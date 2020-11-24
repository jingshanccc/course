package public

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(mysql.Open(MySQLUrl), &gorm.Config{})
	if err != nil {
		log.Println("open database failed, err is " + err.Error())
		return
	}
	//DB.SetMaxOpenConns(20)
	//DB.SetConnMaxLifetime(time.Second * 10)
	//DB.SetMaxIdleConns(10)
}
