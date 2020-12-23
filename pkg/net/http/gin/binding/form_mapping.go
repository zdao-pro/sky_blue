// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/zdao-pro/sky_blue/pkg/net/http/gin/internal/bytesconv"
	"github.com/zdao-pro/sky_blue/pkg/net/http/gin/internal/json"
)

var errUnknownType = errors.New("unknown type")

func mapUri(ptr interface{}, m map[string][]string) error {
	return mapFormByTag(ptr, m, "uri")
}

func mapForm(ptr interface{}, form map[string][]string) error {
	// fmt.Println("form:", form)
	return mapFormByTag(ptr, form, "form")
}

var emptyField = reflect.StructField{}

func mapFormByTag(ptr interface{}, form map[string][]string, tag string) error {
	return mappingByPtr(ptr, formSource(form), tag)
}

// setter tries to set value on a walking by fields of a struct
type setter interface {
	TrySet(value reflect.Value, field reflect.StructField, key string, opt setOptions) (isSetted bool, err error)
}

type formSource map[string][]string

var _ setter = formSource(nil)

// TrySet tries to set a value by request's form source (like map[string][]string)
func (form formSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string, opt setOptions) (isSetted bool, err error) {
	isSetOk, err := setByForm(value, field, form, tagValue, opt)
	if nil != err && "" != opt.message {
		return isSetOk, fmt.Errorf(opt.message)
	}
	return isSetOk, err
}

func mappingByPtr(ptr interface{}, setter setter, tag string) error {
	_, err := mapping(reflect.ValueOf(ptr), emptyField, setter, tag)
	return err
}

func mapping(value reflect.Value, field reflect.StructField, setter setter, tag string) (bool, error) {
	if field.Tag.Get(tag) == "-" { // just ignoring this field
		return false, nil
	}

	var vKind = value.Kind()

	if vKind == reflect.Ptr {
		var isNew bool
		vPtr := value
		if value.IsNil() {
			isNew = true
			vPtr = reflect.New(value.Type().Elem())
		}
		isSetted, err := mapping(vPtr.Elem(), field, setter, tag)
		if err != nil {
			return false, err
		}
		if isNew && isSetted {
			value.Set(vPtr)
		}
		return isSetted, nil
	}

	if vKind != reflect.Struct || !field.Anonymous {
		ok, err := tryToSetValue(value, field, setter, tag)
		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}

	if vKind == reflect.Struct {
		tValue := value.Type()

		var isSetted bool
		for i := 0; i < value.NumField(); i++ {
			sf := tValue.Field(i)
			if sf.PkgPath != "" && !sf.Anonymous { // unexported
				continue
			}
			ok, err := mapping(value.Field(i), tValue.Field(i), setter, tag)
			if err != nil {
				return false, err
			}
			isSetted = isSetted || ok
		}
		return isSetted, nil
	}
	return false, nil
}

type setOptions struct {
	isDefaultExists bool //是否有默认值
	defaultValue    string
	isNeed          bool //参数是否必需
	regexp          string
	isRegexp        bool   //是否需要正则匹配
	message         string //指定的错误信息
	assert          string
	isAssert        bool
	length          int
	isLimitLength   bool
	isPatternRegexp bool
	pattern         string
	gtInt           int64
	ltInt           int64
	intFlag         bool
	intEqualFlag    bool
}

