package main

import (
	"fmt"
)

/*
递归函数
 */

//阶乘
func jiechen(x int) int {
	if x == 0{
		return 1
	}else{
		return x * jiechen(x-1)
	}
}

//斐波那契数列
func fibonaqi(n int)int{
	if n < 2{
		return n
	}else{
		return fibonaqi(n - 2) + fibonaqi(n - 1)
	}
}

func main() {
	x := 3
	y := jiechen(x)
	fmt.Println(y)

	for i :=0; i < 10; i++{
		fmt.Printf("%d \t", fibonaqi(i))
	}
}
