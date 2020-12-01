// Copyright 2019 gf Author(https://github.com/yyf330/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yyf330/gf.

package gerror_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/yyf330/gf/errors/gerror"
	"github.com/yyf330/gf/test/gtest"
)

func nilError() error {
	return nil
}

func Test_Nil(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gerror.New(""), nil)
		t.Assert(gerror.Wrap(nilError(), "test"), nil)
	})
}

func Test_Wrap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := errors.New("1")
		err = gerror.Wrap(err, "2")
		err = gerror.Wrap(err, "3")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "3: 2: 1")
	})

	gtest.C(t, func(t *gtest.T) {
		err := gerror.New("1")
		err = gerror.Wrap(err, "2")
		err = gerror.Wrap(err, "3")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "3: 2: 1")
	})
}

func Test_Cause(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := errors.New("1")
		t.Assert(gerror.Cause(err), err)
	})

	gtest.C(t, func(t *gtest.T) {
		err := errors.New("1")
		err = gerror.Wrap(err, "2")
		err = gerror.Wrap(err, "3")
		t.Assert(gerror.Cause(err), "1")
	})

	gtest.C(t, func(t *gtest.T) {
		err := gerror.New("1")
		t.Assert(gerror.Cause(err), "1")
	})

	gtest.C(t, func(t *gtest.T) {
		err := gerror.New("1")
		err = gerror.Wrap(err, "2")
		err = gerror.Wrap(err, "3")
		t.Assert(gerror.Cause(err), "1")
	})
}

func Test_Format(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := errors.New("1")
		err = gerror.Wrap(err, "2")
		err = gerror.Wrap(err, "3")
		t.AssertNE(err, nil)
		t.Assert(fmt.Sprintf("%s", err), "3: 2: 1")
		t.Assert(fmt.Sprintf("%v", err), "3: 2: 1")
	})

	gtest.C(t, func(t *gtest.T) {
		err := gerror.New("1")
		err = gerror.Wrap(err, "2")
		err = gerror.Wrap(err, "3")
		t.AssertNE(err, nil)
		t.Assert(fmt.Sprintf("%s", err), "3: 2: 1")
		t.Assert(fmt.Sprintf("%v", err), "3: 2: 1")
	})

	gtest.C(t, func(t *gtest.T) {
		err := gerror.New("1")
		err = gerror.Wrap(err, "2")
		err = gerror.Wrap(err, "3")
		t.AssertNE(err, nil)
		t.Assert(fmt.Sprintf("%-s", err), "3")
		t.Assert(fmt.Sprintf("%-v", err), "3")
	})
}

func Test_Stack(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := errors.New("1")
		t.Assert(fmt.Sprintf("%+v", err), "1")
	})

	gtest.C(t, func(t *gtest.T) {
		err := errors.New("1")
		err = gerror.Wrap(err, "2")
		err = gerror.Wrap(err, "3")
		t.AssertNE(err, nil)
		//fmt.Printf("%+v", err)
	})

	gtest.C(t, func(t *gtest.T) {
		err := gerror.New("1")
		t.AssertNE(fmt.Sprintf("%+v", err), "1")
		//fmt.Printf("%+v", err)
	})

	gtest.C(t, func(t *gtest.T) {
		err := gerror.New("1")
		err = gerror.Wrap(err, "2")
		err = gerror.Wrap(err, "3")
		t.AssertNE(err, nil)
		//fmt.Printf("%+v", err)
	})
}

func Test_Current(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := errors.New("1")
		err = gerror.Wrap(err, "2")
		err = gerror.Wrap(err, "3")
		t.Assert(err.Error(), "3: 2: 1")
		t.Assert(gerror.Current(err).Error(), "3")
	})
}

func Test_Next(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := errors.New("1")
		err = gerror.Wrap(err, "2")
		err = gerror.Wrap(err, "3")
		t.Assert(err.Error(), "3: 2: 1")

		err = gerror.Next(err)
		t.Assert(err.Error(), "2: 1")

		err = gerror.Next(err)
		t.Assert(err.Error(), "1")

		err = gerror.Next(err)
		t.Assert(err, nil)
	})
}
