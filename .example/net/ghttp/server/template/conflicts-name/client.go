package main

import (
	"github.com/yyf330/gf/frame/g"
	"github.com/yyf330/gf/net/ghttp"
)

// https://github.com/yyf330/gf/issues/437
func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.WriteTpl("client/layout.html")
	})
	s.SetPort(8199)
	s.Run()
}
