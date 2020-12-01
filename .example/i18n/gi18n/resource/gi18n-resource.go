package main

import (
	"fmt"

	"github.com/yyf330/gf/frame/g"

	_ "github.com/yyf330/gf/os/gres/testdata"
)

func main() {
	m := g.I18n()
	m.SetLanguage("ja")
	err := m.SetPath("/i18n-dir")
	if err != nil {
		panic(err)
	}
	fmt.Println(m.Translate(`hello`))
	fmt.Println(m.Translate(`{#hello}{#world}!`))
}
