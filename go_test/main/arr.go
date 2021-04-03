package main

import (
	"fmt"
	"sort"
	"strconv"
)
func main (){
	var array = [4][2]string {
		{"a", "b"},
	}

	array[3][1] = "end"

	var oneA [10]string =  [10]string {"style", "aaaa"}




	fmt.Println(array)
	fmt.Println(oneA)


	fmt.Println("=================")

	// 切片
	//qiepian := []string{"a", "b", "c"}
	q2 := oneA[0:1]

	fmt.Println(oneA,q2)
	fmt.Printf("%p | %p\n", q2, &oneA)

	var qie_string []string

	var qie_int = []int{}
	fmt.Println(qie_int == nil)
	fmt.Println(qie_string == nil)

	fmt.Println(make([]int, 10))

	qie_string_new := append(qie_string, "X")
	qie_string_new = append(qie_string_new, []string{"Z1", "Z2"}...)
	fmt.Println(qie_string_new)

	fmt.Println(strconv.Itoa(123))
	var zhuanInt , _ = strconv.Atoi("123")
	fmt.Println(zhuanInt)

	var testQiapian = []string {"a", "b", "c"}

	for i:=0 ; i<10; i++ {
		testQiapian = append(testQiapian, strconv.Itoa(i))
		fmt.Printf("%p %v %v\n",testQiapian, len(testQiapian), cap(testQiapian))
	}
	fmt.Printf("%p \n",testQiapian)
	fmt.Println(testQiapian, testQiapian[:5])
	fmt.Println(append(testQiapian[:5], append([]string{"A","B"}, testQiapian[5:]...)...))

	slice1 := []int{1,2,3,4,5}
	slice2 := []int{10,11,12}
	fmt.Println(copy(slice1,slice2))
	slice2[1] = 99
	fmt.Println("1:",slice1)
	fmt.Println("2:",slice2)

	setData := []int{50, 60, 70}
	setData = append(setData[:1], setData[1+1:]...)

	fmt.Println(setData)

	var erweiQiepian = [][]int{
		{1,2,3},
		{10,20,30},
	}
	erweiQiepian = append(erweiQiepian, []int{100,200,300})
	fmt.Println(erweiQiepian)

	// map
	var myMap = map[string] map[float32]string {"name": map[float32]string{15.2: "21"}}
	delete(myMap, "name")
	fmt.Println("myMap",myMap)
	for key, value := range myMap {
		fmt.Println(key, "=>", value[15.2])
	}

	var paixuPiepian = []string{"X", "Z", "Y"}
	sort.Strings(paixuPiepian)
	fmt.Println(paixuPiepian)

	var list = []*Profile{
		{Name: "张三", Age: 30, Married: true},
	}

	var queryData = func(name string , age int) *Profile {
		var result *Profile
		for _, item := range list {
			if item.Name == name {
				result = item
			}
		}
		return result
	}
	fmt.Println(queryData("张三", 30))
}


type Profile struct {
	Name    string   // 名字
	Age     int      // 年龄
	Married bool     // 已婚
}