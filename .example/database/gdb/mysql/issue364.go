package main

import (
	"github.com/yyf330/gf/frame/g"
	"time"
)

func test1() {
	db := g.DB()
	db.SetDebug(true)
	time.Sleep(1 * time.Minute)
	r, e := db.Table("test").Where("id", 10000).Count()
	if e != nil {
		panic(e)
	}
	g.Dump(r)
}

func test2() {
	db := g.DB()
	db.SetDebug(true)
	dao := db.Table("test").Safe()
	time.Sleep(1 * time.Minute)
	r, e := dao.Where("id", 10000).Count()
	if e != nil {
		panic(e)
	}
	g.Dump(r)
}

func main() {
	test1()
	test2()
}
