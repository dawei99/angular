package main

import (
	"container/list"
	"fmt"
)

func main() {
	var s []int
	fmt.Println(s)

	var wuzhi [10]string
	var wuzhiArr []int
	fmt.Printf("%v\n", wuzhi)
	fmt.Printf("%v\n", wuzhiArr)

	var myList = list.New()
	myList.PushFront("first")
	myList.PushBack("end")
	for value := myList.Front() ; value != nil ; value = value.Next() {
		fmt.Println(value.Value)
	}
}
