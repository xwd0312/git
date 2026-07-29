package main

import "fmt"

func main() {
	var age int
	fmt.Printf("请输入你的年龄：")
	fmt.Scan(&age)
	switch {
	case age <= 0:
		fmt.Printf("未成年")
		fallthrough //满足条件跳出switch语句，会继续走下一个条件
	case age <= 18:
		fmt.Printf("成年")
	default:
		fmt.Printf("青年")
	}
	var week int
	fmt.Printf("请输入周几")
	fmt.Scan(&week)
	switch week {
	case 1:
	case 2:
	default:

	}

}
