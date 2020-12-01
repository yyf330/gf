// Copyright 2018 gf Author(https://github.com/yyf330/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yyf330/gf.

package ghttp_test

import (
	"fmt"
	"github.com/yyf330/gf/debug/gdebug"
	"github.com/yyf330/gf/os/gfile"
	"github.com/yyf330/gf/os/gtime"
	"github.com/yyf330/gf/text/gstr"
	"testing"
	"time"

	"github.com/yyf330/gf/frame/g"
	"github.com/yyf330/gf/net/ghttp"
	"github.com/yyf330/gf/test/gtest"
)

func Test_Params_File_Single(t *testing.T) {
	dstDirPath := gfile.TempDir(gtime.TimestampNanoStr())
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/upload/single", func(r *ghttp.Request) {
		file := r.GetUploadFile("file")
		if file == nil {
			r.Response.WriteExit("upload file cannot be empty")
		}

		if name, err := file.Save(dstDirPath, r.GetBool("randomlyRename")); err == nil {
			r.Response.WriteExit(name)
		}
		r.Response.WriteExit("upload failed")
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()
	time.Sleep(100 * time.Millisecond)
	// normal name
	gtest.C(t, func(t *gtest.T) {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		srcPath := gdebug.TestDataPath("upload", "file1.txt")
		dstPath := gfile.Join(dstDirPath, "file1.txt")
		content := client.PostContent("/upload/single", g.Map{
			"file": "@file:" + srcPath,
		})
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")
		t.Assert(content, "file1.txt")
		t.Assert(gfile.GetContents(dstPath), gfile.GetContents(srcPath))
	})
	// randomly rename.
	gtest.C(t, func(t *gtest.T) {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		srcPath := gdebug.TestDataPath("upload", "file2.txt")
		content := client.PostContent("/upload/single", g.Map{
			"file":           "@file:" + srcPath,
			"randomlyRename": true,
		})
		dstPath := gfile.Join(dstDirPath, content)
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")
		t.Assert(gfile.GetContents(dstPath), gfile.GetContents(srcPath))
	})
}

func Test_Params_File_CustomName(t *testing.T) {
	dstDirPath := gfile.TempDir(gtime.TimestampNanoStr())
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/upload/single", func(r *ghttp.Request) {
		file := r.GetUploadFile("file")
		if file == nil {
			r.Response.WriteExit("upload file cannot be empty")
		}
		file.Filename = "my.txt"
		if name, err := file.Save(dstDirPath, r.GetBool("randomlyRename")); err == nil {
			r.Response.WriteExit(name)
		}
		r.Response.WriteExit("upload failed")
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		srcPath := gdebug.TestDataPath("upload", "file1.txt")
		dstPath := gfile.Join(dstDirPath, "my.txt")
		content := client.PostContent("/upload/single", g.Map{
			"file": "@file:" + srcPath,
		})
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")
		t.Assert(content, "my.txt")
		t.Assert(gfile.GetContents(dstPath), gfile.GetContents(srcPath))
	})
}

func Test_Params_File_Batch(t *testing.T) {
	dstDirPath := gfile.TempDir(gtime.TimestampNanoStr())
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/upload/batch", func(r *ghttp.Request) {
		files := r.GetUploadFiles("file")
		if files == nil {
			r.Response.WriteExit("upload file cannot be empty")
		}
		if names, err := files.Save(dstDirPath, r.GetBool("randomlyRename")); err == nil {
			r.Response.WriteExit(gstr.Join(names, ","))
		}
		r.Response.WriteExit("upload failed")
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()
	time.Sleep(100 * time.Millisecond)
	// normal name
	gtest.C(t, func(t *gtest.T) {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		srcPath1 := gdebug.TestDataPath("upload", "file1.txt")
		srcPath2 := gdebug.TestDataPath("upload", "file2.txt")
		dstPath1 := gfile.Join(dstDirPath, "file1.txt")
		dstPath2 := gfile.Join(dstDirPath, "file2.txt")
		content := client.PostContent("/upload/batch", g.Map{
			"file[0]": "@file:" + srcPath1,
			"file[1]": "@file:" + srcPath2,
		})
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")
		t.Assert(content, "file1.txt,file2.txt")
		t.Assert(gfile.GetContents(dstPath1), gfile.GetContents(srcPath1))
		t.Assert(gfile.GetContents(dstPath2), gfile.GetContents(srcPath2))
	})
	// randomly rename.
	gtest.C(t, func(t *gtest.T) {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		srcPath1 := gdebug.TestDataPath("upload", "file1.txt")
		srcPath2 := gdebug.TestDataPath("upload", "file2.txt")
		content := client.PostContent("/upload/batch", g.Map{
			"file[0]":        "@file:" + srcPath1,
			"file[1]":        "@file:" + srcPath2,
			"randomlyRename": true,
		})
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")

		array := gstr.SplitAndTrim(content, ",")
		t.Assert(len(array), 2)
		dstPath1 := gfile.Join(dstDirPath, array[0])
		dstPath2 := gfile.Join(dstDirPath, array[1])
		t.Assert(gfile.GetContents(dstPath1), gfile.GetContents(srcPath1))
		t.Assert(gfile.GetContents(dstPath2), gfile.GetContents(srcPath2))
	})
}
