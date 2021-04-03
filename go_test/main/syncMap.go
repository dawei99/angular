package main

import (
	"fmt"
	"sync"
)

func main(){
	// 普通map
	//var testMap = make(map[string]int)
	//
	//fmt.Println(testMap)
	//go func () {
	//	for {
	//		testMap["a"] = 1
	//	}
	//}()
	//go func() {
	//	for {
	//		_ = testMap["a"]
	//	}
	//}()

	// sync.map
	var syncMap  sync.Map
	syncMap.Store("a", "b")
	syncMap.Store(1, 10.8)
	fmt.Println(syncMap.Load("a"))
	fmt.Println(syncMap.Load("a"))
	fmt.Println(syncMap.Load("a"))

	//for {
	//
	//}
}
