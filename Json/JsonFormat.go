package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

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

	//取得Query orr Param or Body的參數
	//localhost:8080/users?genger=male&age=25   ?後為query  前面為path or param
	engine.DELETE("/:id", func(context *gin.Context) {
		userName := context.Param("id")
		//context.GetQuery()
		//context.Query()
		//context.Param()
		//context.Request.Body
		context.Writer.WriteString(userName)

	})

	engine.Run(":8090")
}

type Response struct {
	Code    int
	Message string
	Data    interface{}
}
