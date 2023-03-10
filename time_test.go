package gotimeparser_test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aohorodnyk/gotimeparser"
)

func ExampleTime_UnmarshalJSON() {
	type jsonType struct {
		NumSeconds      gotimeparser.Time `json:"num_seconds"`
		NumMilliseconds gotimeparser.Time `json:"num_milliseconds"`
		NumMicroseconds gotimeparser.Time `json:"num_microseconds"`
		NumNanoseconds  gotimeparser.Time `json:"num_nanoseconds"`

		HexSeconds gotimeparser.Time `json:"hex_seconds"`

		StrSeconds      gotimeparser.Time `json:"str_seconds"`
		StrMilliseconds gotimeparser.Time `json:"str_milliseconds"`
		StrMicroseconds gotimeparser.Time `json:"str_microseconds"`
		StrNanoseconds  gotimeparser.Time `json:"str_nanoseconds"`

		StrRFC3339     gotimeparser.Time `json:"str_rfc3339"`
		StrRFC3339Nano gotimeparser.Time `json:"str_rfc3339_nano"`
		StrRFC1123     gotimeparser.Time `json:"str_rfc1123"`
		StrRFC850      gotimeparser.Time `json:"str_rfc850"`
		StrRFC822      gotimeparser.Time `json:"str_rfc822"`
	}

	jsonStr := `{
    "num_seconds": 1651808102,
    "num_milliseconds": 1651808102363,
    "num_microseconds": 1651808102363368,
    "num_nanoseconds": 1651808102363368423,

    "hex_seconds": "0x62749766",

    "str_seconds": "1651808102",
    "str_milliseconds": "1651808102363",
    "str_microseconds": "1651808102363368",
    "str_nanoseconds": "1651808102363368423",

    "str_rfc3339": "2022-05-06T03:35:02Z",
    "str_rfc3339_nano": "2022-05-06T03:35:02.363368423Z",
    "str_rfc1123": "Fri, 06 May 2022 03:35:02 UTC",
    "str_rfc850": "Friday, 06-May-22 03:35:02 UTC",
    "str_rfc822": "06 May 22 03:35 UTC"
  }`

	var parsed jsonType

	err := json.Unmarshal([]byte(jsonStr), &parsed)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("NumSeconds:", parsed.NumSeconds)
	fmt.Println("HexSeconds:", parsed.HexSeconds)
	fmt.Println("NumMicroseconds:", parsed.NumMicroseconds)
	fmt.Println("NumMilliseconds:", parsed.NumMilliseconds)
	fmt.Println("NumNanoseconds:", parsed.NumNanoseconds)
	fmt.Println("StrMicroseconds:", parsed.StrMicroseconds)
	fmt.Println("StrMilliseconds:", parsed.StrMilliseconds)
	fmt.Println("StrNanoseconds:", parsed.StrNanoseconds)
	fmt.Println("StrSeconds:", parsed.StrSeconds)
	fmt.Println("StrRFC1123:", parsed.StrRFC1123)
	fmt.Println("StrRFC3339:", parsed.StrRFC3339)
	fmt.Println("StrRFC3339Nano:", parsed.StrRFC3339Nano)
	fmt.Println("StrRFC822:", parsed.StrRFC822)
	fmt.Println("StrRFC850:", parsed.StrRFC850)

	// Covert time.Time to gotimeparser.Time.
	gotime := time.Date(2023, 3, 7, 19, 33, 23, 523, time.UTC)
	fmt.Println("gotimeparser.Time:", gotimeparser.FromGoTime(gotime))

	// Output:
	// NumSeconds: 2022-05-05 20:35:02 -0700 PDT
	// HexSeconds: 2022-05-05 20:35:02 -0700 PDT
	// NumMicroseconds: 2022-05-05 20:35:02.363368 -0700 PDT
	// NumMilliseconds: 2022-05-05 20:35:02.363 -0700 PDT
	// NumNanoseconds: 2022-05-05 20:35:02.363368423 -0700 PDT
	// StrMicroseconds: 2022-05-05 20:35:02.363368 -0700 PDT
	// StrMilliseconds: 2022-05-05 20:35:02.363 -0700 PDT
	// StrNanoseconds: 2022-05-05 20:35:02.363368423 -0700 PDT
	// StrSeconds: 2022-05-05 20:35:02 -0700 PDT
	// StrRFC1123: 2022-05-06 03:35:02 +0000 UTC
	// StrRFC3339: 2022-05-06 03:35:02 +0000 UTC
	// StrRFC3339Nano: 2022-05-06 03:35:02.363368423 +0000 UTC
	// StrRFC822: 2022-05-06 03:35:00 +0000 UTC
	// StrRFC850: 2022-05-06 03:35:02 +0000 UTC
	// gotimeparser.Time: 2023-03-07 19:33:23.000000523 +0000 UTC
}
