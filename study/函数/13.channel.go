package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan = make(chan int) //声明并初始化一个长度为0的信道

func pay(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s,开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s,购物结束\n", name)
	moneyChan <- money
	wait.Done()
}

func main() {
	var wait sync.WaitGroup
	timeStart := time.Now()
	wait.Add(3)
	go pay("张三", 2, &wait)
	go pay("李四", 3, &wait)
	go pay("王五", 5, &wait)
	go func() {
		wait.Wait()
		close(moneyChan)
	}()
	//for {
	//	money, ok := <-moneyChan
	//	fmt.Println(money, ok)
	//	if !ok {
	//		break
	//	}
	//}
	var moneyList []int
	for money := range moneyChan {
		moneyList = append(moneyList, money)
	}
	fmt.Println("购物完成", time.Since(timeStart))
	fmt.Println("moneyList", moneyList)
}
