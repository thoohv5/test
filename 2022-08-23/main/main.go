package main

import (
	"fmt"
	"os"
)

func main() {

	os.MkdirAll("/Users/thooh/Sites/github.com/thoohv5/test/test", 0666)

	os.Chmod("/Users/thooh/Sites/github.com/thoohv5/test/test", 0666)

	file, err := os.OpenFile("/Users/thooh/Sites/github.com/thoohv5/test/test.txt", os.O_CREATE, 0666)
	if nil != err {
		panic(err)
	}

	err = file.Chmod(0666)
	if nil != err {
		panic(err)
	}

	fileInfo, err := file.Stat()
	if nil != err {
		panic(err)
	}

	fmt.Println(fileInfo)

}
