// Copyright 2017 gf Author(https://github.com/yyf330/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yyf330/gf.

package gmvc

import "github.com/yyf330/gf/database/gdb"

type (
	M     = Model      // M is alias for Model, just for short write purpose.
	Model = *gdb.Model // Model is alias for *gdb.Model.
)
