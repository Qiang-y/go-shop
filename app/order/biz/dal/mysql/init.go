package mysql

import (
	"biz-demo/gomall/app/order/biz/model"
	"biz-demo/gomall/app/order/conf"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:3307)/user?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if os.Getenv("GO_ENV") != "online" {
		err := DB.AutoMigrate(&model.Order{}, &model.OrderItem{})
		if err != nil {
			panic(err)
		}
	}
}