package main

import "fmt"

var a int = 10		// var 一般声明全局变量
var he, wo = "hello", "world"
var x, y int		// 全局变量 没赋值的话为默认值
var (		//这种写法一般用于全局变量
	q int
	w bool // 默认为false
)

//全局变量可以只声明不使用， 局部变量 必须使用

func main()  {
	var b = 5
	c := 15			// := 赋值， 不需要声明， 函数内部，适用于局部变量(函数体内部)
	fmt.Println(a, b, c)
	e, f, g := 1, 2, 3	// 多个值赋值
	fmt.Println(e, f, g)
	fmt.Println(he, wo)
	var name1, name2, name3 int		// 这种声明 后面必须加类型
	name3, name1, name2 = 4, 5, 6
	fmt.Println(name1, name2, name3)

	//int
	//fmt.Println(q, w)
	x = 1					// =赋值的话， 必须声明
	y = 2
	fmt.Println(x, y)

	// str and char
	str := "hello"   // 字符串是常量  必能更改  str[0] = 'c' 编译不通过,  字符打印为2进制
	fmt.Println("hahah", str)
	fmt.Printf("nihao %s\n", str)			// 打印 \n 换行
	char := 'c'
	fmt.Printf("字符为%c\n", char)		//打印字符时，可以用%c

	//值类型 int float bool string 直接指向内存
	// i :=1 j:=2  i = j   进行的是内存的copy 可以通过&i 获取内存地址(指针)
	i :=1
	j :=2
	fmt.Println(i, j, &i, &j)
	i = j
	fmt.Println(i, j, &i, &j)

}