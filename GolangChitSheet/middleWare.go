package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//自定義中間健 條件 ：
//1.符合func函數
//2. return gin.HandleFunc

// https://juejin.cn/post/6844903833164857358

func main() {
	engine := gin.Default()
	//engine全部都需要使用中間件的話
	engine.Use(RequestInfos())

	engine.GET("/query", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  context.FullPath(),
		})
	})

	//單獨接口使用中間件
	engine.GET("/hello", RequestInfos(), func(context *gin.Context) {
		//todo
	})
	engine.Run(":9001")
}

//******
func RequestInfos() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.FullPath()
		method := context.Request.Method
		fmt.Println("請求path " + path)
		fmt.Println("請求method" + method)
		fmt.Println("處理前的狀態碼", context.Writer.Status())
		//context.Next()
		//從客戶發出請求到得到數據 會經過此中間件兩次，如需最後HandleFunc之後的數據就得使用Neㄋt
		context.Next()
		fmt.Println("處理後的狀態碼", context.Writer.Status())
	}
}
