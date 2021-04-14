package main

import "fmt"

func main() {
	b1 := true

	var b2 bool // 默认值是flase

	fmt.Printf("%T\n", b1)
	fmt.Printf("%T\n value:%v\n", b2, b2) // %v是value值
}
