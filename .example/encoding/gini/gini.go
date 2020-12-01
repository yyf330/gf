package main

import (
	"fmt"
	"github.com/yyf330/gf/encoding/gini"
)

func main() {
	s := `
a = b

`
	m, err := gini.Decode([]byte(s))
	fmt.Println(err)
	fmt.Println(m)
}
