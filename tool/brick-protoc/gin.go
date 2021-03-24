package main

import (
	"fmt"
	"os/exec"
)

const (
	_getBMGen = "go get -u github.com/zdao-pro/sky_blue/tool/protobuf/protoc-gen-bm"
	_bmProtoc = "protoc --proto_path=%s --proto_path=%s --proto_path=%s --bm_out=:."
)

func installBMGen() error {
	if _, err := exec.LookPath("protoc-gen-bm"); err != nil {
		if err := goget(_getBMGen); err != nil {
			return err
		}
	}
	return nil
}

func genBM(files []string) error {
	fmt.Println("files", files)
	err := generate(_bmProtoc, files)
	if err != nil {
		fmt.Println("err_r:", err.Error())
	}
	return err
}
