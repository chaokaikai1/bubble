package dao

import (
	"bubble/config"
	"bubble/models"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)
var (
	DB *gorm.DB
)
func  InitMySQL()(err error){
	dsn:=strings.Join([]string{config.Conf.MySqlConfig.User,":",config.Conf.MySqlConfig.PassWord,"@tcp(",config.Conf.MySqlConfig.Host,":",strconv.Itoa(config.Conf.MySqlConfig.Port),")/",config.Conf.MySqlConfig.DB,"?charset=utf8mb4&parseTime=True&loc=Local"},"")
	//dsn:="root:123456@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB,err=gorm.Open("mysql",dsn)
	DB.SingularTable(true)
	DB.AutoMigrate(&models.Todo{})
	return
}
func Close(){
	DB.Close()
}