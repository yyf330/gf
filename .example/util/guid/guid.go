package main

import (
	"fmt"
	"github.com/yyf330/gf/util/guid"
)

func main() {
	for i := 0; i < 100; i++ {
		s := guid.S()
		fmt.Println(s, len(s))
	}
}
