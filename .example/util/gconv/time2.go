package main

import (
	"fmt"

	"github.com/yyf330/gf/util/gconv"
)

func main() {
	fmt.Println(gconv.Time("2018-06-07").String())

	fmt.Println(gconv.Time("2018-06-07 13:01:02").String())

	fmt.Println(gconv.Time("2018-06-07 13:01:02.096").String())

}
