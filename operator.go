package main

import "fmt"
/*
算术运算符 + - * / ^ % ++ --
关系运算符 == != > < >= <=   返回值为bool
逻辑运算符 && || !
位运算符 & | ^ >> <<
赋值运算符 = += -= *= /= %= >>= <<= &= ^= |=
其他运算符  & (返回地址)  * (指针变量)
*/

func main() {
	a, b := 9, 2
	var c int
	c = a + b
	fmt.Printf("a + b = %d \n", c)
	c = a - b
	fmt.Printf("a - b = %d \n", c)

	fmt.Printf("a == b %t \n", a == b)		// 打印bool类型 用%t

	x, y := true, false
	if (x && y){
		fmt.Printf("x && y is true \n")
	}else {
		fmt.Printf("x && y is false \n")
	}
	if (!y){
		fmt.Printf("%t\n", y)
	}
	if (x || y){
		fmt.Printf("x || y is true \n" )
	}


}
