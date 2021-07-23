package models

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
	// "github.com/beego/beego/v2/client/cache"
	// _ "github.com/beego/beego/v2/client/cache/redis"
)

var (
	Location map[string]*Object
)

var (
	route map[int]*Route
)

type Play interface {
	Add()
	Get()
	GetAll()
	Update()
	Delete()
}

type Route struct {
	RouteId orm.BigIntegerField
	Station Station
}

type Station struct {
	StationId orm.BigIntegerField
	Name      orm.CharField
	Location  orm.JSONField
}

type TimeKey struct {
	KeyId     orm.BigIntegerField
	Route     Route
	Center    Station
	Up        Station
	Down      Station
	TimeValue orm.IntegerField
}

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres",
		"user=tummoc password=specsoid host=127.0.0.1 port=5432 dbname=tummoc sslmode=disable")
	orm.SetMaxIdleConns("default", 10)
	orm.SetMaxOpenConns("default", 100)
	orm.DefaultTimeLoc = time.Local
	orm.RegisterModel(new(Route), new(TimeKey), new(Station))

	// r := Route
	// RouteList = &r
	// bm, err := cache.NewCache("redis", `{"key":"collectionName","conn":":6039","dbNum":"0","password":"thePassWord"}`)
}

func main() {
	name := "default"
	force := true
	verbose := true

	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}

func (rot Route) Get(routid int64) []*Route {
	o := orm.NewOrm()
	route := new(Route)
	qs := o.QueryTable(route)
	var r []*Route

	qs.All(r)

	return r

}

func GetRoute(object Route) (RouteId string) {

}

func GetAllRoutes(object Route) (RouteId string) {

}

func UpdateRoute(object Route) (RouteId string) {

}

func DeleteRoute(object Route) (RouteId string) {

}

func AddStation(object Route) (RouteId string) {

}

func GetStation(object Route) (RouteId string) {

}

func GetAllStation(object Route) (RouteId string) {

}

func UpdateStation(object Route) (RouteId string) {

}

func DeleteStation(object Route) (RouteId string) {

}

func AddTimeKey(object Route) (RouteId string) {

}

func GetTime(object Route) (RouteId string) {

}

func GetAllTimeKey(object Route) (RouteId string) {

}

func UpdateTimeKey(object Route) (RouteId string) {

}

func DeleteTimeKey(object Route) (RouteId string) {

}
