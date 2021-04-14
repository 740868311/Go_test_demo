package main

import "fmt"

func main() {
	age := 30
	if age >= 30 {
		fmt.Println("30而立了老铁, 赶紧努力")
	} else if age <= 18 {
		fmt.Println("青春就是浪")
	} else {
		fmt.Println("123")
	}

	// 作用域, age1变量此时只在if条件判断语句中生效
	if age1 := 19; age > 8 {
		fmt.Println(age1)
	}
}
