package main

import (
	"github.com/yyf330/gf/frame/g"
	"github.com/yyf330/gf/util/gvalid"
)

// string默认值校验
func main() {
	type User struct {
		Uid string `gvalid:"uid@integer"`
	}

	user := &User{}

	g.Dump(gvalid.CheckStruct(user, nil))
}
