package main

import "fmt"

func sayHello() {
	fmt.Println("hello")
	var getName = func() string {
		return "XWD"
	}
	fmt.Println(getName())
}

func sum(numbers ...int) { //多个参数
	var sum int
	for _, n := range numbers {
		sum += n
	}
	fmt.Println(sum)

}

func r1() (string, bool) { //返回多个类型
	return "hello", true
}

// 命名返回值
func r2() (val string, ok bool) {
	val = "hello"
	return val, ok
}

func login() {
	fmt.Println("login")
}

func register() {
	fmt.Println("register")

}
func main() {
	sayHello()
	sum(1, 2, 3, 4, 5)
	fmt.Println("请输入要执行的操作：")
	fmt.Println("1:登录")
	fmt.Println("2:注册")
	var num int
	fmt.Scan(&num)
	switch num {
	case 1:
		login()
	case 2:
		register()
	}
	var userMap = map[int]func(){
		1: sayHello,
	}
	fun, ok := userMap[num]
	if ok {
		fun()
	}
}
