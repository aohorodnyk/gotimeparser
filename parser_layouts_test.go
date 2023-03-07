package gotimeparser_test

import (
	"fmt"
	"time"

	"github.com/aohorodnyk/gotimeparser"
)

func ExampleParseAllLayouts() {
	rfc3339 := "2022-05-06T03:35:02Z"
	fmt.Println(gotimeparser.ParseAllLayouts(rfc3339))

	rfc3339Nano := "2022-05-06T03:35:02.363368423Z"
	fmt.Println(gotimeparser.ParseAllLayouts(rfc3339Nano))

	rfc1123 := "Fri, 06 May 2022 03:35:02 UTC"
	fmt.Println(gotimeparser.ParseAllLayouts(rfc1123))

	rfc850 := "Friday, 06-May-22 03:35:02 UTC"
	fmt.Println(gotimeparser.ParseAllLayouts(rfc850))

	rfc822 := "06 May 22 03:35 UTC"
	fmt.Println(gotimeparser.ParseAllLayouts(rfc822))

	fmt.Println("=-=-=-=-=-=-=-=-=")

	wrong := "06 of May 22 03:35 in UTC"
	fmt.Println(gotimeparser.ParseAllLayouts(wrong))

	// Output:
	// 2022-05-06 03:35:02 +0000 UTC <nil>
	// 2022-05-06 03:35:02.363368423 +0000 UTC <nil>
	// 2022-05-06 03:35:02 +0000 UTC <nil>
	// 2022-05-06 03:35:02 +0000 UTC <nil>
	// 2022-05-06 03:35:00 +0000 UTC <nil>
	// =-=-=-=-=-=-=-=-=
	// 0001-01-01 00:00:00 +0000 UTC could not parse time: 06 of May 22 03:35 in UTC
}

func ExampleParseLayouts() {
	rfc3339 := "2022-05-06T03:35:02Z"
	fmt.Println(gotimeparser.ParseLayouts([]string{time.RFC3339}, rfc3339))

	rfc3339Nano := "2022-05-06T03:35:02.363368423Z"
	fmt.Println(gotimeparser.ParseLayouts([]string{time.RFC3339, time.RFC3339Nano}, rfc3339Nano))

	rfc1123 := "Fri, 06 May 2022 03:35:02 UTC"
	fmt.Println(gotimeparser.ParseLayouts([]string{time.RFC1123}, rfc1123))

	fmt.Println("=-=-=-=-=-=-=-=-=")

	// Use RFC3339 layout, but support only RFC1123.
	fmt.Println(gotimeparser.ParseLayouts([]string{time.RFC1123}, rfc3339))

	wrong := "06 of May 22 03:35 in UTC"
	fmt.Println(gotimeparser.ParseAllLayouts(wrong))

	// Output:
	// 2022-05-06 03:35:02 +0000 UTC <nil>
	// 2022-05-06 03:35:02.363368423 +0000 UTC <nil>
	// 2022-05-06 03:35:02 +0000 UTC <nil>
	// =-=-=-=-=-=-=-=-=
	// 0001-01-01 00:00:00 +0000 UTC could not parse time: 2022-05-06T03:35:02Z
	// 0001-01-01 00:00:00 +0000 UTC could not parse time: 06 of May 22 03:35 in UTC
}
