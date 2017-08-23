package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world")
	pwd, _ := os.Getwd()
	fmt.Printf("this is %s\n", pwd)
}
