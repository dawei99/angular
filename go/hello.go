package main
import "fmt"

func main() {
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


    // 管道
//     var guandao chan int


    ch:=make(chan string)
    go func() {
        fmt.Println("并发进行")
        ch <- "abc"
        fmt.Println("并发退出")
    }()
    jieshou := <-ch
    fmt.Printf("接收完成，接受值：%v\n", jieshou)
}
