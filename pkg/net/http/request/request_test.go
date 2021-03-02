package request

import (
	"context"
	"fmt"
	"testing"
)

type M map[string]Upstream

func TestRequest(t *testing.T) {
	r := NewRequest(context.Background())
	p := map[string]interface{}{
		"token": "eee",
	}
	rs, err := r.Get("$api_server/user/token_check", p)
	if err != nil {
		panic(err)
	}
	fmt.Println(rs.Content())
}
