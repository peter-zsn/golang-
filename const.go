package main

import "fmt"


// const 常量 不能被修改的 常量： bool, int, float, double, string, 声明常量必须显示调用 const
// iota 特殊常量，
// 在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1


const a  = "abc"
var c  = "abc"

// 常量,
const (
	unknown = 0
	female = 1
	male				// 默认为上个数值的值
	name = "haha"
	legth = len(name)		// 常量中的函数必须是内置函数
)

const (
	x1 = iota		//后面的值， 默认itoa += 1 直到出现itoa 再显示itoa的值
	x2			//itoa += 1 默认显示itoa的值
	x3
	x4 = 6
	x5
	x6
	x7 = "ha"
	x8
	x9 = iota
	x10
)

func main()  {
 	c = "nihao"
	fmt.Printf("%s\n", a)
	fmt.Printf("%s\n", c)
	const heigh =  1
	const weigh  = 2
	var area int
	fmt.Printf("%d, %d\n", heigh, weigh)
	const x, y, z  = 1, false, "niu"
	area = heigh * weigh
	fmt.Printf("area is %d\n", area)
	fmt.Print(x, y, z)
	fmt.Printf("\n")
	//x = 2			// 常量不能被更改
	fmt.Println(unknown, female, male, name, legth)

	fmt.Println(x1, x2, x3, x4, x5, x6, x7, x8, x8, x10)
}
