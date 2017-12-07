goWebApp  
======
一款基于go的个人网站

本站呢，后端是使用Go所写，框架为beego，数据库为SQLite，使用了beedb的orm。前端则用了H5、bootstrap，和jQuery.

第一页是主站简介，第二页是音乐站简介，第三页是照片墙，点击后可以查看大图，第四页是一个播放器，第五页是个粒子滤镜的飞行战斗游戏。

## Installation
Go version >= 1.3.
```bash
首先需要安装go，如已安装可跳过
export GOROOT=$HOME/go  
export GOBIN=$GOROOT/bin
export PATH=$PATH:$GOROOT/bin

```



## Build
```bash
配置Go环境变量

To install bee use the go get command:
go get github.com/beego/bee

export GOPATH=/项目所放路径/goWebApp
cd $GOPATH/src/webApp
bee run
```
over
liuqiong
