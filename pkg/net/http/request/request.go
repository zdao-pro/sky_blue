package request

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/zdao-pro/sky_blue/pkg/log"
)

//NewRequest ..
/*
	@param1: timeout itime.Duration
	@param2: headers map[string]string
	@param2: cookies map[string]string
*/
func NewRequest(c context.Context, arg ...interface{}) *Request {
	r := &Request{
		timeout: 30,
		headers: map[string]string{},
		cookies: map[string]string{},
		Context: c,
	}
	l := len(arg)
	if l > 0 {
		if t, ok := arg[0].(time.Duration); ok {
			r.timeout = t
		}
	}
	if l > 1 {
		if headers, ok := arg[1].(map[string]string); ok {
			r.headers = headers
		}
	}

	if l > 2 {
		if cookies, ok := arg[2].(map[string]string); ok {
			r.cookies = cookies
		}
	}

	return r
}

//Request ..
type Request struct {
	context.Context
	cli               *http.Client
	transport         *http.Transport
	debug             bool
	url               string
	method            string
	time              int64
	timeout           time.Duration
	headers           map[string]string
	cookies           map[string]string
	username          string
	password          string
	data              interface{}
	disableKeepAlives bool
	tlsClientConfig   *tls.Config
	jar               http.CookieJar
	proxy             func(*http.Request) (*url.URL, error)
	checkRedirect     func(req *http.Request, via []*http.Request) error
}

func (r *Request) DisableKeepAlives(v bool) *Request {
	r.disableKeepAlives = v
	return r
}

func (r *Request) Jar(v http.CookieJar) *Request {
	r.jar = v
	return r
}

func (r *Request) CheckRedirect(v func(req *http.Request, via []*http.Request) error) *Request {
	r.checkRedirect = v
	return r
}

func (r *Request) TLSClient(v *tls.Config) *Request {
	return r.SetTLSClient(v)
}

func (r *Request) SetTLSClient(v *tls.Config) *Request {
	r.tlsClientConfig = v
	return r
}

func (r *Request) Proxy(v func(*http.Request) (*url.URL, error)) *Request {
	r.proxy = v
	return r
}

func (r *Request) Transport(v *http.Transport) *Request {
	r.transport = v
	return r
}

// Debug model
func (r *Request) Debug(v bool) *Request {
	r.debug = v
	return r
}

// Get transport
func (r *Request) getTransport() http.RoundTripper {
	if r.transport == nil {
		return http.DefaultTransport
	}

	r.transport.DisableKeepAlives = r.disableKeepAlives

	if r.tlsClientConfig != nil {
		r.transport.TLSClientConfig = r.tlsClientConfig
	}

	if r.proxy != nil {
		r.transport.Proxy = r.proxy
	}

	return http.RoundTripper(r.transport)
}

// Build client
func (r *Request) buildClient() *http.Client {
	if r.cli == nil {
		r.cli = &http.Client{
			Transport:     r.getTransport(),
			Jar:           r.jar,
			CheckRedirect: r.checkRedirect,
			Timeout:       time.Second * r.timeout,
		}
	}
	return r.cli
}

// Set headers
func (r *Request) SetHeaders(headers map[string]string) *Request {
	if headers != nil || len(headers) > 0 {
		for k, v := range headers {
			r.headers[k] = v
		}
	}
	return r
}

// Init headers
func (r *Request) initHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range r.headers {
		req.Header.Set(k, v)
	}
}

// Set cookies
func (r *Request) SetCookies(cookies map[string]string) *Request {
	if cookies != nil || len(cookies) > 0 {
		for k, v := range cookies {
			r.cookies[k] = v
		}
	}
	return r
}

// Init cookies
func (r *Request) initCookies(req *http.Request) {
	for k, v := range r.cookies {
		req.AddCookie(&http.Cookie{
			Name:  k,
			Value: v,
		})
	}
}

// Set basic auth
func (r *Request) SetBasicAuth(username, password string) *Request {
	r.username = username
	r.password = password
	return r
}

