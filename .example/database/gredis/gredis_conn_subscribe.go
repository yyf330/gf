package main

import (
	"fmt"

	"github.com/yyf330/gf/frame/g"
	"github.com/yyf330/gf/util/gconv"
)

func main() {
	conn := g.Redis().Conn()
	defer conn.Close()
	_, err := conn.Do("SUBSCRIBE", "channel")
	if err != nil {
		panic(err)
	}
	for {
		reply, err := conn.Receive()
		if err != nil {
			panic(err)
		}
		fmt.Println(gconv.Strings(reply))
	}
}
