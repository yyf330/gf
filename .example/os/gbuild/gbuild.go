package main

import (
	"github.com/yyf330/gf/frame/g"
	"github.com/yyf330/gf/os/gbuild"
)

func main() {
	g.Dump(gbuild.Info())
	g.Dump(gbuild.Map())
}
