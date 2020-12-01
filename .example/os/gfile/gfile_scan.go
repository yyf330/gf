package main

import (
	"github.com/yyf330/gf/os/gfile"
	"github.com/yyf330/gf/util/gutil"
)

func main() {
	gutil.Dump(gfile.ScanDir("/Users/john/Documents", "*.*"))
	gutil.Dump(gfile.ScanDir("/home/john/temp/newproject", "*", true))
}
