package models

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type SprintPlay interface {
	SprintAdds()
	SprintGet()
	SprintGetFull()
	SprintUpdates()
	SprintDeletes()
}

type Sprint struct {
	SprintId  orm.BigIntegerField
	StartTime orm.TimeField
	EndTime   orm.TimeField
	Route     Route
	TimeKey   TimeKey
}

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres",
		"user=tummoc password=specsoid host=127.0.0.1 port=5432 dbname=tummoc sslmode=disable")
	orm.SetMaxIdleConns("default", 10)
	orm.SetMaxOpenConns("default", 100)
	orm.DefaultTimeLoc = time.Local
	orm.RegisterModel(new(Sprint))
}

func SprintAdds(s Sprint) Sprint {
	o := orm.NewOrm()
	var sprint Sprint
	sprint.SprintId = s.SprintId
	sprint.StartTime = s.StartTime
	sprint.EndTime = s.EndTime
	sprint.Route = s.Route
	sprint.TimeKey = s.TimeKey

	id, err := o.Insert(&sprint)
	if err == nil {
		fmt.Println(id)
	}

	return sprint
}

func SprintGet(sprintid int64) Sprint {
	o := orm.NewOrm()
	sprint := Sprint{SprintId: orm.BigIntegerField(sprintid)}

	err := o.Read(&sprint)

	if err == orm.ErrNoRows {
		fmt.Println("No result found")
	} else if err == orm.ErrMissPK {
		fmt.Println("No Primary Keys Found")
	} else {
		fmt.Println(sprint.SprintId)
	}

	return sprint
}

func SprintGetFull() []*Sprint {
	o := orm.NewOrm()
	sprint := new(Sprint)
	qs := o.QueryTable(sprint)
	var s []*Sprint
	qs.All(s)
	return s
}

func SprintUpdates(sprintid int64, ss *Sprint) Sprint {
	o := orm.NewOrm()
	sprint := Sprint{SprintId: orm.BigIntegerField(sprintid)}

	if o.Read(&sprint) == nil {
		sprint.Route = Route{RouteId: ss.Route.RouteId}
		if num, err := o.Update(&sprint, "Route"); err == nil {
			fmt.Println(num)
			return sprint
		} else {
			fmt.Println(err)
		}
		sprint.EndTime = ss.EndTime
		if num, err := o.Update(&sprint, "EndTime"); err == nil {
			fmt.Println(num)
			return sprint
		} else {
			fmt.Println(err)
		}
		sprint.StartTime = ss.StartTime
		if num, err := o.Update(&sprint, "StartTime"); err == nil {
			fmt.Println(num)
			return sprint
		} else {
			fmt.Println(err)
		}
		sprint.TimeKey = ss.TimeKey
		if num, err := o.Update(&sprint, "TimeKey"); err == nil {
			fmt.Println(num)
			return sprint
		} else {
			fmt.Println(err)
		}

	}

	return Sprint{SprintId: 1}

}

func SprintDeletes(sprintid int64, ss Sprint) Sprint {
	o := orm.NewOrm()
	sprint := Sprint{SprintId: orm.BigIntegerField(sprintid)}
	if o.Read(&sprint) == nil {
		o.Delete(sprint)
		return sprint
	} else {
		errs := "cannot delete"
		fmt.Println(errs)
	}
	return Sprint{SprintId: 1}
}
