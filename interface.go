package main

import (
	"fmt"
)

/*
定义接口

type interface_name interface{
	method1 [return type]
	method2 [return]
}
 */

type phone interface {
	call()
}

type nokia struct {

}

func (phone nokia) call() {
	fmt.Println("i am nokia I call you")
}

type iphone struct {
	name int
}

func (phone iphone)call()  {
	fmt.Println("i am iphone i call you")
	fmt.Println(phone.name)
}

func main()  {
	//var phone1  phone
	//phone1 = new(nokia)
	//phone1.call()
	//
	//var phone2 phone
	//phone2 = new(iphone)
	//phone2.call()

	var phone3 iphone
	phone3.name = 100
	phone3.call()

	var phone4 nokia
	phone4.call()
}