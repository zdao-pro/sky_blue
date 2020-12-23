package gin

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/zdao-pro/sky_blue/pkg/log"
	"github.com/zdao-pro/sky_blue/pkg/net/http/gin"
)

//Server is a gin Engine
var Server *gin.Engine

type param struct {
	A int       `form:"a" need:"true" message:"a参数缺失"`
	B bool      `form:"b"`
	C string    `form:"c" need:"true" default:"c" regexp:"\\d+"`
	T time.Time `form:"t" need:"true" default:"now" time_format:"unix"`
}

//Init http server
func TestGin(t *testing.T) {
	log.Init(nil)
	Server = gin.Default()
	Server.GET("/ping", func(c *gin.Context) {
		var p param
		err := c.ShouldBindQuery(&p)
		if nil != err {
			fmt.Println(err.Error())
		}
		fmt.Println(p)
		c.AbortWithStatus(http.StatusOK)
	})
	//Server.Run()
}

type user struct {
	a int  `from:"sun" need:"true" time_format:"unix" regexp:"\\d+"`
	b bool `from:"sun"`
}

func TestReflect(t *testing.T) {
	var u user
	v := reflect.ValueOf(u)
	tValue := reflect.TypeOf(u)
	for i := 0; i < v.NumField(); i++ {
		t := tValue.Field(i)
		fmt.Println(t.Tag.Get("regexp"))
	}
	b, err := regexp.MatchString("\\d+", "123")
	if nil != err {
		fmt.Println(err.Error())
	}
	fmt.Println(b)
}
