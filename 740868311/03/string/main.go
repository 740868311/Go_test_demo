package main

import (
	"fmt"
	"strings"
)

func main() {
	// 反斜线转义
	path := "http:\\www.baidu.com\\abc\\123"
	// ``原样显示
	path1 := `http:\www.baidu.com\abc\123`
	fmt.Println(path)
	fmt.Println(path1)
	// 单行字符串
	a := "I'm ok"
	// 多行字符串
	s := `
		人之初
		性本善
	`
	fmt.Println(a)
	fmt.Println(s)

	// 字符串相关操作,打印字符串长度
	fmt.Println(len(path1))
	fmt.Println(len(path))

	//字符串拼接
	name := "张三"
	world := "哈哈哈"

	// 两个字符串可以直接拼接成新的字符串
	ss := name + world

	fmt.Println(ss)
	// Sprintf 格式化后的数据不舒服返回成字符串类型
	ss1 := fmt.Sprintf("%v%v", name, world)
	fmt.Println(ss1)

	// 格式化输出两个需要拼接的字符串
	fmt.Printf("%s%s", name, world)

	// 分割
	ret := strings.Split(path, "\\")
	fmt.Println(ret) //[http: www.baidu.com abc 123]

	// 包含
	fmt.Println(strings.Contains(ss, "张三")) // true
	fmt.Println(strings.Contains(ss, "李四")) // false

	// 前缀
	fmt.Println(strings.HasPrefix(ss, "张三")) // true
	// 后缀
	fmt.Println(strings.HasSuffix(ss, "张三")) // false

	// 字符串出现的位置
	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "c")) // 2
	// 字符串最后出现的位置
	fmt.Println(strings.LastIndex(s4, "b")) // 5

	// 拼接
	fmt.Println(strings.Join(ret, "+")) // http:+www.baidu.com+abc+123

	// 字符串的修改
	s5 := "白萝卜"            // => '白' '萝' '卜'
	s6 := []rune(s5)       // 把字符串强制转换成一个rune切片
	s6[0] = '红'            // 把第一个字符'白'修改成'红'
	fmt.Printf(string(s6)) // 把rune切片强制转换成字符串输出

	b := 65
	fmt.Printf("%c\n", b) //%c 该值对应的unicode码值  输出 A
	c := '中'
	fmt.Printf("%v(%c)", c, c) // 20013(中)

	d := 'a' // rune(int32)
	e := "a" // string

	fmt.Printf("d=%T,e=%T", d, e)   // d=int32,e=string
	fmt.Printf("d=%v,e=%v\n", d, e) // d=97,e=a

	f := "hello小王子"

	g := []rune(f)

	fmt.Println(len(g))

	var aa string
	for i := 0; i < len(g); i++ {
		aa = fmt.Sprintf("%T(%c)(%v)", g[i], g[i], g[i])
		fmt.Println(aa)
	}

}
