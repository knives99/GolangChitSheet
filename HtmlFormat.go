package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	// template 支援html變量的更改 {{.}}
	//設置html目錄
	engine.LoadHTMLGlob("./html/*")
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