func tryToSetValue(value reflect.Value, field reflect.StructField, setter setter, tag string) (bool, error) {
	var tagValue string
	var setOpt setOptions

	tagValue = field.Tag.Get(tag)
	// fmt.Println("etagValue:" + tagValue + "tag:" + tag)
	tagValue, opt := head(tagValue, ",")
	if tagValue == "" { // default value is FieldName
		tagValue = field.Name
	}
	if tagValue == "" { // when field is "emptyField" variable
		return false, nil
	}

	if opt == "" {
	}

	if k, ok := field.Tag.Lookup("default"); ok {
		setOpt.isDefaultExists = true
		setOpt.defaultValue = k
	}

	if "true" == field.Tag.Get("need") {
		setOpt.isNeed = true
	}

	if k, ok := field.Tag.Lookup("regexp"); ok {
		setOpt.regexp = k
		setOpt.isRegexp = true
	}

	if k, ok := field.Tag.Lookup("message"); ok {
		setOpt.message = k
	}

	if k, ok := field.Tag.Lookup("assert"); ok {
		setOpt.isAssert = true
		setOpt.assert = k
	}

	if k, ok := field.Tag.Lookup("pattern"); ok {
		setOpt.isPatternRegexp = true
		setOpt.regexp = patternMap[k]
		setOpt.pattern = k
	}

	gt, gtOk := field.Tag.Lookup("gt")
	lt, ltOk := field.Tag.Lookup("lt")
	ge, geOk := field.Tag.Lookup("ge")
	le, leOk := field.Tag.Lookup("le")

	if gtOk || ltOk {
		setOpt.intFlag = true
		if v, err := strconv.ParseInt(gt, 10, 64); nil == err {
			setOpt.gtInt = v
		} else {
			setOpt.gtInt = math.MinInt64
		}
		if v, err := strconv.ParseInt(lt, 10, 64); nil == err {
			setOpt.ltInt = v
		} else {
			setOpt.ltInt = math.MaxInt64
		}

	}

	if geOk || leOk {
		setOpt.intFlag = true
		setOpt.intEqualFlag = true
		if v, err := strconv.ParseInt(ge, 10, 64); nil == err {
			setOpt.gtInt = v
		} else {
			setOpt.gtInt = math.MinInt64
		}
		if v, err := strconv.ParseInt(le, 10, 64); nil == err {
			setOpt.ltInt = v
		} else {
			setOpt.ltInt = math.MaxInt64
		}
	}

	if k, ok := field.Tag.Lookup("length"); ok {
		if l, err := strconv.ParseUint(k, 10, 16); nil == err {
			setOpt.isLimitLength = true
			setOpt.length = int(l)
		}
	}

	// var opt string
	// for len(opts) > 0 {
	// 	opt, opts = head(opts, ",")

	// 	if k, v := head(opt, "="); k == "default" {
	// 		setOpt.isDefaultExists = true
	// 		setOpt.defaultValue = v
	// 	}
	// }

	return setter.TrySet(value, field, tagValue, setOpt)
}

func setByForm(value reflect.Value, field reflect.StructField, form map[string][]string, tagValue string, opt setOptions) (isSetted bool, err error) {
	vs, ok := form[tagValue]
	if (!ok || "" == vs[0]) && !opt.isDefaultExists {
		if true == opt.isNeed {
			return false, fmt.Errorf("parm %v is mising", tagValue)
		}
		return false, nil
	}

	switch value.Kind() {
	case reflect.Slice:
		if !ok {
			vs = []string{opt.defaultValue}
		}
		return true, setSlice(vs, value, field)
	case reflect.Array:
		if !ok {
			vs = []string{opt.defaultValue}
		}
		if len(vs) != value.Len() {
			return false, fmt.Errorf("%q is not valid value for %s", vs, value.Type().String())
		}
		return true, setArray(vs, value, field)
	default:
		var val string
		if !ok || "" == vs[0] {
			val = opt.defaultValue
		}

		if len(vs) > 0 {
			val = vs[0]
		}

		if true == opt.isLimitLength && len(val) != opt.length {
			return false, fmt.Errorf("the length of %s length is equal %d", tagValue, opt.length)
		}

		if true == opt.isAssert && opt.assert != val {
			return false, fmt.Errorf("the param %s assert error", tagValue)
		}

		if true == opt.intFlag {
			v, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return false, err
			}
			if true == opt.intEqualFlag {
				if !(v >= opt.gtInt && v <= opt.ltInt) {
					return false, fmt.Errorf("the param %s cannot >= %d <= %d", tagValue, opt.gtInt, opt.ltInt)
				}
			} else {
				if !(v > opt.gtInt && v < opt.ltInt) {
					return false, fmt.Errorf("the param %s cannot > %d < %d", tagValue, opt.gtInt, opt.ltInt)
				}
			}
		}

		if true == opt.isRegexp || true == opt.isPatternRegexp {
			//正则匹配
			r, err := regexp.Compile(opt.regexp)
			if nil != err {
				return false, err
			}
			b := r.MatchString(val)
			if true != b {
				if true == opt.isPatternRegexp {
					return false, fmt.Errorf("the param %s cannot match %s", tagValue, opt.pattern)
				}
				return false, fmt.Errorf("regexp match %s is error", tagValue)
			}
		}
		return true, setWithProperType(val, value, field)
	}
}

