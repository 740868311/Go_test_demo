package main

import (
	"fmt"
)

func main() {
	// 十进制
	i1 := 10
	fmt.Printf("%d\n", i1)
	fmt.Printf("%o\n", i1) // 把10进制转换成八进制
	fmt.Printf("%x\n", i1) // 把10进制转换成16进制
	fmt.Printf("%b\n", i1) // 把10进制转换成二进制

	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)

	// 十六进制
	i3 := 0x1ff
	fmt.Printf("%d\n", i3)

	// %T查看变量类型
	fmt.Printf("%T\n", i3)

	// 声明int8类型的变量
	i4 := int8(9)
	fmt.Printf("%T\n", i4)
	fmt.Printf("%o\n", i4)
}
