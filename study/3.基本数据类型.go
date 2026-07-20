package main

import "fmt"

// 默认数字类型是int类型，带个u就是无符号，只能存正整数
func main() {
	var u8 uint8 = 255 //2^8-1
	fmt.Println(u8)
	//0 0000000   uint8
	//1 0000000   int8

	//字符型（不是字符串）
	//比较重要的两个类型是byte（单字节字符）和 rune（多字节字符）
	var a byte = 'a' //ASCII里面的字符
	fmt.Printf("%c %d\n", a, a)
	var a1 uint8 = 97
	fmt.Printf("%c %d\n", a1, a1)
	var a2 rune = '徐' //int32
	fmt.Printf("%c %d\n", a2, a2)
	//字符串类型：字符串类型赋值是双引号

	//转义字符
	fmt.Println("你好\t你好")
	fmt.Println("'你好'你好")
	fmt.Println("\"你好\"你好")
	fmt.Println("\\你好\\你好")

	//反引号
	var a3 = `你好
你好
`
	fmt.Println(a3)
	//go语言中不允许将整型强制转换为布尔型，无法参与数值计算，也无法与其他类型进行转换

	//零值问题，只声明不赋值
	var a4 int
	var a5 string
	var a6 bool
	fmt.Println(a4, a5, a6)
}
