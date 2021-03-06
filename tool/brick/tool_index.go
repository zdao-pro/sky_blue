package main

import "time"

var toolIndexs = []*Tool{
	{
		Name:      "brick",
		Alias:     "brick",
		BuildTime: time.Date(2020, 3, 31, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/zdao-pro/sky_blue/tool/brick@" + Version,
		Summary:   "brick工具集本体",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "brick",
		Hidden:    true,
	},
	{
		Name:      "protoc",
		Alias:     "brick-protoc",
		BuildTime: time.Date(2020, 3, 31, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/zdao-pro/sky_blue/tool/brick-protoc@" + Version,
		Summary:   "快速方便生成pb.go的protoc封装，windows、Linux请先安装protoc工具",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "brick",
	},
	{
		Name:      "genbts",
		Alias:     "brick-gen-bts",
		BuildTime: time.Date(2020, 3, 31, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/zdao-pro/sky_blue/tool/brick-gen-bts@" + Version,
		Summary:   "缓存回源逻辑代码生成器",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "brick",
	},
	{
		Name:      "genmc",
		Alias:     "brick-gen-mc",
		BuildTime: time.Date(2020, 3, 31, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/zdao-pro/sky_blue/tool/brick-gen-mc@" + Version,
		Summary:   "mc缓存代码生成",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "brick",
	},
	{
		Name:         "genproject",
		Alias:        "brick-gen-project",
		Install:      "go get -u github.com/zdao-pro/sky_blue/tool/brick-gen-project@" + Version,
		BuildTime:    time.Date(2020, 3, 31, 0, 0, 0, 0, time.Local),
		Platform:     []string{"darwin", "linux", "windows"},
		Hidden:       true,
		Requirements: []string{"wire"},
	},
	{
		Name:      "testgen",
		Alias:     "testgen",
		BuildTime: time.Date(2020, 3, 31, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/zdao-pro/sky_blue/tool/testgen@" + Version,
		Summary:   "测试代码生成",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "brick",
	},
	{
		Name:      "testcli",
		Alias:     "testcli",
		BuildTime: time.Date(2020, 3, 31, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/zdao-pro/sky_blue/tool/testcli@" + Version,
		Summary:   "测试代码运行",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "brick",
	},
	//  third party
	{
		Name:      "wire",
		Alias:     "wire",
		BuildTime: time.Date(2020, 3, 31, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/google/wire/cmd/wire",
		Platform:  []string{"darwin", "linux", "windows"},
		Hidden:    true,
	},
	{
		Name:      "swagger",
		Alias:     "swagger",
		BuildTime: time.Date(2020, 3, 31, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/go-swagger/go-swagger/cmd/swagger",
		Summary:   "swagger api文档",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "goswagger.io",
	},
}
