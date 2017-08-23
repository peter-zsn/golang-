package main

import (
	"fmt"
	"os"
)


func main(){
	a := os.Args[1]		// 传参 [0] 当前执行文件， [1] 为第一个参数 类推
	fmt.Println(1111)
	fmt.Println(a)
}