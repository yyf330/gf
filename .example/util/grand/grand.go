package main

import (
	"fmt"

	"github.com/yyf330/gf/util/grand"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(grand.S(16))
	}
	for i := 0; i < 100; i++ {
		fmt.Println(grand.N(0, 99999999))
	}
}
