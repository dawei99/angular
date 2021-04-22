//
// 类型转换 - 重点
package main

import (
	"fmt"
	"strconv"
)

func main(){
	// 数字 <-> 字符串
	num := 100
	fmt.Printf("%s \n", strconv.Itoa(num))

	str ,err := strconv.Atoi("51sdf")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d \n", str)
	}

	/**
	 * -------------
	 * parse 系列
	 * -------------
	 */
	// ParseBool 只接受：1 0 t f T F true false True False TRUE FALSE
	boolVal, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("错误：", err)
	}  else {
		fmt.Printf("值：%v \n",boolVal)
	}

	// ParseInt
	num2, err := strconv.ParseInt("-11", 2, 0)
	fmt.Printf("%v \n", num2)

	// parseUnit
	unum, err := strconv.ParseUint("sd", 10,0)
	if err != nil {
		fmt.Printf("%T \n", unum)
	} else{
		fmt.Println(err)
	}

	// parseFloat
	floatNum, err := strconv.ParseFloat("10.5",0)
	fmt.Printf("%v \n", floatNum)

	/**
	 * -------------
	 * Format 系列
	 * -------------
	 */
	// FormatBool
	str2 := strconv.FormatBool(false)
	fmt.Printf("%v \n", str2)

	// FormatInt
	fmt.Printf("%v \n ", strconv.FormatInt(-30, 2))

	// FormatUint
	fmt.Printf("%v \n ", strconv.FormatUint(30, 2))

	// FormatFloat
	fmt.Printf("%v \n ", strconv.FormatFloat(30.58, 'E', -1, 64))

	// append 切片追加 - 没懂
	//b10 := []byte("ab的")
	b10 := []string{"a"}
	b10 = append(b10, "s")
	fmt.Println(b10)
}
