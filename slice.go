package main

import (
	"fmt"
)

/*
slice 是对数组的抽象
数组不可改变长度，
slice 为动态数组， 他的长度是不固定的，可以追加元素，追加时可能使切片的容量增大。
可以定义一个未指定大小的数组定义切片
var slice_arry [] type
或 用make 定义
var slice1 [] type = make([]type, len)
也可以指定容量，其中capacity为可选参数, 默认为len
make([]T, len, capacity)
len 是切片的初始长度也可以是数组的长度

len(slice) 切片的长度
cap(slice) 切片的最大长度（容量） 只要容量足够大，就能一直追加，

如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。

copy() 拷贝切片		copy(s1, s2)
append()切片追加新元素 s1 = append(s1, eval)
 */

func main() {
	s := []int{1,2,3}		//定义
	fmt.Print(s[:])
	fmt.Printf("s length is %d", len(s))
	a := make([]int, 3)		//定义
	a = s[0:2]
	fmt.Print(a[:])
	b := s[:1]			//[index1: index2]slice 截取 ,index1 可以省略默认0
	fmt.Print(b[:])
	fmt.Printf("\n")
	print(cap(s))
	print("\n")

	var nil_s []int		// 空slice
	fmt.Print(nil_s)
	fmt.Printf("\n")


	var s1 []int
	fmt.Println(s1)
	s1 = append(s1, 10)		//append() 切片追加 容量足够大就可以追加
	fmt.Println(s1)
	s1 = append(s1, 11)
	fmt.Println(s1)

	s2 := make([] int , len(s1),2 * cap(s1))
	copy(s2, s1)			// copy 容量必须>=copy对象的容量 必须使用 make建立新的slice
	fmt.Println(s2)
	s2 = append(s2, 13)
	fmt.Println(s2)

}
