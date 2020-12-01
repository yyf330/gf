package main

import (
	"time"

	"github.com/yyf330/gf/os/glog"
	"github.com/yyf330/gf/os/gtimer"
)

func main() {
	interval := time.Second
	gtimer.AddTimes(interval, 2, func() {
		glog.Println("doing1")
	})
	gtimer.AddTimes(interval, 2, func() {
		glog.Println("doing2")
	})

	select {}
}
