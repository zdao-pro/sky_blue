package request

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var (
	r = regexp.MustCompile(`\$[a-zA-Z_]+`)
)

//Upstream ..
type Upstream struct {
	Server    []string `yaml:"server"`    // server addr
	Keepalive int      `yaml:"keepalive"` // keepalive
}

func handleURL(url string) string {
	s := r.FindString(url)
	if s == "" {
		return url
	}

	serverName := s[1:len(s)]
	u, err := UpstreamMap.Get(serverName)
	if err != nil {
		panic(err)
	}

	serverAddr := ""
	serverLen := len(u.Server)
	if 1 > serverLen {
		return url
	}

	if 1 < serverLen {
		rand.Seed(time.Now().Unix())
		serverAddr = u.Server[rand.Intn(serverLen)]
	} else {
		serverAddr = u.Server[0]
	}

	return strings.Replace(url, s, serverAddr, 1)
}