func (r *Request) initBasicAuth(req *http.Request) {
	if r.username != "" && r.password != "" {
		req.SetBasicAuth(r.username, r.password)
	}
}

// Check application/json
func (r *Request) isJson() bool {
	if len(r.headers) > 0 {
		for _, v := range r.headers {
			if strings.Contains(strings.ToLower(v), "application/json") {
				return true
			}
		}
	}
	return false
}

func (r *Request) JSON() *Request {
	r.SetHeaders(map[string]string{"Content-Type": "application/json"})
	return r
}

// Build query data
func (r *Request) buildBody(d ...interface{}) (io.Reader, error) {
	if r.method == "DELETE" || len(d) < 2 {
		return nil, nil
	}

	switch d[1].(type) {
	case string:
		return strings.NewReader(d[1].(string)), nil
	case []byte:
		return bytes.NewReader(d[1].([]byte)), nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return bytes.NewReader(IntByte(d[1])), nil
	case *bytes.Reader:
		return d[1].(*bytes.Reader), nil
	case *strings.Reader:
		return d[1].(*strings.Reader), nil
	case *bytes.Buffer:
		return d[1].(*bytes.Buffer), nil
	default:
		b, err := json.Marshal(d[1])
		if err != nil {
			return nil, err
		}
		return bytes.NewReader(b), nil
	}

	t := reflect.TypeOf(d[1]).String()
	if !strings.Contains(t, "map[string]interface") {
		return nil, errors.New("Unsupported data type.")
	}

	data := make([]string, 0)
	for k, v := range d[0].(map[string]interface{}) {
		if s, ok := v.(string); ok {
			data = append(data, fmt.Sprintf("%s=%v", k, s))
			continue
		}
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		data = append(data, fmt.Sprintf("%s=%s", k, string(b)))
	}

	return strings.NewReader(strings.Join(data, "&")), nil
}

func (r *Request) SetTimeout(d time.Duration) *Request {
	r.timeout = d
	return r
}

// Parse query for GET request
func parseQuery(url string) ([]string, error) {
	urlList := strings.Split(url, "?")
	if len(urlList) < 2 {
		return make([]string, 0), nil
	}
	query := make([]string, 0)
	for _, val := range strings.Split(urlList[1], "&") {
		v := strings.Split(val, "=")
		if len(v) < 2 {
			return make([]string, 0), errors.New("query parameter error")
		}
		query = append(query, fmt.Sprintf("%s=%s", v[0], v[1]))
	}
	return query, nil
}

// Build GET request url
func buildUrl(url string, data ...interface{}) (string, error) {
	query, err := parseQuery(url)
	if err != nil {
		return url, err
	}

	if len(data) > 0 && data[0] != nil {
		t := reflect.TypeOf(data[0]).String()
		switch t {
		case "map[string]interface {}":
			for k, v := range data[0].(map[string]interface{}) {
				vv := ""
				if reflect.TypeOf(v).String() == "string" {
					vv = v.(string)
				} else {
					b, err := json.Marshal(v)
					if err != nil {
						return url, err
					}
					vv = string(b)
				}
				query = append(query, fmt.Sprintf("%s=%s", k, vv))
			}
		case "string":
			param := data[0].(string)
			if param != "" {
				query = append(query, param)
			}
		default:
			return url, errors.New("Unsupported data type.")
		}

	}

	list := strings.Split(url, "?")

	if len(query) > 0 {
		return fmt.Sprintf("%s?%s", list[0], strings.Join(query, "&")), nil
	}

	return list[0], nil
}

func (r *Request) elapsedTime(n int64, resp *Response) {
	end := time.Now().UnixNano() / 1e6
	resp.time = end - n
}

func (r *Request) log() {
	if r.debug {
		fmt.Printf("[HttpRequest]\n")
		fmt.Printf("-------------------------------------------------------------------\n")
		fmt.Printf("Request: %s %s\nHeaders: %v\nCookies: %v\nTimeout: %ds\nReqBody: %v\n\n", r.method, r.url, r.headers, r.cookies, r.timeout, r.data)
		//fmt.Printf("-------------------------------------------------------------------\n\n")
	}
}

