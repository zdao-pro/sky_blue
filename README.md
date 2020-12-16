[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)

# Sky_blue
sky_blue是一套Go微服务框架，包含大量微服务相关框架及工具。  

> 名字起的随意,但是出自真心

## Features
* bone：是对[gin](https://github.com/gin-gonic/gin)简单包装，简单易用、核心足够轻量；
* Cache：优雅的接口化设计，非常方便的缓存序列化，推荐结合代理模式[overlord](https://github.com/bilibili/overlord)；
* database：集成MySQL，添加熔断保护和统计支持，可快速发现数据层压力；
* peach：方便易用的配置中心，目前实现集成apollo远程配置中心，实现配置版本管理和更新；
* log：类似[zap](https://github.com/uber-go/zap)的field实现高性能日志库,集成syslog,intsig的udp日志；
* trace：基于opentracing，集成了全链路trace支持；
* env：项目环境配置相关代码,可配置DEV,UAT,PRO,方便获取当前服务环境信息；
* util：集成公共工具类函数；
* tool：工具链，可快速生成标准项目文件；

## Quick start

### Requirments

Go version>=1.13

### Installation
```shell
# Linux/macOS
go get github.com/zdao-pro/sky_blue
```

## Documentation

> [简体中文]() 

## License
Kratos is under the MIT license. See the [LICENSE](./LICENSE) file for details.

-------------

*Please report bugs, concerns, suggestions by issues,
email: 1543510543@qq.com to discuss problems around source code.*
