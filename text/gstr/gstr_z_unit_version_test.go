// Copyright 2019 gf Author(https://github.com/yyf330/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yyf330/gf.

// go test *.go -bench=".*"

package gstr_test

import (
	"testing"

	"github.com/yyf330/gf/test/gtest"
	"github.com/yyf330/gf/text/gstr"
)

func Test_CompareVersion(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gstr.CompareVersion("1", "v0.99"), 1)
		t.AssertEQ(gstr.CompareVersion("v1.0", "v0.99"), 1)
		t.AssertEQ(gstr.CompareVersion("v1.0.1", "v1.1.0"), -1)
		t.AssertEQ(gstr.CompareVersion("1.0.1", "v1.1.0"), -1)
		t.AssertEQ(gstr.CompareVersion("1.0.0", "v0.1.0"), 1)
		t.AssertEQ(gstr.CompareVersion("1.0.0", "v1.0.0"), 0)
	})
}

func Test_CompareVersionGo(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gstr.CompareVersionGo("v1.0.1", "v1.1.0"), -1)
		t.AssertEQ(gstr.CompareVersionGo("1.0.1", "v1.1.0"), -1)
		t.AssertEQ(gstr.CompareVersionGo("1.0.0", "v0.1.0"), 1)
		t.AssertEQ(gstr.CompareVersionGo("1.0.0", "v1.0.0"), 0)
		t.AssertEQ(gstr.CompareVersionGo("v0.0.0-20190626092158-b2ccc519800e", "0.0.0-20190626092158"), 0)
		t.AssertEQ(gstr.CompareVersionGo("v0.0.0-20190626092159-b2ccc519800e", "0.0.0-20190626092158"), 1)
		t.AssertEQ(gstr.CompareVersionGo("v4.20.0+incompatible", "4.20.0"), 0)
		t.AssertEQ(gstr.CompareVersionGo("v4.20.0+incompatible", "4.20.1"), -1)
		// Note that this comparison a < b.
		t.AssertEQ(gstr.CompareVersionGo("v1.12.2-0.20200413154443-b17e3a6804fa", "v1.12.2"), -1)
	})
}
