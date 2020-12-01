// Copyright 2019 gf Author(https://github.com/yyf330/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yyf330/gf.

package gdb

import (
	"github.com/yyf330/gf/test/gtest"
	"testing"
)

func Test_Oracle_parseSql(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		o := new(DriverOracle)
		sql := `UPDATE user SET name='john'`
		newSql := o.parseSql(sql)
		t.Assert(newSql, sql)
	})

	gtest.C(t, func(t *gtest.T) {
		o := new(DriverOracle)
		sql := `SELECT * FROM user`
		newSql := o.parseSql(sql)
		t.Assert(newSql, sql)
	})

	gtest.C(t, func(t *gtest.T) {
		o := new(DriverOracle)
		sql := `SELECT * FROM user LIMIT 0, 10`
		newSql := o.parseSql(sql)
		t.Assert(newSql, `SELECT * FROM (SELECT GFORM.*, ROWNUM ROWNUM_ FROM (SELECT  * FROM user ) GFORM WHERE ROWNUM <= 10) WHERE ROWNUM_ >= 0`)
	})

	gtest.C(t, func(t *gtest.T) {
		o := new(DriverOracle)
		sql := `SELECT * FROM user LIMIT 1`
		newSql := o.parseSql(sql)
		t.Assert(newSql, `SELECT * FROM (SELECT GFORM.*, ROWNUM ROWNUM_ FROM (SELECT  * FROM user ) GFORM WHERE ROWNUM <= 1) WHERE ROWNUM_ >= 0`)
	})

	gtest.C(t, func(t *gtest.T) {
		o := new(DriverOracle)
		sql := `SELECT ENAME FROM USER_INFO WHERE ID=2 LIMIT 1`
		newSql := o.parseSql(sql)
		t.Assert(newSql, `SELECT * FROM (SELECT GFORM.*, ROWNUM ROWNUM_ FROM (SELECT  ENAME FROM USER_INFO WHERE ID=2 ) GFORM WHERE ROWNUM <= 1) WHERE ROWNUM_ >= 0`)
	})
}
