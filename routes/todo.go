package routes

import (
	"bubble/config"
	"bubble/controller"
	"github.com/gin-gonic/gin"
)
func InitRoute() *gin.Engine{
	if config.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r:=gin.Default()
	v1Group:=r.Group("v1")
	{
		v1Group.GET("/say",controller.SayHello)
		//添加
		v1Group.POST("/todo",controller.CreateTodo)

		//删除某一个
		v1Group.DELETE("/todo:id", controller.DeleteTodo)

		//修改某一个待办
		v1Group.PUT("/todo:id", controller.UpdateTodo)

		//查看所有的待办
		v1Group.GET("/todo",controller.GetTodoList)
	}
	return r
}