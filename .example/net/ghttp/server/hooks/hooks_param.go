package main

import (
	"github.com/yyf330/gf/frame/g"
	"github.com/yyf330/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln(r.Get("name"))
	})
	s.BindHookHandlerByMap("/", map[string]ghttp.HandlerFunc{
		ghttp.HOOK_BEFORE_SERVE: func(r *ghttp.Request) {
			r.SetParam("name", "john")
		},
	})
	s.SetPort(8199)
	s.Run()
}
