package main

import (
	"fmt"
	"github.com/yyf330/gf/os/gproc"
)

func main() {
	fmt.Println(gproc.Pid())
	err := gproc.ShellRun("sleep 99999s")
	fmt.Println(err)
}
