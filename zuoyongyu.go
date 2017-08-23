package main

import "fmt"

/*
函数内定义的变量称为局部变量
函数外定义的变量称为全局变量
函数定义中的变量称为形式参数
 */

func sum(a, b int) int {
	return a + b
}

var name string = "hello"               //全局变量

func main()  {
	fmt.Println(name)
	name = "world"
	fmt.Println(name)


	var a, b, c int			//局部变量
	a = 1
	b = 2
	c = a + b
	fmt.Println(a, b, c)
	var d = 0
	d = sum(a, b)			// 形式参数
	fmt.Println(d)
}
