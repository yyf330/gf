package main

import (
	"github.com/yyf330/gf/frame/g"
	"github.com/yyf330/gf/net/ghttp"
	"github.com/yyf330/gf/os/gtime"
)

func main() {
	s := g.Server()
	s.BindHandler("/cookie", func(r *ghttp.Request) {
		datetime := r.Cookie.Get("datetime")
		r.Cookie.Set("datetime", gtime.Datetime())
		r.Response.Write("datetime:", datetime)
	})
	s.SetPort(8199)
	s.Run()
}
