// Copyright 2019 gf Author(https://github.com/yyf330/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yyf330/gf.

package gutil_test

import (
	"github.com/yyf330/gf/frame/g"
	"testing"

	"github.com/yyf330/gf/test/gtest"
	"github.com/yyf330/gf/util/gutil"
)

func Test_MapCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := g.Map{
			"k1": "v1",
		}
		m2 := gutil.MapCopy(m1)
		m2["k2"] = "v2"

		t.Assert(m1["k1"], "v1")
		t.Assert(m1["k2"], nil)
		t.Assert(m2["k1"], "v1")
		t.Assert(m2["k2"], "v2")
	})
}

func Test_MapContains(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := g.Map{
			"k1": "v1",
		}
		t.Assert(gutil.MapContains(m1, "k1"), true)
		t.Assert(gutil.MapContains(m1, "K1"), false)
		t.Assert(gutil.MapContains(m1, "k2"), false)
	})
}

func Test_MapMerge(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := g.Map{
			"k1": "v1",
		}
		m2 := g.Map{
			"k2": "v2",
		}
		m3 := g.Map{
			"k3": "v3",
		}
		gutil.MapMerge(m1, m2, m3, nil)
		t.Assert(m1["k1"], "v1")
		t.Assert(m1["k2"], "v2")
		t.Assert(m1["k3"], "v3")
		t.Assert(m2["k1"], nil)
		t.Assert(m3["k1"], nil)
	})
}

func Test_MapMergeCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := g.Map{
			"k1": "v1",
		}
		m2 := g.Map{
			"k2": "v2",
		}
		m3 := g.Map{
			"k3": "v3",
		}
		m := gutil.MapMergeCopy(m1, m2, m3, nil)
		t.Assert(m["k1"], "v1")
		t.Assert(m["k2"], "v2")
		t.Assert(m["k3"], "v3")
		t.Assert(m1["k1"], "v1")
		t.Assert(m1["k2"], nil)
		t.Assert(m2["k1"], nil)
		t.Assert(m3["k1"], nil)
	})
}

func Test_MapPossibleItemByKey(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := g.Map{
			"name":     "guo",
			"NickName": "john",
		}
		k, v := gutil.MapPossibleItemByKey(m, "NAME")
		t.Assert(k, "name")
		t.Assert(v, "guo")

		k, v = gutil.MapPossibleItemByKey(m, "nick name")
		t.Assert(k, "NickName")
		t.Assert(v, "john")

		k, v = gutil.MapPossibleItemByKey(m, "none")
		t.Assert(k, "")
		t.Assert(v, nil)
	})
}

func Test_MapContainsPossibleKey(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := g.Map{
			"name":     "guo",
			"NickName": "john",
		}
		t.Assert(gutil.MapContainsPossibleKey(m, "name"), true)
		t.Assert(gutil.MapContainsPossibleKey(m, "NAME"), true)
		t.Assert(gutil.MapContainsPossibleKey(m, "nickname"), true)
		t.Assert(gutil.MapContainsPossibleKey(m, "nick name"), true)
		t.Assert(gutil.MapContainsPossibleKey(m, "nick_name"), true)
		t.Assert(gutil.MapContainsPossibleKey(m, "nick-name"), true)
		t.Assert(gutil.MapContainsPossibleKey(m, "nick.name"), true)
		t.Assert(gutil.MapContainsPossibleKey(m, "none"), false)
	})
}

func Test_MapOmitEmpty(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := g.Map{
			"k1": "john",
			"e1": "",
			"e2": 0,
			"e3": nil,
			"k2": "smith",
		}
		gutil.MapOmitEmpty(m)
		t.Assert(len(m), 2)
		t.AssertNE(m["k1"], nil)
		t.AssertNE(m["k2"], nil)
	})
}