func setWithProperType(val string, value reflect.Value, field reflect.StructField) error {
	switch value.Kind() {
	case reflect.Int:
		return setIntField(val, 0, value)
	case reflect.Int8:
		return setIntField(val, 8, value)
	case reflect.Int16:
		return setIntField(val, 16, value)
	case reflect.Int32:
		return setIntField(val, 32, value)
	case reflect.Int64:
		switch value.Interface().(type) {
		case time.Duration:
			return setTimeDuration(val, value, field)
		}
		return setIntField(val, 64, value)
	case reflect.Uint:
		return setUintField(val, 0, value)
	case reflect.Uint8:
		return setUintField(val, 8, value)
	case reflect.Uint16:
		return setUintField(val, 16, value)
	case reflect.Uint32:
		return setUintField(val, 32, value)
	case reflect.Uint64:
		return setUintField(val, 64, value)
	case reflect.Bool:
		return setBoolField(val, value)
	case reflect.Float32:
		return setFloatField(val, 32, value)
	case reflect.Float64:
		return setFloatField(val, 64, value)
	case reflect.String:
		value.SetString(val)
	case reflect.Struct:
		switch value.Interface().(type) {
		case time.Time:
			return setTimeField(val, field, value)
		}
		return json.Unmarshal(bytesconv.StringToBytes(val), value.Addr().Interface())
	case reflect.Map:
		return json.Unmarshal(bytesconv.StringToBytes(val), value.Addr().Interface())
	default:
		return errUnknownType
	}
	return nil
}

func setIntField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0"
	}
	intVal, err := strconv.ParseInt(val, 10, bitSize)
	if err == nil {
		field.SetInt(intVal)
	}
	return err
}

func setUintField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0"
	}
	uintVal, err := strconv.ParseUint(val, 10, bitSize)
	if err == nil {
		field.SetUint(uintVal)
	}
	return err
}

func setBoolField(val string, field reflect.Value) error {
	if val == "" {
		val = "false"
	}
	boolVal, err := strconv.ParseBool(val)
	if err == nil {
		field.SetBool(boolVal)
	}
	return err
}

func setFloatField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0.0"
	}
	floatVal, err := strconv.ParseFloat(val, bitSize)
	if err == nil {
		field.SetFloat(floatVal)
	}
	return err
}

func setTimeField(val string, structField reflect.StructField, value reflect.Value) error {
	timeFormat := structField.Tag.Get("time_format")
	if timeFormat == "" {
		timeFormat = "unixnano"
	}

	if val == "now" {
		value.Set(reflect.ValueOf(time.Now()))
		return nil
	}

	switch tf := strings.ToLower(timeFormat); tf {
	case "unix", "unixnano":
		tv, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return err
		}

		d := time.Duration(1)
		if tf == "unixnano" {
			tv = tv * 1000000
			d = time.Second
		}

		t := time.Unix(tv/int64(d), tv%int64(d))
		value.Set(reflect.ValueOf(t))
		return nil

	}

	if val == "" {
		value.Set(reflect.ValueOf(time.Time{}))
		return nil
	}

	l := time.Local
	if isUTC, _ := strconv.ParseBool(structField.Tag.Get("time_utc")); isUTC {
		l = time.UTC
	}

	if locTag := structField.Tag.Get("time_location"); locTag != "" {
		loc, err := time.LoadLocation(locTag)
		if err != nil {
			return err
		}
		l = loc
	}

	t, err := time.ParseInLocation(timeFormat, val, l)
	if err != nil {
		return err
	}

	value.Set(reflect.ValueOf(t))
	return nil
}

func setArray(vals []string, value reflect.Value, field reflect.StructField) error {
	for i, s := range vals {
		err := setWithProperType(s, value.Index(i), field)
		if err != nil {
			return err
		}
	}
	return nil
}

func setSlice(vals []string, value reflect.Value, field reflect.StructField) error {
	slice := reflect.MakeSlice(value.Type(), len(vals), len(vals))
	err := setArray(vals, slice, field)
	if err != nil {
		return err
	}
	value.Set(slice)
	return nil
}

func setTimeDuration(val string, value reflect.Value, field reflect.StructField) error {
	d, err := time.ParseDuration(val)
	if err != nil {
		return err
	}
	value.Set(reflect.ValueOf(d))
	return nil
}

func head(str, sep string) (head string, tail string) {
	idx := strings.Index(str, sep)
	if idx < 0 {
		return str, ""
	}
	return str[:idx], str[idx+len(sep):]
}
