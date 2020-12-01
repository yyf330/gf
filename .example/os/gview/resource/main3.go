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
	s, err := v.Parse("index.html")
	fmt.Println(err)
	fmt.Println(s)
}
