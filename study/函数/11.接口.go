package main

import "fmt"

type Sing interface {
	sing()
	getName() string
}
type Chicken struct {
	Name string
}

type Cat struct {
	Chicken
}

//	type Dog struct {
//		Chicken
//	}
func (c Chicken) sing() {
	fmt.Println(c.Name, "在唱歌")
}
func (c Chicken) getName() string {
	return c.Name
}
func (c Cat) sing() {
	fmt.Println(c.Name, "在唱歌")
}
func (c Cat) getName() string {
	return c.Name
}
func sing(c Sing) {
	c.sing()
	//fmt.Println(c.getName())
	//接口断言
	switch service := c.(type) {
	case Chicken:
		fmt.Println(service)

	}

}

func Print(val any) {
	fmt.Println(val)
}

func main() {
	c := Chicken{Name: "ik"}
	cat := Cat{
		Chicken{
			Name: "咪咪",
		},
	}
	sing(c)
	sing(cat)
	Print(c)
}
