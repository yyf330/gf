// Copyright 2017 gf Author(https://github.com/yyf330/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yyf330/gf.

package gtcp_test

import (
	"github.com/yyf330/gf/container/garray"
)

var (
	ports = garray.NewIntArray(true)
)

func init() {
	for i := 9000; i < 10000; i++ {
		ports.Append(i)
	}
}
