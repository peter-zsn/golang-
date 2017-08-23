package main

import "fmt"
/*
条件语句
if
if(){} else if(){} else{}
switch var {
	case var1:
	case var2:
	...
	default:
}

select{  //每个case必须是一个通信操作，要么是发送要么是接收。
    case
}
*/


func main() {
	var grade = "B"
	mark := 90
	switch mark {
		case 90: grade = "A"
		case 80: grade = "B"
		case 70: grade = "C"
		case 60: grade = "D"
		default: grade = "E"
	}
	switch {
	case grade == "A" :
		fmt.Printf("优秀%d\n", mark)
	case grade == "B":
		fmt.Printf("良好%d\n", mark)
	default:
		fmt.Printf("不及格%d\n", mark)
	}
	fmt.Printf("你的等级是%s\n", grade)


	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	      case i1 = <-c1:
		 fmt.Printf("received ", i1, " from c1\n")
	      case c2 <- i2:
		 fmt.Printf("sent ", i2, " to c2\n")
	      case i3, ok := (<-c3):  // same as: i3, ok := <-c3
		 if ok {
		    fmt.Printf("received ", i3, " from c3\n")
		 } else {
		    fmt.Printf("c3 is closed\n")
		 }
	      default:
		 fmt.Printf("no communication\n")
	   }

}
