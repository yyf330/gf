// Copyright 2017 gf Author(https://github.com/yyf330/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yyf330/gf.

package gins

import (
	"fmt"
	"github.com/yyf330/gf/debug/gdebug"
	"github.com/yyf330/gf/os/gcfg"
	"testing"

	"github.com/yyf330/gf/os/gfile"
	"github.com/yyf330/gf/os/gtime"
	"github.com/yyf330/gf/test/gtest"
)

func Test_View(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertNE(View(), nil)
		b, e := View().ParseContent(`{{"我是中国人" | substr 2 -1}}`, nil)
		t.Assert(e, nil)
		t.Assert(b, "中国人")
	})
	gtest.C(t, func(t *gtest.T) {
		tpl := "t.tpl"
		err := gfile.PutContents(tpl, `{{"我是中国人" | substr 2 -1}}`)
		t.Assert(err, nil)
		defer gfile.Remove(tpl)

		b, e := View().Parse("t.tpl", nil)
		t.Assert(e, nil)
		t.Assert(b, "中国人")
	})
	gtest.C(t, func(t *gtest.T) {
		path := fmt.Sprintf(`%s/%d`, gfile.TempDir(), gtime.TimestampNano())
		tpl := fmt.Sprintf(`%s/%s`, path, "t.tpl")
		err := gfile.PutContents(tpl, `{{"我是中国人" | substr 2 -1}}`)
		t.Assert(err, nil)
		defer gfile.Remove(tpl)
		err = View().AddPath(path)
		t.Assert(err, nil)

		b, e := View().Parse("t.tpl", nil)
		t.Assert(e, nil)
		t.Assert(b, "中国人")
	})
}

func Test_View_Config(t *testing.T) {
	// view1 test1
	gtest.C(t, func(t *gtest.T) {
		dirPath := gdebug.TestDataPath("view1")
		gcfg.SetContent(gfile.GetContents(gfile.Join(dirPath, "config.toml")))
		defer gcfg.ClearContent()
		defer instances.Clear()

		view := View("test1")
		t.AssertNE(view, nil)
		err := view.AddPath(dirPath)
		t.Assert(err, nil)

		str := `hello ${.name},version:${.version}`
		view.Assigns(map[string]interface{}{"version": "1.9.0"})
		result, err := view.ParseContent(str, nil)
		t.Assert(err, nil)
		t.Assert(result, "hello test1,version:1.9.0")

		result, err = view.ParseDefault()
		t.Assert(err, nil)
		t.Assert(result, "test1:test1")
	})
	// view1 test2
	gtest.C(t, func(t *gtest.T) {
		dirPath := gdebug.TestDataPath("view1")
		gcfg.SetContent(gfile.GetContents(gfile.Join(dirPath, "config.toml")))
		defer gcfg.ClearContent()
		defer instances.Clear()

		view := View("test2")
		t.AssertNE(view, nil)
		err := view.AddPath(dirPath)
		t.Assert(err, nil)

		str := `hello #{.name},version:#{.version}`
		view.Assigns(map[string]interface{}{"version": "1.9.0"})
		result, err := view.ParseContent(str, nil)
		t.Assert(err, nil)
		t.Assert(result, "hello test2,version:1.9.0")

		result, err = view.ParseDefault()
		t.Assert(err, nil)
		t.Assert(result, "test2:test2")
	})
	// view2
	gtest.C(t, func(t *gtest.T) {
		dirPath := gdebug.TestDataPath("view2")
		gcfg.SetContent(gfile.GetContents(gfile.Join(dirPath, "config.toml")))
		defer gcfg.ClearContent()
		defer instances.Clear()

		view := View()
		t.AssertNE(view, nil)
		err := view.AddPath(dirPath)
		t.Assert(err, nil)

		str := `hello {.name},version:{.version}`
		view.Assigns(map[string]interface{}{"version": "1.9.0"})
		result, err := view.ParseContent(str, nil)
		t.Assert(err, nil)
		t.Assert(result, "hello test,version:1.9.0")

		result, err = view.ParseDefault()
		t.Assert(err, nil)
		t.Assert(result, "test:test")
	})
	// view2
	gtest.C(t, func(t *gtest.T) {
		dirPath := gdebug.TestDataPath("view2")
		gcfg.SetContent(gfile.GetContents(gfile.Join(dirPath, "config.toml")))
		defer gcfg.ClearContent()
		defer instances.Clear()

		view := View("test100")
		t.AssertNE(view, nil)
		err := view.AddPath(dirPath)
		t.Assert(err, nil)

		str := `hello {.name},version:{.version}`
		view.Assigns(map[string]interface{}{"version": "1.9.0"})
		result, err := view.ParseContent(str, nil)
		t.Assert(err, nil)
		t.Assert(result, "hello test,version:1.9.0")

		result, err = view.ParseDefault()
		t.Assert(err, nil)
		t.Assert(result, "test:test")
	})
}
