package main

import (
	"time"

	"github.com/yyf330/gf/os/glog"
	"github.com/yyf330/gf/os/gtimer"
)

func main() {
	interval := time.Second
	gtimer.AddSingleton(interval, func() {
		glog.Println("doing")
		time.Sleep(5 * time.Second)
	})

	select {}
}
