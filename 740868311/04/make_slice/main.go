package main

import "fmt"

func main() {
	// make() 函数创造切片
	s1 := make([]int, 5, 10)                                          // make(类型, 长度, 容量)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1)) //s1=[0 0 0 0 0] len(s1)=5 cap(s1)=10
	s2 := make([]int, 0, 10)
	fmt.Printf("s2=%v len(s2)=%d cap(s2)=%d", s2, len(s2), cap(s2)) //s1=[0 0 0 0 0] len(s1)=5 cap(s1)=10

	// 切片的赋值 , 共享内存地址, 修改一个,其他引用值都会跟着变
	s3 := []int{1, 3, 5}
	s4 := s3
	fmt.Println(s3, s4) //[1 3 5] [1 3 5]
	s3[0] = 1000
	fmt.Println(s3, s4) //[1000 3 5] [1000 3 5]

	// 切片的遍历
	// 1.索引遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}

	// 1.索引遍历
	for i, v := range s3 {
		fmt.Println(i, v)
	}

}
