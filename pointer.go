package main

import (
	"fmt"
)

/*
& 取地址符
指针  指向任何一个值的内存地址它指向那个值的内存地址
和c一样 也存在空指针 值为nil 和 null None nil NULL 一样


声明
var p_name *type
 */

var A *int
var B *float32

var D [2]*int
// 多维数组指针
var C [2][2]*int

func set_arry_rel(arry [3] int)  {
	arry[0] = 10
}
func set_arry_P(arry [3]*int)  {
	*arry[0] = 20
}

func set_arry_arry1(arry [3] int) [3]int {
	//var result [3]int
	//result = arry
	//result[0] = 100
	//return result
	arry[0] = 100
	return arry
}

func main() {
	a := 1
	fmt.Printf("a 的地址是 %x \n", &a)

	// 指向一个值的指针
	var a_P *int
	a_P = &a
	fmt.Printf("a 的地址是 %x \n", a_P)
	fmt.Printf("p_a 指向的值是 %d\n", *a_P)

	// 空指针  *null_p 取值时会报错
	var null_P *int
	fmt.Printf("null_P is %x\n", null_P)

	// 判断是不是空指针
	if(null_P == nil){
		fmt.Printf("null_p is 空指针")
	}else{
		fmt.Printf("null_p is 不是空指针")
	}

	arry := [3] int{1,2,3}
	var arry_p [3]* int
	for i :=0; i < 3; i++{
		arry_p[i] = &arry[i]
	}
	for i :=0; i < 3; i++{
		fmt.Printf("arry_p[%d] 的值是%x, 指向的值是%d \n", i, arry_p[i], *arry_p[i])
	}

	set_arry_rel(arry)
	fmt.Print(arry[0])	// 数组传参 改变参数的值
	fmt.Printf("\n")

	set_arry_P(arry_p)		// 指针传参 改变参数的值
	fmt.Print(arry[0])
	fmt.Printf("\n")

	arry = set_arry_arry1(arry)	// 数组传参 若想改变参数的值， 需要返回值 同c语言一样
	fmt.Print(arry[0])
	fmt.Printf("\n")

	// 多维数组
	arrys := [2][2]int{ {1, 2}, {3, 4} }
	fmt.Print(arrys[0][0])
	fmt.Printf("\n")

	// 多维指针
	var arrys_p [2][2]*int
	for i:=0; i < 2; i++{
		for j := 0; j < 2; j++{
			arrys_p[i][j] = &arrys[i][j]
		}
	}
	fmt.Print(*arrys_p[1][0])

}
