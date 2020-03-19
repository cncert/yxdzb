package main

import (
	_ "yxdzb/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "yxdzb:12345678@/yxdzb?charset=utf8")
	orm.RunCommand()
}

func main() {
	beego.Run()
}
