package main

import (
	"fmt"
	"time"

	// "tummoc/models"
	"tummoc/models"
	_ "tummoc/models"
	_ "tummoc/routers"

	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

func init() {
	// orm.ResetModelCache()

	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres",
		"user=tummoc password=specsoid host=127.0.0.1 port=5432 dbname=tummoc sslmode=disable")
	orm.SetMaxIdleConns("default", 10)
	orm.SetMaxOpenConns("default", 100)
	orm.DefaultTimeLoc = time.Local
	orm.RegisterModel(new(models.Route), new(models.Station), new(models.TimeKey), new(models.Sprint), new(models.Location))

}

func main() {
	name := "default"
	force := false
	verbose := true

	orm.ResetModelCache()
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
