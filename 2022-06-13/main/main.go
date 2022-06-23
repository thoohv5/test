package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	B    B
	Name string
	age  int
}

type B struct {
	C string
}

func main() {

	v := &A{
		B: B{
			C: "111",
		},
		Name: "111",
	}

	fmt.Printf("%+v", v)
	fmt.Printf("%#v", v)
	fmt.Println(prettyPrint(v))
	PrettyPrint(v)
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

// print the contents of the obj
func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}
