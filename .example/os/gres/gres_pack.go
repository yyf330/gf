package main

import (
	"github.com/yyf330/gf/crypto/gaes"
	"github.com/yyf330/gf/os/gfile"
	"github.com/yyf330/gf/os/gres"
)

var (
	CryptoKey = []byte("x76cgqt36i9c863bzmotuf8626dxiwu0")
)

func main() {
	binContent, err := gres.Pack("public,config")
	if err != nil {
		panic(err)
	}
	binContent, err = gaes.Encrypt(binContent, CryptoKey)
	if err != nil {
		panic(err)
	}
	if err := gfile.PutBytes("data.bin", binContent); err != nil {
		panic(err)
	}
}
