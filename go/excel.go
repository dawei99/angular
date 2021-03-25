package main

import (
    "github.com/360EntSecGroup-Skylar/excelize/v2"
    "github.com/tidwall/gjson"
    "fmt"
    "os"
    _ "image/gif"
    _ "image/jpeg"
    _ "image/png"
    "reflect"
)

func main () {
    reflect.TypeOf(10)

        jsonstr := os.Args[1] // 传入json
        parse := gjson.Parse(jsonstr)
        //excelFile := parse.Get("file").String() // Excel文件
        file, err := excelize.OpenFile("Book1.xlsx") // 打开excel
        // 检测打开失败
        if err != nil {
            fmt.Println(err)
            return
        }
        var firstSheetName string = file.GetSheetMap()[1]

        _, errStat := os.Stat("/home/rjh/www/angular/go/image.png")
        if errStat != nil {
            fmt.Println("excel文件不存在")
            os.Exit(1) //
        }

        // 处理excel
        parse.Get("images").ForEach(func(key, value gjson.Result) bool {
            path := value.Get("path").String()
            pos := value.Get("pos").String()
            if len(path)==0 || len(pos)==0 {
                fmt.Printf("图片添加失败：\n位置：%v\n路径：%v", pos, path)
                return true
            }
            // 设置单元格高度
            if errSetHeight := file.SetRowHeight(firstSheetName, 15, 33) ; errSetHeight != nil {
                fmt.Printf("签字单元格高度设置失败：\n位置：%v\n路径：%v", pos, path)
                return true
            }
            // 添加签字/签章图片
            if err := file.AddPicture(firstSheetName, pos, path, `{ "x_scale": 0.5,"y_scale": 0.5,"positioning": "absolute"}`) ; err != nil {
              fmt.Println("图片添加失败：", err)
            }
            return true
        })

        // 保存文件
        if err := file.Save() ; err != nil {
            fmt.Println(err)
        }

}
