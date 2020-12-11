package peach

import (
	"github.com/zdao-pro/sky_blue/pkg/peach"
	"github.com/zdao-pro/sky_blue/pkg/peach/apollo"
	"fmt"
	"testing"
)

func TestApollo(t *testing.T) {
	peach.Init("/", apollo.PeachDriverApollo)
	a, _ := peach.Get("test").String()
	fmt.Println(a)
}
