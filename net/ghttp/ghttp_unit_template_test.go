// Copyright 2018 gf Author(https://github.com/yyf330/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yyf330/gf.

// static service testing.

package ghttp_test

import (
	"fmt"
	"github.com/yyf330/gf/debug/gdebug"
	"github.com/yyf330/gf/encoding/ghtml"
	"github.com/yyf330/gf/os/gview"
	"testing"
	"time"

	"github.com/yyf330/gf/frame/g"
	"github.com/yyf330/gf/net/ghttp"
	"github.com/yyf330/gf/test/gtest"
)

func Test_Template_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := gview.New(gdebug.TestDataPath("template", "basic"))
		p, _ := ports.PopRand()
		s := g.Server(p)
		s.SetView(v)
		s.BindHandler("/", func(r *ghttp.Request) {
			err := r.Response.WriteTpl("index.html", g.Map{
				"name": "john",
			})
			t.Assert(err, nil)
		})
		s.SetDumpRouterMap(false)
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Name:john")
		t.Assert(client.GetContent("/"), "Name:john")
	})
}

func Test_Template_Encode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := gview.New(gdebug.TestDataPath("template", "basic"))
		v.SetAutoEncode(true)
		p, _ := ports.PopRand()
		s := g.Server(p)
		s.SetView(v)
		s.BindHandler("/", func(r *ghttp.Request) {
			err := r.Response.WriteTpl("index.html", g.Map{
				"name": "john",
			})
			t.Assert(err, nil)
		})
		s.SetDumpRouterMap(false)
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Name:john")
		t.Assert(client.GetContent("/"), "Name:john")
	})
}

func Test_Template_Layout1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := gview.New(gdebug.TestDataPath("template", "layout1"))
		p, _ := ports.PopRand()
		s := g.Server(p)
		s.SetView(v)
		s.BindHandler("/layout", func(r *ghttp.Request) {
			err := r.Response.WriteTpl("layout.html", g.Map{
				"mainTpl": "main/main1.html",
			})
			t.Assert(err, nil)
		})
		s.BindHandler("/nil", func(r *ghttp.Request) {
			err := r.Response.WriteTpl("layout.html", nil)
			t.Assert(err, nil)
		})
		s.SetDumpRouterMap(false)
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Not Found")
		t.Assert(client.GetContent("/layout"), "123")
		t.Assert(client.GetContent("/nil"), "123")
	})
}

func Test_Template_Layout2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := gview.New(gdebug.TestDataPath("template", "layout2"))
		p, _ := ports.PopRand()
		s := g.Server(p)
		s.SetView(v)
		s.BindHandler("/main1", func(r *ghttp.Request) {
			err := r.Response.WriteTpl("layout.html", g.Map{
				"mainTpl": "main/main1.html",
			})
			t.Assert(err, nil)
		})
		s.BindHandler("/main2", func(r *ghttp.Request) {
			err := r.Response.WriteTpl("layout.html", g.Map{
				"mainTpl": "main/main2.html",
			})
			t.Assert(err, nil)
		})
		s.BindHandler("/nil", func(r *ghttp.Request) {
			err := r.Response.WriteTpl("layout.html", nil)
			t.Assert(err, nil)
		})
		s.SetDumpRouterMap(false)
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Not Found")
		t.Assert(client.GetContent("/main1"), "a1b")
		t.Assert(client.GetContent("/main2"), "a2b")
		t.Assert(client.GetContent("/nil"), "ab")
	})
}

func Test_Template_XSS(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := gview.New()
		v.SetAutoEncode(true)
		c := "<br>"
		p, _ := ports.PopRand()
		s := g.Server(p)
		s.SetView(v)
		s.BindHandler("/", func(r *ghttp.Request) {
			err := r.Response.WriteTplContent("{{if eq 1 1}}{{.v}}{{end}}", g.Map{
				"v": c,
			})
			t.Assert(err, nil)
		})
		s.SetDumpRouterMap(false)
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), ghtml.Entities(c))
	})
}
