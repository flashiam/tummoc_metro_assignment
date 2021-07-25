package models

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
)

type RoutePlay interface {
	RouteAdd()
	RouteGet()
	routeGetFull()
	RouteUpdates()
	RouteDeletes()
}

type Route struct {
	RouteId      int64    `orm:"index;unique;pk"`
	StartStation *Station `orm:"rel(one);"`
	EndStation   *Station `orm:"rel(one);"`
}

func init() {

}

func RouteAdd(r Route) Route {
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

func RouteGet(routid int64) Route {
	o := orm.NewOrm()
	route := Route{RouteId: routid}

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

func RouteGetFull() []*Route {
	o := orm.NewOrm()
	route := new(Route)
	qs := o.QueryTable(route)
	var r []*Route
	qs.All(r)
	return r
}

func RouteUpdates(routeid int64, rr *Route) Route {
	o := orm.NewOrm()
	route := Route{RouteId: routeid}

	if o.Read(&route) == nil {
		route.StartStation = &Station{StationId: rr.StartStation.StationId}
		if num, err := o.Update(&route, "StartStation"); err == nil {
			fmt.Println(num)
			return route
		} else {
			fmt.Println(err)
		}
		route.EndStation = &Station{StationId: rr.EndStation.StationId}
		if num, err := o.Update(&route, "EndStation"); err == nil {
			fmt.Println(num)
			return route
		} else {
			fmt.Println(err)
			return route
		}
	}
	return Route{RouteId: 1}
}

func RouteDeletes(routeid int64, rr Route) Route {
	o := orm.NewOrm()
	route := Route{RouteId: routeid}
	if o.Read(&route) == nil {
		o.Delete(route)
		return route
	} else {
		errs := "cannot delete"
		fmt.Println(errs)
	}
	return Route{RouteId: 1}
}
