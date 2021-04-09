package main
import (
    "fmt"
    "github.com/tidwall/gjson"
    "reflect"
)

func main() {


    fmt.Println(rune('a'))

     reflect.TypeOf(10)

    //fmt.Println(gjson.Parse(os.Args).Get("file"))

    var i int = 1
    i2 := 10.8
    i3 := float64(i2)
    fmt.Println(i3)
    i = 10
    var f float64 = 1.5
    const chang int = 5
    i++
    fmt.Printf("%v %v %v\n",i,f,chang)

    var a bool = true
    var b bool = false
    if( a && b ) {
        fmt.Printf("成功\n")
    } else {
        fmt.Printf("失败\n")
    }

    const test string = "abc"
    switch test {
        case "1" : fmt.Printf("匹配到了\n");
        default : fmt.Printf("no\n");
    }

    switch {
        case test == "abc" : fmt.Printf("第二次匹配到了\n")
        default : fmt.Printf("第二次no\n")
    }

    a2:=1+1
    a3:=`style`
    a4:='s'
    fmt.Printf("%v %v %v \n", a2, a3, a4)

    //shuzi:=10
    //fmt.Println(shuzi + "5")


    // 管道
//     var guandao chan int

    // 管道
    //var guandao chan int

    ch:=make(chan string)
    go func() {
        fmt.Println("并发进行")
        ch <- "abc"
        fmt.Println("并发退出")
    }()
    jieshou := <-ch
    fmt.Printf("接收完成，接受值：%v\n", jieshou)

    // 并发管道
    guandao2 := make(chan int)
    go printer(guandao2)

    for i:=1 ; i<10; i++ {
        guandao2<-i
    }
    guandao2<-0

    fmt.Println("-----------------------")

    sets2 := []int{1,2,3}

    for i, value := range sets2 {
        fmt.Println(i,value)
    }

    fmt.Print("aaa", "\t")

    // 集合
    fmt.Println("")
    fmt.Println("集合-----------------")
    var my_map map[string]float64                // 申明类型
    my_map = map[string]float64 {"a": 10.5}

    var my_map2 = make(map[string]int)           // 申明类型
    my_map2 = map[string]int {"a": 50}

    var my_map3 = map[string]float64 {"a": 10.5}   // 自动识别类型方式

    //my_map["a"] = 50.2
    fmt.Println(my_map)
    fmt.Println(my_map2)
    my_map3["a"] = -200.946164529459
    fmt.Println(my_map3["a"])
//     my_map3.forEach(func(key, value float64) {
//        fmt.Println(key, value)
//     })

    // 处理json
    fmt.Println("处理json----------------------")
//     jsonstr1 := `{"file":"Book1.xlsx","images":[{"path":"/home/rjh/www/angular/go/image.png"}]}`
//     p.Get("images").ForEach(func(key, value gjson.Result) bool {
//          //fmt.Println(value.Get("path"))
//          fmt.Println(key)
//          return true
//     })

    fmt.Println("处理json----------------------")
    Args := make(map[string]interface{})
    Args["id"] =  1
    Args["plan_id"] =  "第一个"
    fmt.Println(Args["plan_id"])

    //panic("我是PHP,我要抛出一个异常了，等下defer会通过recover捕获这个异常，然后正常处理，使后续程序正常运行。")

//    var r1, r2 float64 = test()
//    fmt.Println(r1, r2)

    a2, b2 := typedTwoValues()
    fmt.Println(a2, b2)
}

func printer(c chan int){
    for {
        data := <-c
        fmt.Printf("获取到数据：%v\n", data)
        if data == 0 {
            fmt.Println("关闭数据接收")
            break
        }
    }
}

func test2() (float64, float64) {
    return 1.2, 2.3
}

func typedTwoValues() (int, int) {
   return 1, 2
}
