package main

import "fmt"

func main() {
	s1 := []string{"留下", "西溪"}
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1)) //s1=[留下 西溪] len(s1)=2 cap(s1)=2
	// 调用append函数必须用原来的切片变量接受返回值
	s1 = append(s1, "闲林")                                             // append追加元素,原来的底层数组放不下的时候,go底层就会把底层数组换一个
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1)) //s1=[留下 西溪 闲林] len(s1)=3 cap(s1)=4
	s1 = append(s1, "西湖")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1)) //s1=[留下 西溪 闲林 西湖] len(s1)=4 cap(s1)=4
	s1 = append(s1, "古荡")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1)) //s1=[留下 西溪 闲林 西湖 古荡] len(s1)=5 cap(s1)=8
	ss := []string{"青岛湖", "临平", "临安"}
	s1 = append(s1, ss...)                                            // ...表示拆开
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1)) //s1=[留下 西溪 闲林 西湖 古荡 青岛湖 临平 临安] len(s1)=8 cap(s1)=8
}
