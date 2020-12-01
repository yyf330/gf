// Copyright 2017 gf Author(https://github.com/yyf330/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yyf330/gf.

package driver

import (
	"database/sql"
	"github.com/yyf330/gf/database/gdb"
	"github.com/yyf330/gf/frame/g"
	"github.com/yyf330/gf/os/gtime"
)

// MyDriver is a custom database driver, which is used for testing only.
// For simplifying the unit testing case purpose, MyDriver struct inherits the mysql driver
// gdb.DriverMysql and overwrites its functions DoQuery and DoExec.
// So if there's any sql execution, it goes through MyDriver.DoQuery/MyDriver.DoExec firstly
// and then gdb.DriverMysql.DoQuery/gdb.DriverMysql.DoExec.
// You can call it sql "HOOK" or "HiJack" as your will.
type MyDriver struct {
	*gdb.DriverMysql
}

var (
	// customDriverName is my driver name, which is used for registering.
	customDriverName = "MyDriver"
)

func init() {
	// It here registers my custom driver in package initialization function "init".
	// You can later use this type in the database configuration.
	if err := gdb.Register(customDriverName, &MyDriver{}); err != nil {
		panic(err)
	}
}

// New creates and returns a database object for mysql.
// It implements the interface of gdb.Driver for extra database driver installation.
func (d *MyDriver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error) {
	return &MyDriver{
		&gdb.DriverMysql{
			Core: core,
		},
	}, nil
}

// DoQuery commits the sql string and its arguments to underlying driver
// through given link object and returns the execution result.
func (d *MyDriver) DoQuery(link gdb.Link, sql string, args ...interface{}) (rows *sql.Rows, err error) {
	tsMilli := gtime.TimestampMilli()
	rows, err = d.DriverMysql.DoQuery(link, sql, args...)
	if _, err := d.DriverMysql.InsertIgnore("monitor", g.Map{
		"sql":   gdb.FormatSqlWithArgs(sql, args),
		"cost":  gtime.TimestampMilli() - tsMilli,
		"time":  gtime.Now(),
		"error": err.Error(),
	}); err != nil {
		panic(err)
	}
	return
}

// DoExec commits the query string and its arguments to underlying driver
// through given link object and returns the execution result.
func (d *MyDriver) DoExec(link gdb.Link, sql string, args ...interface{}) (result sql.Result, err error) {
	tsMilli := gtime.TimestampMilli()
	result, err = d.DriverMysql.DoExec(link, sql, args...)
	if _, err := d.DriverMysql.InsertIgnore("monitor", g.Map{
		"sql":   gdb.FormatSqlWithArgs(sql, args),
		"cost":  gtime.TimestampMilli() - tsMilli,
		"time":  gtime.Now(),
		"error": err.Error(),
	}); err != nil {
		panic(err)
	}
	return
}
