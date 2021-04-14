package main

import "fmt"

const pi = 3.1415926

const (
	STATUSOK = "200"
	NOTFOUND = 404
)

// 批量常量赋值， 如果某一行声明后没有赋值，默认和上一行一致
const (
	n1 = 200
	n2 = 100
	n3
)

const (
	a1 = iota
	a2
	a3
)

const (
	b1 = iota
	b2
	_
	b3
)

// 插队
const (
	c1 = iota
	c2 = 100
	c3 = iota
	c4
)

// 多个常量声明在一行,iota相当于声明常量行数的索引值, 多个常量在一行声明,iota不会增加, 所以d1 = 1, d2 = 2, d3 = 2, d4 = 3
const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) // << 想当于向左移动一定的位数,  此时iota等于1,也就是1 << 10, 1左移10位,也就是10000000000,二进制换算成10进制,也就是1024
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println(pi)
	fmt.Printf("当前服务器状态码：%s", STATUSOK)

	fmt.Println("n1", n1)
	fmt.Println("n2", n2)
	fmt.Println("n3", n3)

	fmt.Println("a1", a1)
	fmt.Println("a2", a2)
	fmt.Println("a3", a3)

	fmt.Println("b1", b1)
	fmt.Println("b2", b2)
	fmt.Println("b3", b3)

	fmt.Println("c1", c1)
	fmt.Println("c2", c2)
	fmt.Println("c3", c3)
	fmt.Println("c4", c4)

	fmt.Println("d1", d1)
	fmt.Println("d2", d2)
	fmt.Println("d3", d3)
	fmt.Println("d4", d4)
}
