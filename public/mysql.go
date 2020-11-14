package public
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", MySQLUrl)
	if err != nil {
		log.Println("open database failed, err is " + err.Error())
		return
	}
	//DB.SetMaxOpenConns(20)
	//DB.SetConnMaxLifetime(time.Second * 10)
	//DB.SetMaxIdleConns(10)
	if err = DB.Ping(); err != nil {
		log.Println("connect database failed, err is " + err.Error())
		return
	}
}