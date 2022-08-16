package main

import (
	"encoding/json"
	"fmt"
)

type Extra struct {
	E int
	F string
}

type Test struct {
	A      string `json:"a"`
	B      int    `json:"b"`
	c      interface{}
	D      float64 `json:"d,omitempty"`
	*Extra `json:"prefile,omitempty"`
}

func main() {

	bd, err := json.Marshal(&Test{
		A: "1",
		B: 2,
		c: 3,
	})
	fmt.Println(string(bd), err)

}
