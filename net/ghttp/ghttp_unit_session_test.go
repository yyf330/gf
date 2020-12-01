// Copyright 2018 gf Author(https://github.com/yyf330/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yyf330/gf.

package ghttp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/yyf330/gf/frame/g"
	"github.com/yyf330/gf/net/ghttp"
	"github.com/yyf330/gf/test/gtest"
)

func Test_Session_Cookie(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/set", func(r *ghttp.Request) {
		r.Session.Set(r.GetString("k"), r.GetString("v"))
	})
	s.BindHandler("/get", func(r *ghttp.Request) {
		r.Response.Write(r.Session.Get(r.GetString("k")))
	})
	s.BindHandler("/remove", func(r *ghttp.Request) {
		r.Session.Remove(r.GetString("k"))
	})
	s.BindHandler("/clear", func(r *ghttp.Request) {
		r.Session.Clear()
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := ghttp.NewClient()
		client.SetBrowserMode(true)
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		r1, e1 := client.Get("/set?k=key1&v=100")
		if r1 != nil {
			defer r1.Close()
		}
		t.Assert(e1, nil)
		t.Assert(r1.ReadAllString(), "")

		t.Assert(client.GetContent("/set?k=key2&v=200"), "")

		t.Assert(client.GetContent("/get?k=key1"), "100")
		t.Assert(client.GetContent("/get?k=key2"), "200")
		t.Assert(client.GetContent("/get?k=key3"), "")
		t.Assert(client.GetContent("/remove?k=key1"), "")
		t.Assert(client.GetContent("/remove?k=key3"), "")
		t.Assert(client.GetContent("/remove?k=key4"), "")
		t.Assert(client.GetContent("/get?k=key1"), "")
		t.Assert(client.GetContent("/get?k=key2"), "200")
		t.Assert(client.GetContent("/clear"), "")
		t.Assert(client.GetContent("/get?k=key2"), "")
	})
}

func Test_Session_Header(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/set", func(r *ghttp.Request) {
		r.Session.Set(r.GetString("k"), r.GetString("v"))
	})
	s.BindHandler("/get", func(r *ghttp.Request) {
		r.Response.Write(r.Session.Get(r.GetString("k")))
	})
	s.BindHandler("/remove", func(r *ghttp.Request) {
		r.Session.Remove(r.GetString("k"))
	})
	s.BindHandler("/clear", func(r *ghttp.Request) {
		r.Session.Clear()
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		response, e1 := client.Get("/set?k=key1&v=100")
		if response != nil {
			defer response.Close()
		}
		sessionId := response.GetCookie(s.GetSessionIdName())
		t.Assert(e1, nil)
		t.AssertNE(sessionId, nil)
		t.Assert(response.ReadAllString(), "")

		client.SetHeader(s.GetSessionIdName(), sessionId)

		t.Assert(client.GetContent("/set?k=key2&v=200"), "")

		t.Assert(client.GetContent("/get?k=key1"), "100")
		t.Assert(client.GetContent("/get?k=key2"), "200")
		t.Assert(client.GetContent("/get?k=key3"), "")
		t.Assert(client.GetContent("/remove?k=key1"), "")
		t.Assert(client.GetContent("/remove?k=key3"), "")
		t.Assert(client.GetContent("/remove?k=key4"), "")
		t.Assert(client.GetContent("/get?k=key1"), "")
		t.Assert(client.GetContent("/get?k=key2"), "200")
		t.Assert(client.GetContent("/clear"), "")
		t.Assert(client.GetContent("/get?k=key2"), "")
	})
}

func Test_Session_StorageFile(t *testing.T) {
	sessionId := ""
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/set", func(r *ghttp.Request) {
		r.Session.Set(r.GetString("k"), r.GetString("v"))
		r.Response.Write(r.GetString("k"), "=", r.GetString("v"))
	})
	s.BindHandler("/get", func(r *ghttp.Request) {
		r.Response.Write(r.Session.Get(r.GetString("k")))
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)

	gtest.C(t, func(t *gtest.T) {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		response, e1 := client.Get("/set?k=key&v=100")
		if response != nil {
			defer response.Close()
		}
		sessionId = response.GetCookie(s.GetSessionIdName())
		t.Assert(e1, nil)
		t.AssertNE(sessionId, nil)
		t.Assert(response.ReadAllString(), "key=100")
	})
	time.Sleep(time.Second)
	gtest.C(t, func(t *gtest.T) {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		client.SetHeader(s.GetSessionIdName(), sessionId)
		t.Assert(client.GetContent("/get?k=key"), "100")
		t.Assert(client.GetContent("/get?k=key1"), "")
	})
}

func Test_Session_Custom_Id(t *testing.T) {
	var (
		sessionId = "1234567890"
		key       = "key"
		value     = "value"
		p, _      = ports.PopRand()
		s         = g.Server(p)
	)
	s.BindHandler("/id", func(r *ghttp.Request) {
		if err := r.Session.SetId(sessionId); err != nil {
			r.Response.WriteExit(err.Error())
		}
		if err := r.Session.Set(key, value); err != nil {
			r.Response.WriteExit(err.Error())
		}
		r.Response.WriteExit(r.Session.Id())
	})
	s.BindHandler("/value", func(r *ghttp.Request) {
		r.Response.WriteExit(r.Session.Get(key))
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)

	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		r, err := client.Get("/id")
		t.Assert(err, nil)
		defer r.Close()
		t.Assert(r.ReadAllString(), sessionId)
		t.Assert(r.GetCookie(s.GetSessionIdName()), sessionId)
	})
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		client.SetHeader(s.GetSessionIdName(), sessionId)
		t.Assert(client.GetContent("/value"), value)
	})
}
