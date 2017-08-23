package main

import (
	"fmt"
)
/*
数组中的类型必须相同 与c一样 长度固定
var arry_name = [size] type{}
 */

var arry1 = [5]int{1,2,3,4,5}	// 一维数组.
	// 多维数组
var arrys2 = [2][2]int{ {1, 2}, {3, 4} }

// 数组传参    // 在函数内该改变参数的值， 主函数中参数的值斌不改变，
// 若想改变 可传指针 arry_p [2][3]* int 或返回参数，进行赋值
func print_arry(arry [2][3] int)  {
	var j int
	for i :=0; i < 2; i++{
		for j = 0; j < 3; j++{
			fmt.Printf("第%d行 第%d列的内存地址is %d \n", i, j, &arry[i][j])
		}
	}
	arry[0][0] = 10
}

// 不指定大小数组传参, 外边的参数也许是不指定大小的， 感觉不变使用
func print_arry_no(arry [] int)  {
	fmt.Printf("%d \n", arry[0])
}

func main() {
	//fmt.Print(arry1)
	arry1[1] = 10
	//fmt.Print(arry1)
	// 数据定义和赋值
	a := arry1[0]
	fmt.Print(a)
	//fmt.Printf("\n")
	var arry2 [20] int
	var i, j int
	for i = 0; i < 20; i++{
		arry2[i] = i + 100
	}
	//for j = 0; j < 10; j++{
	//	fmt.Printf("arry[%d] is %d \n", j, arry2[j])
	//}
	//fmt.Print(&arry1[0], &arry1[1])   //指针位置（内存）存储相同

	// 二维数组 多维数组同c一样，
	var arrys [2][3] int
	x := 1
	for i=0; i<2;i++{
		for j=0; j<3; j++{
			x += 1
			arrys[i][j] = x
		}
	}

	var arrys1 =  [] int{1,2,3}
	print_arry(arrys)
	print_arry_no(arrys1)
	fmt.Printf("函数中传数组的话，值不变 %d\n", arrys[0][0])
}
