package main

import (
	"fmt"
)

var sum int

func main() {
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("sum = %d\n", sum)
	//for { //死循环可以省略i
	//fmt.Println(time.Now())
	//time.Sleep(2 * time.Second)
	//}
	var List = []string{"a", "b", "c", "d", "e"}
	for i := 0; i < len(List); i++ {
		fmt.Println(i, List[i])
	}
	for index, item := range List { //索引与值
		fmt.Println(index, item)

	}
	var UserMap = map[string]string{"name": "XWD"}
	for k, v := range UserMap {
		fmt.Println(k, v)
	}
}
