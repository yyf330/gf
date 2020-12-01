package main

import (
	"github.com/yyf330/gf/frame/g"
	"github.com/yyf330/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/main1", func(r *ghttp.Request) {
		r.Response.WriteTpl("layout.html", g.Map{
			"name":    "smith",
			"mainTpl": "main/main1.html",
		})
	})
	s.BindHandler("/main2", func(r *ghttp.Request) {
		r.Response.WriteTpl("layout.html", g.Map{
			"name":    "john",
			"mainTpl": "main/main2.html",
		})
	})
	s.SetPort(8199)
	s.Run()
}
