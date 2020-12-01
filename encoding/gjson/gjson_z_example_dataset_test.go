// Copyright 2017 gf Author(https://github.com/yyf330/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yyf330/gf.

package gjson_test

import (
	"fmt"
	"github.com/yyf330/gf/encoding/gjson"
)

func Example_dataSetCreate1() {
	j := gjson.New(nil)
	j.Set("name", "John")
	j.Set("score", 99.5)
	fmt.Printf(
		"Name: %s, Score: %v\n",
		j.GetString("name"),
		j.GetFloat32("score"),
	)
	fmt.Println(j.MustToJsonString())

	// Output:
	// Name: John, Score: 99.5
	// {"name":"John","score":99.5}
}

func Example_dataSetCreate2() {
	j := gjson.New(nil)
	for i := 0; i < 5; i++ {
		j.Set(fmt.Sprintf(`%d.id`, i), i)
		j.Set(fmt.Sprintf(`%d.name`, i), fmt.Sprintf(`student-%d`, i))
	}
	fmt.Println(j.MustToJsonString())

	// Output:
	// [{"id":0,"name":"student-0"},{"id":1,"name":"student-1"},{"id":2,"name":"student-2"},{"id":3,"name":"student-3"},{"id":4,"name":"student-4"}]
}

func Example_dataSetRuntimeEdit() {
	data :=
		`{
        "users" : {
            "count" : 2,
            "list"  : [
                {"name" : "Ming", "score" : 60},
                {"name" : "John", "score" : 59}
            ]
        }
    }`
	if j, err := gjson.DecodeToJson(data); err != nil {
		panic(err)
	} else {
		j.Set("users.list.1.score", 100)
		fmt.Println("John Score:", j.GetFloat32("users.list.1.score"))
		fmt.Println(j.MustToJsonString())
	}
	// Output:
	// John Score: 100
	// {"users":{"count":2,"list":[{"name":"Ming","score":60},{"name":"John","score":100}]}}
}
