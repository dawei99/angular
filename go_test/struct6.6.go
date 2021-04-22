//
// 普通类型添加方法
package main

import "fmt"

//type MyInt int

//func (p putongInt) add(number int) int {
//	//return p + number
//	return p
//}

func main()  {

	test, err := returnTowValue("b")
	test2, err := returnTowValue("c")

	fmt.Printf("%v %v \n", test, err)
	fmt.Printf("%v \n", test2)
}

func returnTowValue (str string) (string, string) {
	return "aaa", str
}

