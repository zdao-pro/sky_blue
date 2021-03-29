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
		"extend_flag": 1,
	}
	rs, err := r.Get("$api_server/internal/enterprise/getMultiCompanyBrief", p, []string{"eeqeq"})
	if err != nil {
		panic(err)
	}
	fmt.Println(rs.Content())
}
