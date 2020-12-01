package main

import (
	"github.com/yyf330/gf/frame/g"
	"github.com/yyf330/gf/os/glog"
)

func main() {
	err := glog.SetConfigWithMap(g.Map{
		"prefix": "[TEST]",
	})
	if err != nil {
		panic(err)
	}
	glog.Info(1)
}
