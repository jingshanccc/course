package public

import (
	"gitee.com/jingshanccc/course/public/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(mysql.Open(config.MySQLUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println("open database failed, err is " + err.Error())
		return
	}
	//DB.SetMaxOpenConns(20)
	//DB.SetConnMaxLifetime(time.Second * 10)
	//DB.SetMaxIdleConns(10)
}
