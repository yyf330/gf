// Copyright 2018 gf Author(https://github.com/yyf330/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yyf330/gf.

package gtype_test

import (
	"github.com/yyf330/gf/internal/json"
	"github.com/yyf330/gf/util/gconv"
	"sync"
	"testing"

	"github.com/yyf330/gf/container/gtype"
	"github.com/yyf330/gf/test/gtest"
)

func Test_Byte(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var wg sync.WaitGroup
		addTimes := 127
		i := gtype.NewByte(byte(0))
		iClone := i.Clone()
		t.AssertEQ(iClone.Set(byte(1)), byte(0))
		t.AssertEQ(iClone.Val(), byte(1))
		for index := 0; index < addTimes; index++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				i.Add(1)
			}()
		}
		wg.Wait()
		t.AssertEQ(byte(addTimes), i.Val())

		//空参测试
		i1 := gtype.NewByte()
		t.AssertEQ(i1.Val(), byte(0))
	})
}

func Test_Byte_JSON(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		i := gtype.NewByte(49)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)
	})
	// Unmarshal
	gtest.C(t, func(t *gtest.T) {
		var err error
		i := gtype.NewByte()
		err = json.Unmarshal([]byte("49"), &i)
		t.Assert(err, nil)
		t.Assert(i.Val(), "49")
	})
}

func Test_Byte_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *gtype.Byte
	}
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "2",
		}, &v)
		t.Assert(err, nil)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.Val(), "2")
	})
}
