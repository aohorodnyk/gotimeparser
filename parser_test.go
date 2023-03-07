package gotimeparser_test

import (
	"fmt"

	"github.com/aohorodnyk/gotimeparser"
)

func ExampleParse() {
	dateTime := "1651808102"
	fmt.Println(gotimeparser.Parse(dateTime))

	dateTime = `"1651808102"`
	fmt.Println(gotimeparser.Parse(dateTime))

	dateTime = `'1651808102'`
	fmt.Println(gotimeparser.Parse(dateTime))

	dateTime = "'2022-05-06T03:35:02Z'"
	fmt.Println(gotimeparser.Parse(dateTime))

	dateTime = `"Fri, 06 May 2022 03:35:02 UTC"`
	fmt.Println(gotimeparser.Parse(dateTime))

	fmt.Println("=-=-=-=-=-=-=")

	dateTime = ``
	fmt.Println(gotimeparser.Parse(dateTime))

	dateTime = `Friday on 06 May 2022 03:35:02 in UTC`
	fmt.Println(gotimeparser.Parse(dateTime))

	// Output:
	// 2022-05-05 20:35:02 -0700 PDT <nil>
	// 2022-05-05 20:35:02 -0700 PDT <nil>
	// 2022-05-05 20:35:02 -0700 PDT <nil>
	// 2022-05-06 03:35:02 +0000 UTC <nil>
	// 2022-05-06 03:35:02 +0000 UTC <nil>
	// =-=-=-=-=-=-=
	// 0001-01-01 00:00:00 +0000 UTC Unsupported time format for string: ''
	// 0001-01-01 00:00:00 +0000 UTC Unsupported time format for string: 'Friday on 06 May 2022 03:35:02 in UTC'
}
