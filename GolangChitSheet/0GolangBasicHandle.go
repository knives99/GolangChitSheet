package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//取得資料

func main() {

	engine := gin.Default()

	//http://localhost:8080/hello?name=davie     ?後為query  前面為path or param or 接口
	//Handle(代指 GET POST DELETE 等操作)	//HandleFunc
	//func (group *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) IRoutes
	engine.Handle("GET", "/hello", HandleFunc)
	engine.Handle("DELETE", "/user/:id", HandleDelete)
	engine.POST("/post", HandlePost)

	engine.Run("8090") //端口 港口 port

}

func HandleFunc(ctx *gin.Context) {
	//印出path only
	path := ctx.FullPath()
	fmt.Println(path)

	//取得前端輸出的資料
	name := ctx.DefaultQuery("name", "hello")
	//ctx.DefaultQuery() 可設置默認參數
	//ctx.GetQuery()
	//ctx.Query()
	//ctx.Param()
	//ctx.Request.Body
	//ctx.PostForm() 從Request Body 取得單一key value

	//如果Query有多數的情況下 ?name=davie&classes=project
	var student Student
	err := ctx.ShouldBindQuery(&student) //直接綁定query參數到student
	if err != nil {
		log.Fatal(err.Error())
	}

	//輸出 返回給前端
	ctx.Writer.Write([]byte("hello" + name))
	ctx.Writer.Write([]byte(student.Name))

}

func HandlePost(ctx *gin.Context) {

	//擷取綁定user 傳來的post數據
	//POST-Content-Type ==  application/x-www-form-urlencoded
	var register Register
	err := ctx.ShouldBind(&register)
	if err != nil {

	}

	//POST- Content-Type == application/json
	var person Person2
	err = ctx.ShouldBindJSON(&person)
	if err != nil {

	}

	ctx.Writer.Write([]byte(register.UserName))
	ctx.Writer.Write([]byte(person.Name))

}

func HandleDelete(ctx *gin.Context) {
	userId := ctx.Param("id")
	fmt.Println(userId)
	ctx.Writer.Write([]byte(userId))

	var input struct {
		ID int64 `uri:"id"`
	}
	err := ctx.ShouldBindUri(&input)
	if err != nil {
	}
	ctx.JSON(http.StatusOK, input.ID)
}

type Student struct {
	//?name=davie&classes=project
	Name    string `form:"name"`
	Classes string `form:"classes"`
}

type Register struct {
	UserName string `form:"name"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

type Person2 struct {
	Name string `form:"name"`
	Sex  string `form:"sex"`
	Age  int    `form:"age"`
}
