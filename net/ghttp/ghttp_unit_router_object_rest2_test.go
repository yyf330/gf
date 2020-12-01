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

type ObjectRest2 struct{}

func (o *ObjectRest2) Init(r *ghttp.Request) {
	r.Response.Write("1")
}

func (o *ObjectRest2) Shut(r *ghttp.Request) {
	r.Response.Write("2")
}

func (o *ObjectRest2) Get(r *ghttp.Request) {
	r.Response.Write("Object Get", r.Get("id"))
}

func (o *ObjectRest2) Put(r *ghttp.Request) {
	r.Response.Write("Object Put", r.Get("id"))
}

func (o *ObjectRest2) Post(r *ghttp.Request) {
	r.Response.Write("Object Post", r.Get("id"))
}

func (o *ObjectRest2) Delete(r *ghttp.Request) {
	r.Response.Write("Object Delete", r.Get("id"))
}

func Test_Router_ObjectRest_Id(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindObjectRest("/object/:id", new(ObjectRest2))
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/object/99"), "1Object Get992")
		t.Assert(client.PutContent("/object/99"), "1Object Put992")
		t.Assert(client.PostContent("/object/99"), "1Object Post992")
		t.Assert(client.DeleteContent("/object/99"), "1Object Delete992")
	})
}
