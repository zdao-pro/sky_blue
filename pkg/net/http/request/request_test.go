package request

import (
	"fmt"
	"testing"

	"github.com/zdao-pro/sky_blue/pkg/log"
)

type M map[string]Upstream

func TestRequest(t *testing.T) {
	log.Init(nil)
	s := `user_server:
            server:
                - pre.zhaodao88.com
            keepalive: 100`
	InitUpstream(s)
	r := NewRequest()
	p := map[string]interface{}{
		"token": "eee",
	}
	rs, err := r.Get("https://$user_server/user/token_check", p)
	if err != nil {
		panic(err)
	}
	fmt.Println(rs.Content())
}
