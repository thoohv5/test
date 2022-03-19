package main

import "fmt"

func main() {
	// t1 := Test{name: "初始值"}
	// // t1.setByPointer("1")
	// t1.setByValue("1")
	// t1.displayByPointer()
	// t1.displayByValue()

	t2 := &Test{name: "初始值"}
	t2.setByPointer("1")
	// t2.setByValue("1")
	t2.displayByPointer()
	t2.displayByValue()

}

type Test struct {
	name string
}

func (t *Test) setByPointer(name string) {
	t.name = name
}

func (t *Test) displayByPointer() {
	fmt.Println(t.name)
}

func (t Test) setByValue(name string) {
	t.name = name
}

func (t Test) displayByValue() {
	fmt.Println(t.name)
}
