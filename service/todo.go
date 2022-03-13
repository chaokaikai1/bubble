package service

import (
	"bubble/dao"
	"bubble/models"
)
type ToDoService struct{}
var helloService HelloService
//创建todo
func (service *ToDoService)CreateTodo(todo *models.Todo) error{
	err:=dao.DB.Create(&todo).Error
	return err
}

func (service *ToDoService)SayHello(name string) string{
	s:=helloService.SayHello(name)
	return s
}
