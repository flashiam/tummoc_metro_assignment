package models

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type StationPlay interface {
	StationAdd()
	StationGet()
	StationGetFull()
	StationUpdates()
	StationDeletes()
}

type TimeKeyPlay interface {
	TimeKeyAdd()
	TimeKeyGet()
	TimeKeyGetFull()
	TimeKeyUpdates()
	TimeKeyDeletes()
}

type Location struct {
	Latitude  orm.FloatField
	Longitude orm.FloatField
}
type Station struct {
	StationId orm.BigIntegerField
	Name      orm.CharField
	Location  Location
}
type TimeKey struct {
	KeyId     orm.BigIntegerField
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
}

func TimeKeyAdd(r Route) Route {
	o := orm.NewOrm()
	var route Route
	route.RouteId = r.RouteId
	route.StartStation = r.StartStation
	route.EndStation = r.EndStation

	id, err := o.Insert(&route)
	if err == nil {
		fmt.Println(id)
	}

	return route
}

func TimeKeyGet(routid int64) Route {
	o := orm.NewOrm()
	route := Route{RouteId: orm.BigIntegerField(routid)}

	err := o.Read(&route)

	if err == orm.ErrNoRows {
		fmt.Println("No result found")
	} else if err == orm.ErrMissPK {
		fmt.Println("No Primary Keys Found")
	} else {
		fmt.Println(route.RouteId)
	}

	return route
}

func TimeKeyGetFull() []*Route {
	o := orm.NewOrm()
	route := new(Route)
	qs := o.QueryTable(route)
	var r []*Route
	qs.All(r)
	return r
}

func TimeKeyUpdates(routeid int64, rr *Route) Route {
	o := orm.NewOrm()
	route := Route{RouteId: orm.BigIntegerField(routeid)}

	if o.Read(&route) == nil {
		route.StartStation = Station{StationId: rr.StartStation.StationId}
		if num, err := o.Update(&route, "StartStation"); err == nil {
			fmt.Println(num)
			return route
		} else {
			errs := "cannot update start station"
			fmt.Println(errs)
		}
		route.EndStation = Station{StationId: rr.EndStation.StationId}
		if num, err := o.Update(&route, "EndStation"); err == nil {
			fmt.Println(num)
			return route
		} else {
			errs := "cannot update end station"
			fmt.Println(errs)
			return route
		}
	}
	return Route{RouteId: 1}
}

func TimeKeyDeletes(routeid int64, rr Route) Route {
	o := orm.NewOrm()
	route := Route{RouteId: orm.BigIntegerField(routeid)}
	if o.Read(&route) == nil {
		o.Delete(route)
		return route
	} else {
		errs := "cannot delete"
		fmt.Println(errs)
	}
	return Route{RouteId: 1}
}

func StationAdd(r Route) Route {
	o := orm.NewOrm()
	var route Route
	route.RouteId = r.RouteId
	route.StartStation = r.StartStation
	route.EndStation = r.EndStation

	id, err := o.Insert(&route)
	if err == nil {
		fmt.Println(id)
	}

	return route
}

func StationGet(routid int64) Route {
	o := orm.NewOrm()
	route := Route{RouteId: orm.BigIntegerField(routid)}

	err := o.Read(&route)

	if err == orm.ErrNoRows {
		fmt.Println("No result found")
	} else if err == orm.ErrMissPK {
		fmt.Println("No Primary Keys Found")
	} else {
		fmt.Println(route.RouteId)
	}

	return route
}

func StationGetFull() []*Route {
	o := orm.NewOrm()
	route := new(Route)
	qs := o.QueryTable(route)
	var r []*Route
	qs.All(r)
	return r
}

func StationUpdates(routeid int64, rr *Route) Route {
	o := orm.NewOrm()
	route := Route{RouteId: orm.BigIntegerField(routeid)}

	if o.Read(&route) == nil {
		route.StartStation = Station{StationId: rr.StartStation.StationId}
		if num, err := o.Update(&route, "StartStation"); err == nil {
			fmt.Println(num)
			return route
		} else {
			errs := "cannot update start station"
			fmt.Println(errs)
		}
		route.EndStation = Station{StationId: rr.EndStation.StationId}
		if num, err := o.Update(&route, "EndStation"); err == nil {
			fmt.Println(num)
			return route
		} else {
			errs := "cannot update end station"
			fmt.Println(errs)
			return route
		}
	}
	return Route{RouteId: 1}
}

func StationDeletes(routeid int64, rr Route) Route {
	o := orm.NewOrm()
	route := Route{RouteId: orm.BigIntegerField(routeid)}
	if o.Read(&route) == nil {
		o.Delete(route)
		return route
	} else {
		errs := "cannot delete"
		fmt.Println(errs)
	}
	return Route{RouteId: 1}
}
