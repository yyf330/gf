package main

import (
	"encoding/json"
	"fmt"
	"github.com/yyf330/gf/container/gset"
)

func main() {
	type Student struct {
		Id     int
		Name   string
		Scores *gset.IntSet
	}
	s := Student{
		Id:     1,
		Name:   "john",
		Scores: gset.NewIntSetFrom([]int{100, 99, 98}),
	}
	b, _ := json.Marshal(s)
	fmt.Println(string(b))
}
