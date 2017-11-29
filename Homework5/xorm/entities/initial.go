package entities

import (
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"time"

	// 不可缺，否则会报错：panic: sql: unknown driver "mysql" (forgotten import?)
	_ "github.com/go-sql-driver/mysql"
)

var engine *xorm.Engine

func init() {
	var err error
	//https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	engine, err = xorm.NewEngine("mysql", "root:qyhfbqzh@tcp(127.0.0.1:3306)/Homework5?charset=utf8&parseTime=true")
	//密码：岂因祸福避趋之（q y h f b q zh）
	checkErr(err)
	engine.TZLocation, err = time.LoadLocation("Asia/Shanghai")
	checkErr(err)
	engine.SetMapper(core.SameMapper{})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
