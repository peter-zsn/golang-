package main

import "fmt"

/*
func fun_name(args1, args2) return_type {}
fun_name 函数名字
args* 参数, 需要指定参数的类型
return_type 返回参数

args  值传递  在函数内部对值进行修改， 不影响实际参数
	引用传递， 传参数的地址， 在函数内部对值进行修改， 实际参数跟着改变
 */

/* 方法 */
func max(x, y int) int{
	var res int
	if x > y{
		res = x
	}else{
		res = y
	}
	return res
}

func swap(x, y string) (string, string){
	return y, x
}

/* 闭包 匿名函数, 可直接使用函数内的变量
在匿名函数中递增i，同一个匿名函数i的值一直更新
执行完毕后，返回函数（得以保存）  常用于缓存和封装
*/
func getSequence() func() int  {
	i := 0
	return func() int{
		i += 1
		return i
	}
}

/* 函数作为值 */
func main() {
 	x := 1
	y := 2
	var c int
	c = max(x, y)
	fmt.Println(c)
	a := "hello"
	b := "world"
	fmt.Printf("%s-%s\n", a,b)
	a, b = swap(a, b)
	fmt.Printf("%s-%s\n", a, b)

	number := getSequence()
	fmt.Println(number())
	fmt.Println(number())
	fmt.Println(number())
	number1 := getSequence()
	fmt.Println(number1())
	fmt.Println(number1())
}
