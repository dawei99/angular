//
// Go语言方法和接收器

package main

import "fmt"

type students struct {
	lists []string
}

func main() {
	student := &students{}
	insert(student, "A")
	insert(student, "B")
	fmt.Printf("%#v \n", student)

	bag := Bag{}
	bag.insert("AAA")
	fmt.Printf("%#v \n", bag)

	Chapter3()
}

func insert(b *students, value string){
	b.lists = append(b.lists, value)
}

//
// （2） 结构体接收器
type Bag struct {
	nodes []string
}

func (b *Bag) insert (value string) {
	b.nodes = append(b.nodes, value)
}

//
// （3）指针与非指针类型接收器
// 使用总结：在计算机中，小对象用复制大对象用指针
func Chapter3 ()  {
	// 指针类型
	test := Property{}
	test.set("ONE2")
	fmt.Println(test.get())
	// 非指针
	// 当方法作用于非指针接收器时，Go语言会在代码运行时将接收器的值复制一份，在非指针接收器的方法中可以获取接收器的成员值，但修改后无效。
	fmt.Println(test.clone(&Property{"safsafsadfsdf"}))

}

type Property struct {
	value string
}

func (p *Property) set (value string) {
	p.value = value
}

func (p *Property) get () string {
	return p.value
}

func (p Property) clone(newValue *Property) Property {
	p.value = newValue.value+"_clone"
	return p
}
