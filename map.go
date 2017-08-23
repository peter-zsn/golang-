package main

import (
	"fmt"
)

/*
map 是无序的键值对集合, 用hash实现的

定义方法
map 关键字
var map_name map[key_type] value_type
make
var map_name = make(map[key_type] value_type)


若不初始化，就会产生一个nil map, nil map 不能存放键值对

delete()  删除集合中的元素， 参数是map 和key
delete(map1, key1)
 */

func main() {
	var map1 map[string]string		//声明
	map1 = make(map[string]string)         //创建

	map1["a"] = "apple"
	map1["b"] = "banaer"
	map1["c"] = "capater"

	for key, value := range map1{
		fmt.Printf("map1[%s] is %s\n", key, value)
	}

	//values, ok := map1["haha"]
	values, ok := map1["a"]			// value 是存在的， ok 是判断key是否存在(bool类型)，
	if ok{
		fmt.Printf("map1[haha] is %s \n", values)
	}else{
		fmt.Printf("map1[hah] no key haha")
	}

	fmt.Println(map1["a"])

	delete(map1, "a")				// 删除map中的key，value
	values1, ok1 := map1["a"]
	if ok1{
		fmt.Printf("map1[haha] is %s \n", values1)
	}else{
		fmt.Printf("map1[hah] no key a")
	}
	fmt.Printf("\n")
	fmt.Println(map1["a"])			// 好像是和 \n
}
