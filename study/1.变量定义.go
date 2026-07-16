package main

import "fmt"

func hello() {
	fmt.Println("Hello World")
}

// 全局变量不能省略 var 关键字，也就是说不能使用短变量声明符号 :=。
var age = 12

var (
	s1 string = "hello"
	s2 string = "world"
)

// const必须声明后赋值
const version string = "1.0.0"

func main() {
	var name string
	name = "xwd"
	fmt.Println(name)

	// 声明并赋值时，类型可以省略。
	var name1 string = "xwd"
	fmt.Println(name1)

	// 函数内部可以使用短变量声明符号 :=。
	name2 := "XWD"
	fmt.Println(name2)

	hello()

	var a1, a2 = 1, 2
	fmt.Println(a1, a2)
	fmt.Println("age:", age)
}
