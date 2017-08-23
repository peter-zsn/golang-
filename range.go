package main

import (
	"fmt"
)
/*
_  表示省略值

range 同python 中的，用于迭代 arry, slice, 通道， 集合(map)的元素
在arry和slice 中返回元素的索引值 返回 （index, arry[index]）
集合（map）中返回key值
 */


func main() {
	nums := []int{1,2,3,4,5}
	sum := 0
	//数组和切片 返回index, arry[index]
	//_ 表示缺省值
	for _, num := range nums{
		sum += num
	}
	fmt.Println(sum)
	for i, num:= range nums{
		if num == 3{
			fmt.Printf("%d 的index is %d\n", num, i)
		}
	}

	//range map集合 返回的是key, value
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs{
		fmt.Println(k, v)
	}

	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for x, y := range "hello"{
		fmt.Printf("x is %d, y is %c\n", x, y)
	}
}
