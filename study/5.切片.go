package main

import (
	"fmt"
	"sort"
)

func main() {
	var List []string
	List = append(List, "a")
	fmt.Println(List[0])
	//切片实例化
	//通常只有以下情况才需要明确创建非 nil 空切片：
	//接口规定必须返回空数组，而不是空值。
	//转换为 JSON 时，需要区分 null 和 []。
	//程序需要用 slice == nil 判断“是否初始化”。
	var List2 []string
	fmt.Println(List2)
	var name []string
	fmt.Println(name)
	var List3 = make([]string, 0) //初始化成长度为 0 并不能直接通过下标赋值,需要用append
	fmt.Println(List3)
	ageList := make([]int, 3)
	fmt.Println(ageList)
	//切片
	array := [3]int{1, 2, 3}
	slices := array[:]
	fmt.Println(slices)
	fmt.Println(array[0:2]) //左闭右开

	//排序，sort
	var ints = []int{3, 1, 2, 7, 4, 5}
	sort.Ints(ints) //默认升序，完整写法sort.Sort(sort.IntSlice(ints))
	fmt.Println(ints)
	sort.Sort(sort.Reverse(sort.IntSlice(ints))) //降序
	//sort.Sort 这个函数只接受“带有排序说明书的数据”
}
