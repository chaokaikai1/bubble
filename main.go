package main

import (
	"bubble/config"
	"bubble/dao"
	"bubble/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
)


func main() {
	//加载配置文件
	if  err:=config.Init("./config/config.ini");err!=nil{
		panic(err.Error())
	}

	//数据库初始化
	if err:=dao.InitMySQL();err!=nil{
		panic(err.Error())
	}
	defer dao.Close()

	//routes 绑定
	r:=routes.InitRoute()
	r.Run(":"+strconv.Itoa(config.Conf.Port))
}
