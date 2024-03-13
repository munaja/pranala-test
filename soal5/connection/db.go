package connection

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("error connecting to DB")
	}
}
