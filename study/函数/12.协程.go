package main

import (
	"fmt"
	"sync"
	"time"
)

func shopping(name string, wait *sync.WaitGroup) {
	fmt.Printf("%s,开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s,购物结束\n", name)
	wait.Done()

}

func main() {
	var wait sync.WaitGroup
	timeStart := time.Now()
	wait.Add(3)
	go shopping("张三", &wait)
	go shopping("李四", &wait)
	go shopping("王五", &wait)
	wait.Wait()
	fmt.Println("购物完成", time.Since(timeStart))
}
