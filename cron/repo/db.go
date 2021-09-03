package repo

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	sh, _ := time.LoadLocation("Asia/Shanghai")

	time.Local = sh

	engine = connectDb()
}

func connectDb() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:root@(127.0.0.1:3306)/test?charset=utf8mb4")
	if err != nil {
		panic(err)
	}

	engine.SetTZLocation(time.Local)
	engine.ShowSQL(true)
	return engine
}

func GetEngine() *xorm.Engine {
	return engine
}
