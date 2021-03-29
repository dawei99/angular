package main

import (
    "github.com/360EntSecGroup-Skylar/excelize/v2"
    "github.com/buger/jsonparser"
    "fmt"
    "os"
    _ "image/gif"
    _ "image/jpeg"
    _ "image/png"
    "reflect"
    "flag"
    "regexp"
    "strconv"
    //"unsafe"
)

func main () {

    var EXCEL_PATH string   // 操作的Excel文件绝对路径
    var IMGS_JSON string    // 需要添加的图片json

    // 获取并赋值命令行参数
    flag.StringVar(&EXCEL_PATH, "f", "./excel.xlsx", "excel文件绝对路径")
    flag.StringVar(&IMGS_JSON, "i", "", "需要插入图片json")
    flag.Parse()

    if len(IMGS_JSON) == 0 {
        fmt.Println("请传入-i参数")
        os.Exit(1)
    }

    file, err := excelize.OpenFile(EXCEL_PATH) // 打开excel
    throwErr(err, "Excel文件不存在!") // 检测打开失败

    var firstSheetName string = file.GetSheetMap()[1]
    regRowNum := regexp.MustCompile(`[a-zA-Z]+([0-9]+)`) // 匹配行号正则
    // 解析json
    jsonparser.ArrayEach([]byte(IMGS_JSON), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
        pathByte,_,_,errPath:=jsonparser.Get(value, "path") ; path := string(pathByte) // 图片绝对路径
        posByte,_,_,errPos:=jsonparser.Get(value, "pos") ; pos := string(posByte) // 图片位置
        heightByte,_,_,errHeight:=jsonparser.Get(value, "height") ; height := string(heightByte) // 单元格高度
        widthByte,_,_,errWidth:=jsonparser.Get(value, "width") ; width := string(widthByte) // 单元格高度

        rowNumResult := regRowNum.FindAllStringSubmatch(pos, -1)

        // 判断值是否正确
        if errPath != nil || errPos != nil {
            return
        }
        // 添加图片
        if errAddPic := file.AddPicture(firstSheetName, string(pos), string(path), `{ "x_scale": 1,"y_scale": 1,"positioning": "absolute","x_offset": 1,"y_offset": 1}`);err != nil {
            fmt.Println("图片添加失败", path, pos)
            fmt.Println(errAddPic)
            return
        }
        // 设置行高
        if errHeight == nil {
            rowNumString := rowNumResult[0][1]
            rowNum, errRowNum := strconv.Atoi(rowNumString) // 行号
            heightNum, errHeightNum := strconv.ParseFloat(height, 64) // 行高
            if errRowNum == nil && errHeightNum == nil {
                file.SetRowHeight(firstSheetName, rowNum, heightNum)
            } else {
                fmt.Println("行高设置失败：", pos, errRowNum, errHeightNum)
                return
            }
        }
        // 设置列宽
        if errWidth == nil {
            rowNumString := rowNumResult[0][1]
            rowNum, errRowNum := strconv.Atoi(rowNumString) // 行号
            heightNum, errHeightNum := strconv.ParseFloat(height, 64) // 行高
            if errRowNum == nil && errHeightNum == nil {
                file.SetRowHeight(firstSheetName, rowNum, heightNum)
            } else {
                fmt.Println("行高设置失败：", pos, errRowNum, errHeightNum)
                return
            }
        }

    })

    // 保存图像
    errSave := file.Save()
    throwErr(errSave, "Excel文件保存失败!") // 检测打开失败

    reflect.TypeOf(10)
}

func throwErr(err error, msg string){
    if err != nil {
        fmt.Println(msg)
        fmt.Println(err)
        os.Exit(1)
    }
}
