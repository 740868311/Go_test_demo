package main

import "fmt"

func main() {
	n := 100
	fmt.Printf("%T\n", n)  // 查看类型
	fmt.Printf("%#v\n", n) // 查看类型的值
	fmt.Printf("%b\n", n)  // 二进制
	fmt.Printf("%d\n", n)  // 十进制
	fmt.Printf("%o\n", n)  // 八进制
	fmt.Printf("%x\n", n)  // 十六进制

	s := "Hello"
	fmt.Printf("%s\n", s)  // 字符串类型
	fmt.Printf("%v\n", s)  //万能的可以打印所有的类型
	fmt.Printf("%#v\n", s) //%#v 会打印出来对应的类型,字符串会带双引号

	b := 65
	fmt.Printf("%c\n", b) //%c 该值对应的unicode码值  输出 A
	c := '中'
	fmt.Printf("%v(%c)", c, c) // 20013(中)

	d := 'a' // rune(int32)
	e := "a" // string

	fmt.Printf("d=%T,e=%T", d, e) // d=int32,e=string
	fmt.Printf("d=%v,e=%v", d, e) // d=97,e=a
}
