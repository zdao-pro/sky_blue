package peach

import (
	"fmt"
	"testing"

	"github.com/zdao-pro/sky_blue/pkg/peach"
	_ "github.com/zdao-pro/sky_blue/pkg/peach/apollo"
)

func TestApollo(t *testing.T) {
	peach.Init(peach.PeachDriverApollo, "zdao_backend.sky_blue")
	a, _ := peach.Get("name").String()
	fmt.Println(a)
}
