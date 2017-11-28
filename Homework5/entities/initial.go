package entities

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var mydb *sql.DB

func init() {
	//https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	db, err := sql.Open("mysql", "root:qyhfbqzh@tcp(127.0.0.1:3306)/Homework5?charset=utf8&parseTime=true")
	//密码：岂因祸福避趋之（q y h f b q zh）
	checkErr(err)
	mydb = db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