// Get is a get http request
func (r *Request) Get(url string, data ...interface{}) (*Response, error) {
	return r.request(http.MethodGet, url, data...)
}

// Post is a post http request
func (r *Request) Post(url string, data ...interface{}) (*Response, error) {
	return r.request(http.MethodPost, url, data...)
}

// Put is a put http request
func (r *Request) Put(url string, data ...interface{}) (*Response, error) {
	return r.request(http.MethodPut, url, data...)
}

// Delete is a delete http request
func (r *Request) Delete(url string, data ...interface{}) (*Response, error) {
	return r.request(http.MethodDelete, url, data...)
}

// Upload file
func (r *Request) Upload(url, filename, fileinput string) (*Response, error) {
	return r.sendFile(url, filename, fileinput)
}

// Send http request
func (r *Request) request(method, url string, data ...interface{}) (*Response, error) {
	// Build Response

	response := &Response{}

	// Start time
	start := time.Now().UnixNano() / 1e6
	// Count elapsed time
	defer r.elapsedTime(start, response)

	if method == "" || url == "" {
		return nil, errors.New("parameter method and url is required")
	}

	url = handleURL(url)

	s := opentracing.SpanFromContext(r.Context)
	var span2 opentracing.Span
	if nil != s {
		span2 = opentracing.StartSpan(url, opentracing.ChildOf(s.Context()))
		span2.SetTag("url", url)
		defer span2.Finish()
	}
	// Debug infomation
	defer r.log()

	r.url = url
	if len(data) > 0 {
		r.data = data[0]
	} else {
		r.data = ""
	}

	var (
		err  error
		req  *http.Request
		body io.Reader
	)
	r.cli = r.buildClient()

	method = strings.ToUpper(method)
	r.method = method

	if method == "GET" || method == "DELETE" || method == "POST" {
		url, err = buildUrl(url, data...)
		if err != nil {
			return nil, err
		}
		r.url = url
	}

	body, err = r.buildBody(data...)
	if err != nil {
		return nil, err
	}

	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	r.initHeaders(req)
	r.initCookies(req)
	r.initBasicAuth(req)

	//trace decorate
	if nil != s {
		ecarrier := opentracing.TextMapWriter(req.Header)
		er := opentracing.GlobalTracer().Inject(span2.Context(), opentracing.HTTPHeaders, ecarrier)
		if nil != er {
			log.Warnc(r.Context, er.Error())
		}
	}

	// fmt.Println("ff", ecarrier)

	resp, err := r.cli.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Warnc(r.Context, "[Http Request error]StatusCode:%d,url:%s,header:%v,request_body:%v,response:%v", resp.StatusCode, req.URL, req.Header, "", string(response.body))
	}

	response.url = url
	response.resp = resp

	return response, nil
}

// Send file
func (r *Request) sendFile(url, filename, fileinput string) (*Response, error) {
	if url == "" {
		return nil, errors.New("parameter url is required")
	}

	url = handleURL(url)

	fileBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(fileBuffer)
	fileWriter, er := bodyWriter.CreateFormFile(fileinput, filename)
	if er != nil {
		return nil, er
	}

	f, er := os.Open(filename)
	if er != nil {
		return nil, er
	}
	defer f.Close()

	_, er = io.Copy(fileWriter, f)
	if er != nil {
		return nil, er
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	// Build Response
	response := &Response{}

	// Start time
	start := time.Now().UnixNano() / 1e6
	// Count elapsed time
	defer r.elapsedTime(start, response)

	// Debug infomation
	defer r.log()

	r.url = url
	r.data = nil

	var (
		err error
		req *http.Request
	)
	r.cli = r.buildClient()
	r.method = "POST"

	req, err = http.NewRequest(r.method, url, fileBuffer)
	if err != nil {
		return nil, err
	}

	r.initHeaders(req)
	r.initCookies(req)
	r.initBasicAuth(req)
	req.Header.Set("Content-Type", contentType)

	resp, err := r.cli.Do(req)
	if err != nil {
		return nil, err
	}

	response.url = url
	response.resp = resp

	return response, nil
}
