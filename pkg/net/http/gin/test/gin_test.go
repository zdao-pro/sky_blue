package gin

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/zdao-pro/sky_blue/pkg/log"
	"github.com/zdao-pro/sky_blue/pkg/net/http/gin"
)

//Server is a gin Engine
var Server *gin.Engine

// type param struct {
// 	A int       `form:"a" need:"true" message:"a参数缺失"`
// 	B bool      `form:"b" need:"false"`
// 	C string    `form:"c" need:"true" default:"c" regexp:"^\\d+$"`
// 	D string    `form:"d" need:"true" assert:"sunjin"`
// 	E string    `form:"e" need:"true" length:"4"`
// 	F string    `form:"f" need:"true" pattern:"email"`
// 	H string    `form:"h" need:"true" pattern:"mobile"`
// 	G string    `form:"g" need:"true" pattern:"common"`
// 	I int       `form:"i" need:"true" gt:"67" lt:"344"`
// 	J int       `form:"j" need:"true" ge:"64" le:"145"`
// 	K []int     `form:"k" need:"true" split:","`
// 	P string    `form:"p" need:"true" split:","`
// 	T time.Time `form:"t" need:"true" default:"now" time_format:"unix"`
// }

type param struct {
	A int       `form:"a" message:"a参数缺失"`
	B string    `form:"b" default:"true"`
	C string    `form:"c" default:"4453" regexp:"^\\d+$"`
	I int       `form:"i" validate:"max=20"`
	K []int     `form:"k" split:","`
	T time.Time `form:"t" default:"now" time_format:"unix"`
}

type s struct {
	Foo string `json:"foo"`
}
type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

//Init http server
func TestGin(t *testing.T) {
	log.Init(nil)
	Server = gin.Default()
	Server.GET("/ping", func(c *gin.Context) {
		// var p param
		group := ColorGroup{
			ID:     1,
			Name:   "Reds",
			Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
		}
		// fmt.Println(a)
		c.Exit(200, group)
		// c.JSON(200, group)
		// err := c.ShouldBindQuery(&p)
		// if nil != err {
		// 	fmt.Println(err.Error())
		// }
		// var b s
		// err := c.ShouldBindJSON(&b)
		// if nil != err {
		// 	fmt.Println(err.Error())
		// }
		// s
	})
	Server.Run()
}

type user struct {
	a int  `from:"sun" need:"true" time_format:"unix" regexp:"\\d+"`
	b bool `from:"sun"`
}

const (
	EmailRegexpStr  = `^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`
	PhoneRegexpStr  = `(^13|14|15|17|18\d{9}$)|(^201|165|195|166|167|168|191|198|199\d{8}$)`
	CommonRegexpStr = `^[a-zA-Z0-9_]+$`
)

func TestReflect(t *testing.T) {
	// var u user
	// v := reflect.ValueOf(u)
	// tValue := reflect.TypeOf(u)
	// for i := 0; i < v.NumField(); i++ {
	// 	t := tValue.Field(i)
	// 	fmt.Println(t.Tag.Get("regexp"))
	// }
	// b, err := regexp.MatchString(EmailRegexpStr, "1543510543@intsig.net")
	// b, err := regexp.MatchString(PhoneRegexpStr, "20121065085")
	// b, err := regexp.MatchString(CommonRegexpStr, "20121065085@")
	r, err := regexp.Compile(CommonRegexpStr)
	if nil != err {
		fmt.Println(err.Error())
	}
	b := r.MatchString("34434")
	fmt.Println(b)
}
