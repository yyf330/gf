package main

import (
	"github.com/yyf330/gf/frame/g"
	"github.com/yyf330/gf/net/ghttp"
)

// 试试模糊匹配规则不带名称会怎么样
func main() {
	s := g.Server()
	s.BindHandler("/hello/*", func(r *ghttp.Request) {
		r.Response.Writeln("哈喽世界！")
	})
	s.SetPort(8199)
	s.Run()
}
