package main

import "fmt"

func main() {
	//map创建的时候一定要初始化，就是一定要复制
	var UserMap map[int]string = map[int]string{
		1: "xwd",
		2: "xwd123",
		4: "",
	}
	fmt.Println(UserMap)
	fmt.Printf("%#v\n", UserMap[3]) //不会报错，会打印类型的默认值
	value := UserMap[3]
	fmt.Println(value)
	value, ok := UserMap[4]
	fmt.Println(ok)

	delete(UserMap, 4)

	var AMap = make(map[int]string)
	fmt.Println(AMap)
}
