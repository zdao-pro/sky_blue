package peach

import (
	"fmt"
	"testing"
)

func TestEnv(t *testing.T) {
	ad := appoloDriver{}
	ac, err := ad.New([]string{"zdao_backend.sky_blue", "zdao_backend.common"})
	if nil != err {
		panic(err)
	}

	v := ac.Get("test_int")
	fmt.Print(v.Int())
}
