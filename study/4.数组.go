package main

import "fmt"

func main() {
	var nameList [3]string = [3]string{"xwd", "张三", "李四"}
	fmt.Println(nameList)
	//go语言不支持负向索引
	fmt.Println(nameList[len(nameList)-1])
	var sList = []string{"a", "b", "c", "d"}
	fmt.Println(sList)

}
