package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {

	// cmd := exec.Command("rm", "-rf", "/tmp/*")

	// fmt.Println(cmd.Output())

	fmt.Println(exec.LookPath("go"))

	fmt.Println(os.Getwd())
	fmt.Println(os.Args[0])
	fmt.Println(exec.LookPath(os.Args[0]))
	fmt.Println(filepath.Abs(func() string {
		s, _ := exec.LookPath(os.Args[0])
		return s
	}()))

}
