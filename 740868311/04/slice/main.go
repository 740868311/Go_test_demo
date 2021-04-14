package main

import "fmt"

func main() {
	// 切片的定义
	var s1 []int    // 定义一个存放int类型元素的切片
	var s2 []string // 定义一个存放字符串类型元素的切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // true , 相当于没有开辟内存空间
	fmt.Println(s2 == nil) // true
	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"留下", "闲林", "西溪"}

	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // false
	fmt.Println(s2 == nil) // false

	// 长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1)) //len(s1):3 cap(s1):3
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2)) //len(s2):3 cap(s2):3

	// 2.由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	a3 := a1[0:4] // 基于一个数组切割,左包含右不包含,左闭右开 [1 3 5 7]
	fmt.Println(a3)
	a4 := a1[1:6] // 冒号两边是索引值,  包含左边索引值,不包含右边索引值, 所以是切从1到5的所有值
	fmt.Println(a4)
	a5 := a1[:4] // -> [0:4]
	a6 := a1[3:] // -> [3:len(a1)]	[7 9 11 13]
	a7 := a1[:]  // -> [0:len(a1)]  [1 3 5 7 9 11 13]
	fmt.Println(a5, a6, a7)

	// 切片的容量是指底层数组的容量
	fmt.Printf("len(a5):%d cap(a5):%d\n", len(a5), cap(a5)) //len(a5):4 cap(a5):7

	// 底层数组从切片的第一个元素到最后的元素数量
	fmt.Printf("len(a6):%d cap(a6):%d\n", len(a6), cap(a6)) //len(a6):4 cap(a6):4

	// 切片再切片
	a8 := a6[3:]
	fmt.Printf("len(a8):%d cap(a8):%d\n", len(a8), cap(a8)) //len(s8):1 cap(s8):1

	a1[6] = 13000          // 修改了底层数组的值
	fmt.Println("a6:", a6) //a6: [7 9 11 13000]
	fmt.Println("a8:", a8) //a8: [13000]

}
