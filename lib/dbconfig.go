package lib

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Gorm *gorm.DB

func init() {
	Gorm = gormDB()
}
func gormDB() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:55001)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB.SetMaxOpenConns(10)
	mysqlDB.SetMaxIdleConns(5)
	return db
}
