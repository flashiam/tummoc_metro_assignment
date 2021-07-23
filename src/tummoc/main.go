package main

import (
	_ "tummoc/routers"

	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres",
		"user=tummoc password=specsoid host=127.0.0.1 port=5432 dbname=tummoc sslmode=disable")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
