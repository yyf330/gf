// Copyright 2020 gf Author(https://github.com/yyf330/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yyf330/gf.

package gview_test

import (
	"github.com/yyf330/gf/debug/gdebug"
	"github.com/yyf330/gf/frame/g"
	"github.com/yyf330/gf/os/gfile"
	"github.com/yyf330/gf/os/gview"
	"github.com/yyf330/gf/test/gtest"
	"testing"
)

func Test_Encode_Parse(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := gview.New()
		v.SetPath(gdebug.TestDataPath("tpl"))
		v.SetAutoEncode(true)
		result, err := v.Parse("encode.tpl", g.Map{
			"title": "<b>my title</b>",
		})
		t.Assert(err, nil)
		t.Assert(result, "<div>&lt;b&gt;my title&lt;/b&gt;</div>")
	})
}

func Test_Encode_ParseContent(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := gview.New()
		tplContent := gfile.GetContents(gdebug.TestDataPath("tpl", "encode.tpl"))
		v.SetAutoEncode(true)
		result, err := v.ParseContent(tplContent, g.Map{
			"title": "<b>my title</b>",
		})
		t.Assert(err, nil)
		t.Assert(result, "<div>&lt;b&gt;my title&lt;/b&gt;</div>")
	})
}
