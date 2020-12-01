package main

import (
	"fmt"
	"github.com/yyf330/gf/debug/gdebug"
)

func main() {
	gdebug.PrintStack()
	fmt.Println(gdebug.CallerPackage())
	fmt.Println(gdebug.CallerFunction())
}
