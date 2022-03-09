package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//輸出資料 返回請求

func main() {
	engine := gin.Default()

	engine.GET("/hellobyte", func(context *gin.Context) {
		fullPath := context.FullPath()
		fmt.Println(fullPath)
		context.Writer.Write([]byte("fullPath" + fullPath))
	})

	engine.GET("/hellostring", func(context *gin.Context) {
		fullPath := "入境:" + context.FullPath()
		fmt.Println(fullPath)
		context.Writer.WriteString(fullPath)
	})

	//返回JSON
	engine.GET("/hellojson", func(context *gin.Context) {
		fullpath := "請求路徑" + context.FullPath()
		fmt.Println("請求入境ＪＳＯＮ" + fullpath)
		context.JSON(200, map[string]interface{}{
			"code":    1,
			"message": "ok",
			"data":    fullpath,
		})
	})

	engine.GET("/jsonstruct", func(context *gin.Context) {
		fullPath := "JsonStruct請求路徑：" + context.FullPath()
		fmt.Println(fullPath)

		resp := Response{1, "ok", fullPath}
		context.JSON(200, &resp)
	})

	//返回Html
	// template 支援html變量的更改 {{.}}
	//設置html目錄
	engine.LoadHTMLGlob("./html/*")
	//要加載靜態文件 需用static
	engine.Static("/img", "./img")
	engine.GET("/hellohtml", func(context *gin.Context) {
		fullPath := "hellohtml" + context.FullPath()
		fmt.Println(fullPath)
		context.HTML(http.StatusOK, "index.html", gin.H{
			"fullPath": fullPath,
			"title":    "gin教程",
		})

	})

	engine.Run(":8090")
}

type Response struct {
	Code    int
	Message string
	Data    interface{}
}
