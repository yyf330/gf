// Copyright 2017-2018 gf Author(https://github.com/yyf330/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yyf330/gf.

package gfile_test

import (
	"github.com/yyf330/gf/debug/gdebug"
	"testing"

	"github.com/yyf330/gf/os/gfile"
	"github.com/yyf330/gf/test/gtest"
)

func Test_NotFound(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		teatFile := gfile.Dir(gdebug.CallerFilePath()) + gfile.Separator + "testdata/readline/error.log"
		callback := func(line string) {
		}
		err := gfile.ReadLines(teatFile, callback)
		t.AssertNE(err, nil)
	})
}

func Test_ReadLines(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expectList := []string{"a", "b", "c", "d", "e"}

		getList := make([]string, 0)
		callback := func(line string) {
			getList = append(getList, line)
		}

		teatFile := gfile.Dir(gdebug.CallerFilePath()) + gfile.Separator + "testdata/readline/file.log"
		err := gfile.ReadLines(teatFile, callback)

		t.AssertEQ(getList, expectList)
		t.AssertEQ(err, nil)
	})
}

func Test_ReadByteLines(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expectList := [][]byte{[]byte("a"), []byte("b"), []byte("c"), []byte("d"), []byte("e")}

		getList := make([][]byte, 0)
		callback := func(line []byte) {
			getList = append(getList, line)
		}

		teatFile := gfile.Dir(gdebug.CallerFilePath()) + gfile.Separator + "testdata/readline/file.log"
		err := gfile.ReadByteLines(teatFile, callback)

		t.AssertEQ(getList, expectList)
		t.AssertEQ(err, nil)
	})
}
