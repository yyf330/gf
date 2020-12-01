// Copyright 2017 gf Author(https://github.com/yyf330/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yyf330/gf.

package gdb

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/yyf330/gf/errors/gerror"
	"github.com/yyf330/gf/os/gtime"
	"github.com/yyf330/gf/text/gstr"
	"github.com/yyf330/gf/util/gconv"
	"github.com/yyf330/gf/util/gutil"
	"reflect"
)

// Update does "UPDATE ... " statement for the model.
//
// If the optional parameter <dataAndWhere> is given, the dataAndWhere[0] is the updated data field,
// and dataAndWhere[1:] is treated as where condition fields.
// Also see Model.Data and Model.Where functions.
func (m *Model) Update(dataAndWhere ...interface{}) (result sql.Result, err error) {
	if len(dataAndWhere) > 0 {
		if len(dataAndWhere) > 2 {
			return m.Data(dataAndWhere[0]).Where(dataAndWhere[1], dataAndWhere[2:]...).Update()
		} else if len(dataAndWhere) == 2 {
			return m.Data(dataAndWhere[0]).Where(dataAndWhere[1]).Update()
		} else {
			return m.Data(dataAndWhere[0]).Update()
		}
	}
	defer func() {
		if err == nil {
			m.checkAndRemoveCache()
		}
	}()
	if m.data == nil {
		return nil, errors.New("updating table with empty data")
	}
	var (
		updateData                                    = m.data
		fieldNameCreate                               = m.getSoftFieldNameCreated()
		fieldNameUpdate                               = m.getSoftFieldNameUpdated()
		fieldNameDelete                               = m.getSoftFieldNameDeleted()
		conditionWhere, conditionExtra, conditionArgs = m.formatCondition(false, false)
	)
	// Automatically update the record updating time.
	if !m.unscoped && fieldNameUpdate != "" {
		var (
			refValue = reflect.ValueOf(m.data)
			refKind  = refValue.Kind()
		)
		if refKind == reflect.Ptr {
			refValue = refValue.Elem()
			refKind = refValue.Kind()
		}
		switch refKind {
		case reflect.Map, reflect.Struct:
			dataMap := ConvertDataForTableRecord(m.data)
			gutil.MapDelete(dataMap, fieldNameCreate, fieldNameUpdate, fieldNameDelete)
			if fieldNameUpdate != "" {
				dataMap[fieldNameUpdate] = gtime.Now().String()
			}
			updateData = dataMap
		default:
			updates := gconv.String(m.data)
			if fieldNameUpdate != "" && !gstr.Contains(updates, fieldNameUpdate) {
				updates += fmt.Sprintf(`,%s='%s'`, fieldNameUpdate, gtime.Now().String())
			}
			updateData = updates
		}
	}
	newData, err := m.filterDataForInsertOrUpdate(updateData)
	if err != nil {
		return nil, err
	}
	conditionStr := conditionWhere + conditionExtra
	if !gstr.ContainsI(conditionStr, " WHERE ") {
		return nil, gerror.New("there should be WHERE condition statement for UPDATE operation")
	}
	return m.db.DoUpdate(
		m.getLink(true),
		m.tables,
		newData,
		conditionStr,
		m.mergeArguments(conditionArgs)...,
	)
}
