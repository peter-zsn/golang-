package main

import (
	"fmt"
	"time"
)

func main() {
	i := 3
	// go 关键词， 后跟函数执行语句--表示创建一个线程，这里调用匿名函数-func，
	// （）传递实参---值为下面（）形参的值
	go func(a int) {
		a += 1
		fmt.Println(a)
	}(i)	// () 里面表示传递的形参
	fmt.Println(i)
	time.Sleep(1)	// 程序优先执行主线程，主线程执行完毕后，等待会，给予执行线程的时间(wait)
	fmt.Println(i)

}
