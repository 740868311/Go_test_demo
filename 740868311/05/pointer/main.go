package main

import "fmt"

func main() {
	// 1.&:取内存地址符号
	// n := 18
	// fmt.Println(&n) //0xc000128058

	// p := &n
	// fmt.Printf("%T\n", p) //*int  *标识指针,int表示类型,  *int表示int型的指针

	// // 2. *:根据内存地址取值
	// m := *p
	// fmt.Println(m)        // 18
	// fmt.Printf("%T\n", m) // int

	// 空内存地址会报错
	// var a *int // nil pointer  todo panic: runtime error: invalid memory address or nil pointer dereference
	// *a = 100
	// fmt.Println(*a)

	// new 函数申请一个内存地址
	var a = new(int)
	*a = 100
	fmt.Println(*a) // 100
	fmt.Println(a)  // 0xc0000ac0a8
	var a2 = new(int)
	fmt.Println(a2)  // 0xc0000140f8
	fmt.Println(*a2) // 0
	*a2 = 101
	fmt.Println(a2)  // 0xc0000140f8
	fmt.Println(*a2) // 101

}
