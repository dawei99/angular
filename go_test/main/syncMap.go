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
	var syncMap sync.Map
	syncMap.Store("name", "style");
	syncMap.Store(1, 10.5);
	value,_ := syncMap.Load(1)
	fmt.Println(value)

	syncMap.Range(func(k, value interface{}) bool {
		fmt.Println(k, "=", value)
		return true
	})
}
