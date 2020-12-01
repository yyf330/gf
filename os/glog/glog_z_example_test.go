// Copyright 2020 gf Author(https://github.com/yyf330/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yyf330/gf.

package glog_test

import (
	"context"
	"github.com/yyf330/gf/frame/g"
)

func Example_context() {
	ctx := context.WithValue(context.Background(), "Trace-Id", "123456789")
	g.Log().Ctx(ctx).Error("runtime error")

	// May Output:
	// 2020-06-08 20:17:03.630 [ERRO] {Trace-Id: 123456789} runtime error
	// Stack:
	// ...
}
