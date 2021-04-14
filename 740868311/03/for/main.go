package main

import (
	"fmt"
	"unicode"
)

func main() {
	// 变种1, 变量在外部定义
	i := 1
	for ; i < 10; i++ {
		if i == 5 {
			continue // 和PHP一致
		} else if i > 5 {
			fmt.Println("over")
			break // 和PHP一致
		}
		fmt.Println(i)
	}

	// 变种2, 定义变量和增量在外部和内部去实现,不在for语句本上实现
	i1 := 5
	for i1 < 10 {
		fmt.Println(i1)
		i1++
	}

	// 变种3, 死循环
	for {
		i++
		if i == 30 {
			break
		}
		fmt.Println(i)
	}

	// for range循环, 专门循环字符串/数组/切片/map/通道(map和通道暂且不知道是什么)
	// 有以下规律 1.数组/切片/字符串返回索引和值  2.map返回键和值 3.通道值返回通道内的值
	s := "Hello杭州"
	for _, v := range s {
		if unicode.Is(unicode.Scripts["Han"], v) {
			fmt.Printf("%c", v)
		}
	}

	switch w := 1; w {
	case 1:
		fmt.Println(w)
	case 2:
		fmt.Println(123)
	}

	w := 1
	switch {
	case w > 0 && w < 2:
		fmt.Println("123")
	case w > 2 && w < 4:
		fmt.Println("456")
	}

	// goto
	for i := 1; i <= 10; i++ {
		for j := 1; j <= i; j++ {
			if j+i == 10 {
				goto XX
			}
			fmt.Printf("%v*%v=%v\t", i, j, i*j)
		}
		fmt.Println()
	}

XX:
	fmt.Println("over")
}
