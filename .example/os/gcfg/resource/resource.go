package main

import (
	"github.com/yyf330/gf/frame/g"
	_ "github.com/yyf330/gf/os/gres/testdata"
)

func main() {
	g.Res().Dump()
	g.Dump(g.Config().Get("redis"))

	g.Config().SetFileName("my.ini")
	g.Dump(g.Config().Get("redis"))

	g.Config().SetPath("config-custom")
	g.Config().SetFileName("my.ini")
	g.Dump(g.Config().Get("redis"))

	g.Config().SetFileName("config.toml")
	g.Dump(g.Config().Get("redis"))
}
