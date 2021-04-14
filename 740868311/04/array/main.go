package main

import "fmt"

//数组

// 存放元素的容器
// 必须指定存放的元素的类型和容量(长度)
// 数组的长度是数据类型的一部分
func main() {
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("a1:%T a2:%T", a1, a2)

	// 数组的初始化
	// 如果不初始化: 默认元素都是零值(布尔值:false, 整型和浮点型都是0, 字符串"")
	fmt.Println(a1, a2)

	// 1.初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	// 2. 初始化方式2
	// ...根据初始值自动腿短数组的长度是多少
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a10) //[0 1 2 3 4 5 6 7]
	// 3. 初始化方式3
	a3 := [5]int{1, 2} // 后边的置为领值[1 2 0 0 0]
	fmt.Println(a3)
	a31 := [5]int{0: 1, 4: 2} //根据索引赋值, 没赋值的置为领值[1 0 0 0 2]
	fmt.Println(a31)

	// 数组的遍历
	citys := [...]string{"杭州", "北京", "上海"}
	// 根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	for _, city := range citys {
		fmt.Println(city)
	}

	// 多维数组
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)

	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 数组是值类型
	b1 := [3]int{1, 2, 3} // [1 2 3]
	b2 := b1              // [1 2 3]
	b2[0] = 100           // [100 2 3]
	fmt.Println(b1, b2)

	// 作业, 求数组的和
	c1 := [...]int{1, 3, 5, 7, 8}

	total := 0
	for _, i := range c1 {
		total += i
	}

	fmt.Println(total)

	// 找出数组中和为指定值的两个元素的下标
	target := 8
	for k, v := range c1 {
		for a := k + 1; a < len(c1); a++ {
			if v+c1[a] == target {
				fmt.Printf("%d %d\t", k, a)
			}
		}
	}
}
