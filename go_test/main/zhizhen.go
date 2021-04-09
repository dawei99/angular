package main

import "fmt"

type wokao struct {
	name  string
	sex int
}
func main(){

	var valueP *int
	valueP = new(int)
	*valueP = 10

	var valueP2 *int = valueP
	fmt.Println(valueP2, valueP)
	*valueP2 = 50
	fmt.Printf("%v\n%v\n",*valueP2, *valueP)

	var bianliang int = 10
	var bianliang2 = &bianliang
	*bianliang2 = 60
	fmt.Println(bianliang, bianliang2)

	my(bianliang2)

/*	var sum *int
	sum = new(int)
	fmt.Printf("第一个指针地址:%v\n",sum)
	var yinyong **wokao = &sum

	fmt.Printf("yinyong:%p\n",yinyong)
	sum = new(wokao)
	fmt.Printf("第二个指针地址:%p\n",sum)
	fmt.Printf("yinyong:%v\n",**yinyong)
	var yinyongSumValue wokao = **yinyong
	yinyongSumValue.name = "通过yinyong设置的name"
	fmt.Printf("第二个指针地址:%v\n",*sum)
	//bakSum := sum
	// 指针变量赋新值

	fmt.Printf("%p\n",sum)

	// 拿出指针变量的中具体值
	var sum2 wokao = *sum
	sum2.name = "liang"
	fmt.Println(*sum)
	fmt.Println(sum2)*/

	// make
	fmt.Println("----------------")
	fmt.Println(make([]int, 10))
	fmt.Println(make(map[string]int, 10))
	fmt.Println(make(chan int))

}

func my(yiyong *int)int{
	fmt.Println(*yiyong)
	return 1
}
