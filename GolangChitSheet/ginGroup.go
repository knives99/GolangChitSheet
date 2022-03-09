package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//所有功能統一在user這個Group之下（路由組）
//http://localhost:8090/user/register
//http://localhost:8090/user/login
//http://localhost:8090/user/info
//http://localhost:8090/user/:id

func main() {
	engine := gin.Default()

	routerGroup := engine.Group("/user")
	//routerGroup 符合 IRoutes interface
	routerGroup.POST("/register", registerHandle)

	routerGroup.POST(
		"/login",
		func(context *gin.Context) {
			fulpath := context.FullPath()
			fmt.Println(fulpath)
			context.Writer.Write([]byte("123"))

		},
	)

	routerGroup.DELETE("/:id", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		userName := context.Param("id")
		context.Writer.WriteString(userName)
	})

	engine.Run(":9000")
}

func registerHandle(context *gin.Context) {
	fulPath := context.FullPath()
	fmt.Println(fulPath)
	context.Writer.WriteString(fulPath)
}
