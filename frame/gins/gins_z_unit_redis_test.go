// Copyright 2017 gf Author(https://github.com/yyf330/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yyf330/gf.

package gins_test

import (
	"github.com/yyf330/gf/debug/gdebug"
	"github.com/yyf330/gf/frame/gins"
	"github.com/yyf330/gf/os/gtime"
	"testing"
	"time"

	"github.com/yyf330/gf/os/gfile"
	"github.com/yyf330/gf/test/gtest"
)

func Test_Redis(t *testing.T) {
	redisContent := gfile.GetContents(
		gdebug.TestDataPath("redis", "config.toml"),
	)

	gtest.C(t, func(t *gtest.T) {
		var err error
		dirPath := gfile.TempDir(gtime.TimestampNanoStr())
		err = gfile.Mkdir(dirPath)
		t.Assert(err, nil)
		defer gfile.Remove(dirPath)

		name := "config.toml"
		err = gfile.PutContents(gfile.Join(dirPath, name), redisContent)
		t.Assert(err, nil)

		err = gins.Config().AddPath(dirPath)
		t.Assert(err, nil)

		defer gins.Config().Clear()

		// for gfsnotify callbacks to refresh cache of config file
		time.Sleep(500 * time.Millisecond)

		//fmt.Println("gins Test_Redis", Config().Get("test"))

		redisDefault := gins.Redis()
		redisCache := gins.Redis("cache")
		redisDisk := gins.Redis("disk")
		t.AssertNE(redisDefault, nil)
		t.AssertNE(redisCache, nil)
		t.AssertNE(redisDisk, nil)

		r, err := redisDefault.Do("PING")
		t.Assert(err, nil)
		t.Assert(r, "PONG")

		r, err = redisCache.Do("PING")
		t.Assert(err, nil)
		t.Assert(r, "PONG")

		_, err = redisDisk.Do("SET", "k", "v")
		t.Assert(err, nil)
		r, err = redisDisk.Do("GET", "k")
		t.Assert(err, nil)
		t.Assert(r, []byte("v"))
	})
}
