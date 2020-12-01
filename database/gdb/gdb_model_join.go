// Copyright 2017 gf Author(https://github.com/yyf330/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yyf330/gf.

package gdb

import (
	"fmt"
	"github.com/yyf330/gf/text/gstr"
)

// isSubQuery checks and returns whether given string a sub-query sql string.
func isSubQuery(s string) bool {
	s = gstr.TrimLeft(s)
	if p := gstr.Pos(s, " "); p != -1 {
		if gstr.Equal(s[:p], "select") {
			return true
		}
	}
	return false
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (m *Model) LeftJoin(table ...string) *Model {
	var (
		model   = m.getModel()
		joinStr = ""
	)
	if len(table) > 0 {
		if isSubQuery(table[0]) {
			joinStr = "(" + table[0] + ")"
		} else {
			joinStr = m.db.QuotePrefixTableName(table[0])
		}
	}
	if len(table) > 2 {
		model.tables += fmt.Sprintf(
			" LEFT JOIN %s AS %s ON (%s)",
			joinStr, m.db.QuoteWord(table[1]), table[2],
		)
	} else if len(table) == 2 {
		model.tables += fmt.Sprintf(
			" LEFT JOIN %s ON (%s)",
			joinStr, table[1],
		)
	} else if len(table) == 1 {
		model.tables += fmt.Sprintf(
			" LEFT JOIN %s",
			joinStr,
		)
	}
	return model
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (m *Model) RightJoin(table ...string) *Model {
	var (
		model   = m.getModel()
		joinStr = ""
	)
	if len(table) > 0 {
		if isSubQuery(table[0]) {
			joinStr = "(" + table[0] + ")"
		} else {
			joinStr = m.db.QuotePrefixTableName(table[0])
		}
	}
	if len(table) > 2 {
		model.tables += fmt.Sprintf(
			" RIGHT JOIN %s AS %s ON (%s)",
			joinStr, m.db.QuoteWord(table[1]), table[2],
		)
	} else if len(table) == 2 {
		model.tables += fmt.Sprintf(
			" RIGHT JOIN %s ON (%s)",
			joinStr, table[1],
		)
	} else if len(table) == 1 {
		model.tables += fmt.Sprintf(
			" RIGHT JOIN %s",
			joinStr,
		)
	}
	return model
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (m *Model) InnerJoin(table ...string) *Model {
	var (
		model   = m.getModel()
		joinStr = ""
	)
	if len(table) > 0 {
		if isSubQuery(table[0]) {
			joinStr = "(" + table[0] + ")"
		} else {
			joinStr = m.db.QuotePrefixTableName(table[0])
		}
	}
	if len(table) > 2 {
		model.tables += fmt.Sprintf(
			" INNER JOIN %s AS %s ON (%s)",
			joinStr, m.db.QuoteWord(table[1]), table[2],
		)
	} else if len(table) == 2 {
		model.tables += fmt.Sprintf(
			" INNER JOIN %s ON (%s)",
			joinStr, table[1],
		)
	} else if len(table) == 1 {
		model.tables += fmt.Sprintf(
			" INNER JOIN %s",
			joinStr,
		)
	}
	return model
}
