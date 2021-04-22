package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	//"net/http"
)

func main() {
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.GET("/ping/json", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, map[string]interface{}{
	//		"lang" : "lang",
	//	})
	//})
	//
	//// 渲染html
	//r.LoadHTMLGlob("templates/**/*")
	////r.LoadHTMLFiles("my.tmpl", "templates/my.tmpl")
	//r.GET("/common/index/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "common/index.tmpl", gin.H{
	//		"title": "main website",
	//	})
	//})
	//r.GET("/form/index", func(c *gin.Context){
	//	c.HTML(http.StatusOK, "form/index.tmpl", gin.H{
	//
	//	})
	//})
	//r.GET("/my.tmpl", func(c *gin.Context){
	//	c.HTML(http.StatusOK, "my.tmpl", gin.H{
	//
	//	})
	//})

	//r := gin.Default()
	//html := template.Must(template.ParseFiles("templates/my.tmpl"))
	//r.SetHTMLTemplate(html)

	r := gin.Default()
	r.GET("/user/:name/:pid", func(c *gin.Context) {
		fmt.Printf("%#v \n",c)
		name := c.par
		c.String(http.StatusOK, "Hello %s", name)
	})

	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
