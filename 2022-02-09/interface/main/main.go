package main

import "fmt"

type IPerson interface {
	Eat(name string) string
	Work(things string) string
}

type Student struct {
	content string
}

func (s *Student) Eat(name string) string {
	fmt.Println(s.content)
	s.content = name
	return s.content
}

func (s Student) Work(things string) string {
	fmt.Println(s.content)
	s.content = things
	return s.content
}

func main() {
	s1 := Student{
		content: "在家",
	}
	fmt.Println(s1.Eat("米饭"), s1.Work("上学"))
	fmt.Println("1-----------")

	s2 := Student{
		content: "在家",
	}
	fmt.Println(s2.Work("上学"), s2.Eat("米饭"))
	fmt.Println("2-----------")

	s3 := &Student{
		content: "在家",
	}
	fmt.Println(s3)
	fmt.Println(s3.Eat("米饭"))
	fmt.Println(s3)
	fmt.Println(s3.Work("上学"))
	fmt.Println("3-----------")

	s4 := &Student{
		content: "在家",
	}
	fmt.Println(s4)
	fmt.Println(s4.Work("上学"))
	fmt.Println(s4)
	fmt.Println(s4.Eat("米饭"))
	fmt.Println("4-----------")

	var s5 IPerson = &Student{
		content: "在家",
	}
	fmt.Println(s5.Work("上学"), s5.Eat("米饭"))
	fmt.Println("5-----------")

	// var s6 IPerson = Student{
	// 	content: "在家",
	// }
	// fmt.Println(s6.Work("上学"), s6.Eat("米饭"))
	// fmt.Println("-----------")

}
