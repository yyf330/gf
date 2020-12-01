// Copyright 2018 gf Author(https://github.com/yyf330/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yyf330/gf.

package gtest_test

import (
	"testing"

	"github.com/yyf330/gf/test/gtest"
)

func TestC(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(1, 1)
		t.AssertNE(1, 0)
		t.AssertEQ(float32(123.456), float32(123.456))
		t.AssertEQ(float64(123.456), float64(123.456))
	})
}

func TestCase(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(1, 1)
		t.AssertNE(1, 0)
		t.AssertEQ(float32(123.456), float32(123.456))
		t.AssertEQ(float64(123.456), float64(123.456))
	})
}
