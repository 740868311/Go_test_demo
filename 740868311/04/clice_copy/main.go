package main

import "fmt"

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1

	var a3 = make([]int, 3, 3)
	copy(a3, a1)            // copy 会把原先的内存复制到另外一个内存中,和原先内存不在一起,修改原先切片数据,把影响copy后的切片
	fmt.Println(a1, a2, a3) //[1 3 5] [1 3 5] [1 3 5]
	a1[0] = 100
	fmt.Println(a1, a2, a3) //[100 3 5] [100 3 5] [1 3 5]

	// 将a1中的索引为1的3这个元素删掉
	a1 = append(a1[:1], a1[2:]...) // 取出前半部分和后半部分,拼接一下获取删除后的切片, 实际上是修改底层的数组对应索引的值
	fmt.Println(a1)                //[100 5]
	fmt.Println(cap(a1))           //3

	x1 := [...]int{1, 3, 5} // 数组
	s1 := x1[:]             // 切片
	fmt.Println(s1, len(s1), cap(s1))
	// 1.切片不保存具体的值
	// 2切片对应一个底层数组
	// 3.底层数组都是占用一块连续的内存
	fmt.Printf("%p\n", &s1[0])        //0xc0000be0c0
	s1 = append(s1[:1], s1[2:]...)    // 相当于对切片重新赋值,  赋值后会覆盖底层数组对应索引的值
	fmt.Printf("%p\n", &s1[0])        //0xc0000be0c0
	fmt.Println(s1, len(s1), cap(s1)) //[1 5] 2 3
	// 对切片的赋值,修改了底层数组的数据
	fmt.Println(x1) //[1 5 5]
}
