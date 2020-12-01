package main

import "github.com/yyf330/gf/frame/g"

func main() {
	s := g.Server()
	s.SetDenyRoutes([]string{
		"/config*",
	})
	s.SetPort(8299)
	s.Run()
}
