package controller

import (
	"bubble/models"
	"bubble/service"
	"github.com/gin-gonic/gin"
	"net/http"
)
var todoService service.ToDoService
func CreateTodo(c *gin.Context){
 	var todo models.Todo

 	c.BindJSON(&todo)
	rep:=models.Response{
		Code: 500,
		Msg: "faild",
		Data: nil,
	}
	if todo.Title==""{
		rep.Msg="title is not null"
		c.JSON(http.StatusOK, rep)
		return
	}
 	if err:=todoService.CreateTodo(&todo);err!=nil{
		rep.Msg=err.Error()
		c.JSON(http.StatusOK, rep)
	}else{
		rep.Code=200
		rep.Msg="success"
		c.JSON(http.StatusOK,rep)
	}

}
func SayHello(c *gin.Context){
	name:=c.Query("name")
	result:=todoService.SayHello(name)
	c.JSON(http.StatusOK,gin.H{"result":result})
}
func DeleteTodo(c *gin.Context){

}

func UpdateTodo(c *gin.Context){

}

func GetTodoList(c *gin.Context){

}



