package main

import (
	"fmt"
)
/*
struct 结构体  同c一样  可以定义多个类型， 可以定义为简单的对象

type struct_name struct{
	name1 type
	name2 type
	........
}
 */

type Book struct {
	title string
	author string
	price int
}

func main() {
	var book1 Book
	var book2 Book

	book1.author = "zsn"
	book1.title = "good"
	book1.price = 10000
	book2 = book1

	fmt.Printf("book1'title is %s, price is %d \n", book1.title, book1.price)
	fmt.Printf("book2'title is %s, price is %d \n", book2.title, book2.price)

	book1 = set_book(book1)
	fmt.Print(book1.author)
	fmt.Printf("\n")

	//结构体指针
	var book_p *Book
	book_p = &book1
	fmt.Print(book1.price)
	fmt.Printf("\n")
	set_book_p(book_p)
	fmt.Print(book_p.price)
	fmt.Print(book1.price)


}

// 结构体传参
func set_book(book Book) Book {
	book.author = "hello"
	return book
}

// 结构体指针 传参，  不用*取值， 直接指针.元素 赋值
func set_book_p(book * Book){
	book.price = 100
}