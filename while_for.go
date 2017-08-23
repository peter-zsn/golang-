package main

import "fmt"

/*
for  goto break continue
for true{}  死循环

 */

func main()  {
	x := 1
	A: for x < 10{
		x++
		for y := 2; y < x; y++{
			if x % y == 0{
				fmt.Println(x, y)
				goto A
			}
		}
	}
	for x < 20{
		x++
		if x == 15{
			break
		}
		fmt.Println(x)
	}
}
