package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//fmt.Println("hello world")
	//pwd, _ := os.Getwd()
	//fmt.Printf("this is %s\n", pwd)
	WorkDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(WorkDir, err)

}
