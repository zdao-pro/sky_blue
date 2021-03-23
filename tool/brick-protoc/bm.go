package main

import (
	"os/exec"
)

const (
	_getBMGen = "go get -u github.com/zdao-pro/sky_blue/tool/protobuf/protoc-gen-gin"
	_bmProtoc = "protoc --proto_path=%s --proto_path=%s --proto_path=%s --bm_out=:."
)

func installBMGen() error {
	if _, err := exec.LookPath("protoc-gen-gin"); err != nil {
		if err := goget(_getBMGen); err != nil {
			return err
		}
	}
	return nil
}

func genBM(files []string) error {
	return generate(_bmProtoc, files)
}
