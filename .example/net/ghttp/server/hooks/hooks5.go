package main

import (
	"github.com/yyf330/gf/frame/g"
	"github.com/yyf330/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHookHandler("/*any", ghttp.HOOK_BEFORE_SERVE, func(r *ghttp.Request) {
		r.Response.Writeln("/*any")
	})
	s.BindHookHandler("/v1/*", ghttp.HOOK_BEFORE_SERVE, func(r *ghttp.Request) {
		r.Response.Writeln("/v1/*")
		r.ExitHook()
	})
	s.SetPort(8199)
	s.Run()
}
