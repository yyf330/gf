package main

import (
	"fmt"

	"github.com/yyf330/gf/frame/g"
	"github.com/yyf330/gf/util/gconv"
)

func main() {
	type Score struct {
		Name   string
		Result int
	}
	type User struct {
		Scores []Score
	}

	user := new(User)
	scores := g.Map{
		"Scores": g.Map{
			"Name":   "john",
			"Result": 100,
		},
	}

	// 嵌套struct转换，属性为slice类型，数值为map类型
	if err := gconv.Struct(scores, user); err != nil {
		fmt.Println(err)
	} else {
		g.Dump(user)
	}
}
