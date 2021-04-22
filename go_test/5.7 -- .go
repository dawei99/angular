//
// 函数实现接口
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myFuncInter myFuncer
	var function myFunc = element
	myFuncInter = function
	fmt.Printf("%v \n", myFuncInter.Use(10.5, "abc"))
}

type myFuncer interface {
	Use(interface{}, string)(string, string)
}

type myFunc func(interface{}, string) (string, string)

func (m myFunc) Use(i interface{}, s string) (string, string) {
	return m(i, s)
}

func element (i interface{}, s string) (string, string) {
	fmt.Printf("%v \n", reflect.TypeOf(i).String())
	//switch reflect.TypeOf(i).String() {
	//
	//}
	fmt.Printf("%v \n", i.(string))
	return "1", "2"
	//return ""
}
