package models

import (
	"fmt"
	"time"
	. "tummoc/micro"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
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
	LocId     int64   `orm:"index;unique;pk"`
	Latitude  float32 `orm:"digits(8);decimals(6)"`
	Longitude float32 `orm:"digits(8);decimals(6)"`
}
type Station struct {
	StationId int64 `orm:"index;unique;pk"`
	Name      string
	Location  *Location `orm:"rel(one);"`
}
type TimeKey struct {
	KeyId         int64    `orm:"index;unique;pk"`
	Center        *Station `orm:"rel(fk);"`
	Up            *Station `orm:"rel(fk);"`
	Down          *Station `orm:"rel(fk);"`
	TimeValueUp   int8
	TimeValueDown int8
	Sprint        []*Sprint `orm:"reverse(many)"`
}

func TimeKeyAdd(t TimeKey) TimeKey {
	o := orm.NewOrm()
	var timekey TimeKey
	timekey.KeyId = t.KeyId
	timekey.Up = t.Up
	timekey.Down = t.Down

	id, err := o.Insert(&timekey)
	if err == nil {
		fmt.Println(id)
	}

	return timekey
}

func TimeKeyGet(timekeyid int64) TimeKey {
	o := orm.NewOrm()
	timekey := TimeKey{KeyId: timekeyid}

	err := o.Read(&timekey)

	if err == orm.ErrNoRows {
		fmt.Println("No result found")
	} else if err == orm.ErrMissPK {
		fmt.Println("No Primary Keys Found")
	} else {
		fmt.Println(timekey.KeyId)
	}

	return timekey
}

func TimeKeyGetFull() []*TimeKey {
	o := orm.NewOrm()
	timekey := new(TimeKey)
	qs := o.QueryTable(timekey)
	var t []*TimeKey
	qs.All(t)
	return t
}

func TimeKeyUpdates(timekeyid int64, tt *TimeKey) TimeKey {
	o := orm.NewOrm()
	timekey := TimeKey{KeyId: timekeyid}

	if o.Read(&timekey) == nil {
		timekey.Center = &Station{StationId: tt.Center.StationId}
		if num, err := o.Update(&timekey, "Center"); err == nil {
			fmt.Println(num)
			return timekey
		} else {
			fmt.Println(err)
		}
		timekey.Up = &Station{StationId: tt.Up.StationId}
		if num, err := o.Update(&timekey, "Up"); err == nil {
			fmt.Println(num)
			return timekey
		} else {
			fmt.Println(err)
		}
		timekey.Down = &Station{StationId: tt.Down.StationId}
		if num, err := o.Update(&timekey, "Down"); err == nil {
			fmt.Println(num)
			return timekey
		} else {
			fmt.Println(err)
		}
		timekey.TimeValueDown = tt.TimeValueDown
		if num, err := o.Update(&timekey, "TimeValueDown"); err == nil {
			fmt.Println(num)
			return timekey
		} else {
			fmt.Println(err)
		}
		timekey.TimeValueUp = tt.TimeValueUp
		if num, err := o.Update(&timekey, "TimeValueUp"); err == nil {
			fmt.Println(num)
			return timekey
		} else {
			fmt.Println(err)
		}
	}

	return TimeKey{KeyId: 1}

}

func TimeKeyDeletes(timekeyid int64, tt TimeKey) TimeKey {
	o := orm.NewOrm()
	timekey := TimeKey{KeyId: timekeyid}
	if o.Read(&timekey) == nil {
		o.Delete(timekey)
		return timekey
	} else {
		errs := "cannot delete"
		fmt.Println(errs)
	}
	return TimeKey{KeyId: 1}
}

func StationAdd(s Station) Station {
	o := orm.NewOrm()
	var station Station
	station.StationId = s.StationId
	station.Name = s.Name
	station.Location = s.Location

	id, err := o.Insert(&station)
	if err == nil {
		fmt.Println(id)
	}

	return station
}

