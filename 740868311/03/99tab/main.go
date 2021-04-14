package main

import (
	"fmt"
	"unicode"
)

func main() {
	// 打印乘法表
	for x := 1; x <= 9; x++ {
		for y := 1; y <= x; y++ {

			fmt.Printf("%d*%d=%d ", x, y, (x * y))
		}
		fmt.Println()
	}

	// 编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型。
	a := 1
	b := float32(1)
	c := true
	d := "字符串"
	fmt.Printf("%T %b\n", a, a)
	fmt.Printf("%T %v\n", b, b)
	fmt.Printf("%T %v\n", c, c)
	fmt.Printf("%T %v\n", d, d)

	//编写代码统计出字符串"Hello杭州"中汉字的数量。
	s := "Hello杭州"
	count := 0
	for _, v := range s {
		if unicode.Is(unicode.Scripts["Han"], v) {
			count++
		}
	}
	fmt.Printf("\"%v\"字符串中汉字的数量是%v", s, count)
}
