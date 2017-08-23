package main

import (
	"fmt"
)
/*
错误处理
函数通常在最后的返回值中返回错误信息。使用errors.New 可返回一个错误信息：
type error interface{
	Error() string
}
 */

type DError struct {
	a int
	b int
}

func (de *DError) Error() string{
	str := `error
		a is %d
		b is %d
		b not 0
	`
	return fmt.Sprintf(str, de.a, de.b)
}

func dev(a int, b int)(result int, errorMsg string)  {
	if b == 0{
		data :=DError{
			a: a,
			b: b,
		}
		errorMsg = data.Error()
		return
	}else{
		return a / b, ""
	}

}

func main() {
	if result, errorMsg :=dev(100, 0); errorMsg !=""{
		fmt.Println(errorMsg)
	}else {
		fmt.Println(result)
	}

}
