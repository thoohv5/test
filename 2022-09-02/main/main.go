package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	now := "2022-09-02 10:56:43"

	println(time.Now().String())

	parse, _ := time.Parse("2006-01-02 15:04:05", now)
	fmt.Println(parse.Unix(), parse.String())
	location, _ := time.ParseInLocation("2006-01-02 15:04:05", now, time.Local)
	fmt.Println(location.Unix(), location.String())

	strings.TrimSuffix()

	strings.TrimRight()
}
