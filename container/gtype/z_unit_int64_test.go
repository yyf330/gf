// Copyright 2018 gf Author(https://github.com/yyf330/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yyf330/gf.

package gtype_test

import (
	"github.com/yyf330/gf/container/gtype"
	"github.com/yyf330/gf/internal/json"
	"github.com/yyf330/gf/test/gtest"
	"github.com/yyf330/gf/util/gconv"
	"math"
	"sync"
	"testing"
)

func Test_Int64(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var wg sync.WaitGroup
		addTimes := 1000
		i := gtype.NewInt64(0)
		iClone := i.Clone()
		t.AssertEQ(iClone.Set(1), int64(0))
		t.AssertEQ(iClone.Val(), int64(1))
		for index := 0; index < addTimes; index++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				i.Add(1)
			}()
		}
		wg.Wait()
		t.AssertEQ(int64(addTimes), i.Val())

		//空参测试
		i1 := gtype.NewInt64()
		t.AssertEQ(i1.Val(), int64(0))
	})
}

func Test_Int64_JSON(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		i := gtype.NewInt64(math.MaxInt64)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := gtype.NewInt64()
		err := json.Unmarshal(b2, &i2)
		t.Assert(err, nil)
		t.Assert(i2.Val(), i)
	})
}

func Test_Int64_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *gtype.Int64
	}
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "123",
		}, &v)
		t.Assert(err, nil)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.Val(), "123")
	})
}
