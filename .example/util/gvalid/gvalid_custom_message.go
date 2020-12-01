package main

import (
	"fmt"
	"github.com/yyf330/gf/frame/g"
	"github.com/yyf330/gf/util/gvalid"
)

func main() {
	g.I18n().SetLanguage("cn")
	err := gvalid.Check("", "required", nil)
	fmt.Println(err.String())
}
