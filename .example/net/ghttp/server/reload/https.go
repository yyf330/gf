package main

import (
	"github.com/yyf330/gf/frame/g"
	"github.com/yyf330/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln("哈罗！")
	})
	s.EnableHTTPS("/home/john/temp/server.crt", "/home/john/temp/server.key")
	s.EnableAdmin()
	s.SetPort(8200)
	s.Run()
}
