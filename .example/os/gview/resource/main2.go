package main

import (
	"fmt"

	"github.com/yyf330/gf/frame/g"
	"github.com/yyf330/gf/os/gres"
	_ "github.com/yyf330/gf/os/gres/testdata"
)

func main() {
	gres.Dump()

	v := g.View()
	v.SetPath("files/template/layout2")
	s, err := v.Parse("layout.html", g.Map{
		"mainTpl": "main/main1.html",
	})
	fmt.Println(err)
	fmt.Println(s)
}
