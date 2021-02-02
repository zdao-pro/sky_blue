package util

import (
	"fmt"
	"runtime"
	"testing"
)

func runFuncName(skip int) string {
	pc := make([]uintptr, 2)
	runtime.Callers(skip, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func TestFuncName(t *testing.T) {
	fmt.Println(FuncName(2))
}

func TestLocalAddress(t *testing.T) {
	fmt.Println(GetLocalAddress())
}

func TestMD5(t *testing.T) {
	str := "5205720"
	s := MD5([]byte(str))
	t.Error(s)
}

func TestGetUUID(t *testing.T) {
	str := GetUUID()
	t.Error(str)
}
