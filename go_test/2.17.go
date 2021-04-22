// 常量
package main

import (
	"fmt"
)

func main () {
	const (
		one int = iota
		tow
		three
	)
	fmt.Println(one, tow ,three)


	const (
		None myInt = iota
		Cpu
		Gpu
	)

	// 常量模拟枚举
	//fmt.Println(None, Cpu, Gpu)

	var test myInt = 3
	fmt.Printf("调用字符串类型的myInt：%s %s \n", test, test)
	//fmt.Println(test)
}
type myInt int
func (m myInt) String() string {
	if m == 0 {
		fmt.Printf("m的类型：%#v \n", m)
	}
	switch m {
	case 0:
		return "Cpu"
	case 1:
		return "Gpt"
	default:
		return "NO"
	}
}
