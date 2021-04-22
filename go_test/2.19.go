// 类型别名
// 类型定义
package main

import "fmt"

type cloneInt int
type aliasInt = int
func main(){
	var clone cloneInt
	var alias aliasInt
	fmt.Printf("%T\n%T", clone,alias)
}
