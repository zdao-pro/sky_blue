package apollo

import (
	"fmt"
	"testing"
)

func TestEnv(t *testing.T) {
	var d appoloDriver = appoloDriver{}
	c, err := d.New("zdao_backend.sky_blue")
	if nil != err {
		fmt.Print("vss")
	}

	fmt.Println(c)
	v := c.Get("name")
	if nil == v {
		fmt.Print("v")
	}
	fmt.Print(v.String())
}
