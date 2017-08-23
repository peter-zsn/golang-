package main

import (
	"fmt"
)
/*
类型转换
type(value)
 */

func main() {
	var a int = 7
	var b int= 2
	var c float32
	c = float32(a) / float32(b)
	d := a/b
	fmt.Println(c)
	fmt.Println(d)

}
