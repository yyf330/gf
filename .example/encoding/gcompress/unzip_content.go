package main

import (
	"fmt"

	"github.com/yyf330/gf/encoding/gcompress"
	"github.com/yyf330/gf/os/gfile"
)

func main() {
	err := gcompress.UnZipContent(
		gfile.GetBytes(`D:\Workspace\Go\GOPATH\src\github.com\gogf\gf\geg\encoding\gcompress\data.zip`),
		`D:\Workspace\Go\GOPATH\src\github.com\gogf\gf\geg`,
	)
	fmt.Println(err)
}
