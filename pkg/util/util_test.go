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
