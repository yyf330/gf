package gbase64

import (
	"fmt"

	"github.com/yyf330/gf/encoding/gbase64"
)

func main() {
	s := "john"
	b := gbase64.Encode(s)
	c, e := gbase64.Decode(b)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(e)
}
