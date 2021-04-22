//
// 函数实现接口
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Invoker interface {
	Call()string
}

type InvokerFunc interface {
	Call()func(string)string
}

type myStruct struct {
	item int
}

func (m *myStruct) Call() string {
	return strconv.Itoa(m.item)
}

func (m *myStruct) get () int {
	old := m
	m.item = m.item + 1
	fmt.Println("newI：", m.item)
	return old.item
}

type testFunc func(float64) string
func (t testFunc) Call (p float64) string {
	return t(p)
}

func (m myStruct) call()int{
	return 1
}

func main (){

	var myInt myStruct = myStruct{100}
	myInt.get()
	fmt.Printf("%d\n", myInt)

	var test testFunc = func(f float64) string {
		return "abc"
	}

	fmt.Printf("%v \n", test.Call(51.5))

	var invoker Invoker
	invoker = &myStruct{100}

	fmt.Printf("%v \n", invoker.Call())

	var invoker2 Invoker
	var function FuncCaller = aa
	invoker2 = function
	fmt.Printf("%T \n", invoker2.Call())

	fmt.Printf("------\n %v \n", reflect.TypeOf("sdf").String() == "String")


}

type FuncCaller func(interface{})
func (f FuncCaller) Call()string {
	return "1"
}

func aa(p interface{}) {

}