func StationGet(stationid int64) Station {
	o := orm.NewOrm()
	station := Station{StationId: stationid}

	err := o.Read(&station)

	if err == orm.ErrNoRows {
		fmt.Println("No result found")
	} else if err == orm.ErrMissPK {
		fmt.Println("No Primary Keys Found")
	} else {
		fmt.Println(station.StationId)
	}

	return station
}

func StationGetFull() []*Station {
	o := orm.NewOrm()
	station := new(Station)
	qs := o.QueryTable(station)
	var s []*Station
	qs.All(s)
	return s
}

func StationUpdates(stationid int64, ss *Station) Station {
	o := orm.NewOrm()
	station := Station{StationId: stationid}

	if o.Read(&station) == nil {
		station.Name = ss.Name
		if num, err := o.Update(&station, "Name"); err == nil {
			fmt.Println(num)
			return station
		} else {
			fmt.Println(err)
		}
		station.Location = ss.Location
		if num, err := o.Update(&station, "Location"); err == nil {
			fmt.Println(num)
			return station
		} else {
			fmt.Println(err)
		}

	}
	return Station{StationId: 1}
}

func StationDeletes(stationid int64, ss Station) Station {
	o := orm.NewOrm()
	station := Station{StationId: stationid}
	if o.Read(&station) == nil {
		o.Delete(station)
		return station
	} else {
		errs := "cannot delete"
		fmt.Println(errs)
	}
	return Station{StationId: 1}
}

type Coordinate struct {
	Lats  float64
	Longs float64
}

func GetLocation(coordinate *Coordinate) *Station {
	o := orm.NewOrm()
	location := new(Location)
	qs := o.QueryTable(location)
	var l []*Location
	qs.All(l)

	var min float64 = Distance(float64(l[0].Latitude), float64(l[0].Longitude), coordinate.Lats, coordinate.Longs)
	var locid int64 = l[0].LocId

	for _, j := range l {
		d := Distance(float64(j.Latitude), float64(j.Longitude), coordinate.Lats, coordinate.Longs)
		if min > d {
			min = d
			locid = j.LocId
		}
	}

	station := Station{Location: &Location{LocId: locid}}
	err := o.Read(&station)

	if err == orm.ErrNoRows {
		fmt.Println("No result found")
	} else if err == orm.ErrMissPK {
		fmt.Println("No Primary Keys Found")
	} else {
		fmt.Println(station.StationId)
	}

	return &station
}

func GetTime(station *Station, time *time.Time) int64 {
	timekeyforthisstation := TimeKey{Center: station}

	o := orm.NewOrm()
	sprint := new(Sprint)
	qs := o.QueryTable(sprint)
	var s []*Sprint
	qs.All(s)

	var currentSprint *Sprint
	var currentRoute *Route
	var currentStation *Station
	min := s[0].StartTime

	for _, j := range s {
		if time.Sub(min) > time.Sub(j.StartTime) {
			min = j.StartTime
			currentSprint = j
		} else {
			currentSprint = s[0]
		}
	}
	currentRoute = currentSprint.Route
	timelapsedfromsprintstart := int64(time.Sub(min).Minutes())

	var alltimekeysforthisroute []*TimeKey

	for _, k := range s {
		tks := Sprint{Route: currentRoute}
		alltimekeysforthisroute = append(alltimekeysforthisroute, tks.TimeKey...)
		fmt.Println(k.SprintId)
	}

	var timecount int8

	for i, x := range alltimekeysforthisroute {
		timecount += alltimekeysforthisroute[i].TimeValueDown
		if timecount > int8(timelapsedfromsprintstart) {
			currentStation = x.Center
		}
	}

	timekeyforcurrentstation := TimeKey{Center: currentStation}

	var wait int64

	for i := timekeyforcurrentstation.KeyId; i <= timekeyforthisstation.KeyId; i++ {
		key := &TimeKey{KeyId: i}
		wait += int64(key.TimeValueDown)
	}

	return wait
}

type Node struct {
	NodeStation *Station
	Time        time.Time
}

func NodeCreate(station *Station, timer *time.Time) *Node {
	wait := GetTime(station, timer)
	times := timer.Add(time.Minute * time.Duration(wait))

	nd := &Node{NodeStation: station, Time: times}
	return nd
}
