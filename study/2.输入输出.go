package main

import "fmt"

func main() {
	fmt.Println("你好", "你好")
	fmt.Print("XWD")
	fmt.Printf("%s 你好！", "XWD") //格式化输出，没有自动换行，可以手动加\n换行   fmt.Printf("%s 你好！\n", "XWD")
	fmt.Printf("%d\n", 3)
	fmt.Printf("%.2f\n", 3.1415)      //保留两位小数
	fmt.Printf("%T %T\n", "abc", 200) //打印类型
	fmt.Printf("%v\n", "")            //打印任意字符
	fmt.Printf("%#v\n", "")           //可以打印出空字符串

	//将格式化的内容复制给一个变量
	var f = fmt.Sprintf("%.2f", 3.1415)
	fmt.Println(f)

	//输入
	fmt.Println("请输入你的名字：")
	var name string //定义了一个变量即使没有赋值，也会有一个零值
	fmt.Scan(&name)
	n, err := fmt.Scan(&name)
	fmt.Println(n, err) //这里的 n 表示 fmt.Scan 成功读取并赋值的参数数量
	fmt.Printf("你的名字是：%s\n", name)

}
