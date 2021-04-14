package main

import "fmt"

func main() {
	var (
		a = 5
		b = 2
	)

	// 算数运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ // 单独的语句, 不能放在=的右边赋值  => a = a + 1
	a-- // 单独的语句, 不能放在=的右边赋值  => a = a - 1

	// 关系运算符
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a < b)
	fmt.Println(a > b)

	// 逻辑运算符  && || !
	fmt.Println(a > b && b > 10)
	age := 30
	if age >= 30 && age < 35 {
		fmt.Println("而立之年")
	} else if age > 35 && age < 50 {
		fmt.Println("送外卖")
	}

	// 位运算符:针对的是二进制数
	// 5的二进制表示: 101
	// 2的二进制表示:  10

	// &:按位与(两位均为1才为1)
	fmt.Println(5 & 2) // 000
	// |:按位或(两位有1个为1就为1)
	fmt.Println(5 | 2) // 111
	// ^:按位异或(两位不一样则为1)
	fmt.Println(5 ^ 2) // 111
	// <<:将二进制位左移指定位数
	fmt.Println(5 << 1)  // 1010 => 10
	fmt.Println(1 << 10) // 10000000000 => 1024
	// >>:将二进制位向右移指定的位数
	fmt.Println(5 >> 1)
	fmt.Println(5 >> 2)
	fmt.Println(5 >> 3)

	var i = int8(1) // i只能存8为, 超过8为会裁剪
	fmt.Println(i << 10)

	// 赋值运算符,  用来给变量赋值的
	var x int
	x = 10
	x += 1 // x = x + 1
	x -= 1 // x = x - 1
	x *= 2 // x = x * 2
	x /= 2
	x %= 2
	x <<= 1
	x >>= 1
	x &= 1 // x = x & 1
	x ^= 1
	x |= 1
}
