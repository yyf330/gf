package main

import (
	"fmt"

	"github.com/yyf330/gf/frame/g"
)

func main() {
	fmt.Println(g.Config().Get("none"))
}
