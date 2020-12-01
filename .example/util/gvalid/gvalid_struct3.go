package main

import (
	"github.com/yyf330/gf/frame/g"
	"github.com/yyf330/gf/util/gvalid"
)

// same校验
func main() {
	type User struct {
		Pass string `gvalid:"passwd1 @required|length:2,20|password3||密码强度不足"`
	}

	user := &User{
		Pass: "1",
	}

	g.Dump(gvalid.CheckStruct(user, nil).Maps())
}
