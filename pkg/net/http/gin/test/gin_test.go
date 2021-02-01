package gin

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/zdao-pro/sky_blue/pkg/database/sql"
	"github.com/zdao-pro/sky_blue/pkg/ecode"
	"github.com/zdao-pro/sky_blue/pkg/log"
	"github.com/zdao-pro/sky_blue/pkg/net/http/gin"
	"github.com/zdao-pro/sky_blue/pkg/net/http/request"
	"github.com/zdao-pro/sky_blue/pkg/net/trace/zipkin"
	"github.com/zdao-pro/sky_blue/pkg/peach"
	_ "github.com/zdao-pro/sky_blue/pkg/peach/apollo"
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
type Misc struct {
	Email string `json:"email"`
	Phone int    `json:"phone"`
}

type UserInfo struct {
	*sql.Model
	ID         int16     `orm:"id"`
	Name       string    `orm:"name"`
	Age        int       `orm:"age"`
	Status     uint8     `orm:"status"`
	RegistTime time.Time `orm:"regist_time"`
	CreateTime time.Time `orm:"create_time" time_format:"unixmilli"`
	MiscInfo   Misc      `orm:"misc"`
}

func (u *UserInfo) Insert(c context.Context, name string, age int) (err error) {
	_, err = u.Exec(c, "insert into user_info set name = ?,age = ?", name, age)
	return
}

func (u *UserInfo) QueryUserByID(c context.Context, id int) error {
	err := u.Select(c, u, "select id,name,age,regist_time,status,create_time,misc from user_info where id = ?", id)
	if nil != err {
		return err
	}
	fmt.Println(u)
	return nil
}

func (u *UserInfo) QueryUserByName(c context.Context, name string) (*[]UserInfo, error) {
	list := make([]UserInfo, 0)
	err := u.Select(c, &list, "select id,name,age,regist_time,status,create_time,misc from user_info where name = ?", name)
	if nil != err {
		return nil, err
	}
	// fmt.Println(list)
	return &list, nil
}

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

//RegisterInfo ..
type RegisterInfo struct {
	Name string `form:"name" validate:"required"`
	Age  int    `form:"age" validate:"min=1"` //required:表示必传 min表示最小值
}

//Init http server
func TestGin(t *testing.T) {
	log.Init(nil)
	zipkin.Init("gin")
	defer zipkin.Close()
	err := peach.Init(peach.PeachDriverApollo, []string{"zdao_backend.sky_blue", "zdao_backend.common"})
	if nil != err {
		panic(err)
	}
	// 初始化http request
	upstreamStr, _ := peach.Get("upstream.yaml").String()
	request.InitUpstream(upstreamStr)
	var c sql.Config
	peach.Get("mysql_test.yaml").UnmarshalYAML(&c)
	db := sql.NewMySQL(&c)
	if db == nil {
		log.Warn("error")
	}
	user := UserInfo{
		Model: sql.NewModel(db),
	}
	// user.Begin(context.Background())
	// user.Insert("test", 50)
	// user.Rollback()

	Server = gin.Default()
	Server.GET("/internal/ping", func(c *gin.Context) {
		// var p param
		// group := ColorGroup{
		// 	ID:     1,
		// 	Name:   "Reds",
		// 	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
		// }
		user.QueryUserByID(c.Context, 145)
		fmt.Println(user)
		r := request.NewRequest(c.Context)
		p := map[string]interface{}{
			"token": "eee",
		}
		rs, err := r.Get("http://127.0.0.1:8080/ping", p)
		if err != nil {
			panic(err)
		}
		fmt.Println(rs.Content())
		// log.Infoc(c.Context, "trace test")
		// fmt.Println("uid:", c.UserID)

		c.Exit(int(ecode.ParamInvaidErr))
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
		// fmt.Println(c.Context)
	})
	Server.GET("/ping", func(c *gin.Context) {
		// var p param
		// group := ColorGroup{
		// 	ID:     1,
		// 	Name:   "Reds",
		// 	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
		// }
		var registerParm RegisterInfo
		//解析请求参数到registerParm
		err := c.ShouldBindQuery(&registerParm)
		if nil != err {
			return
		}
		r := request.NewRequest(c.Context)
		p := map[string]interface{}{
			"token": "eeerwrwwe",
		}
		rs, err := r.Get("https://$user_server/user/token_check", p)
		if err != nil {
			panic(err)
		}
		fmt.Println(rs.Content())
		c.Exit(int(ecode.PermissionErr))
		// c.Exit(int(ecode.ParamInvaidErr))
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
		// fmt.Println(c.Context)
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
