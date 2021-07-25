package controllers

import (
	// "encoding/json"
	"tummoc/models"

	"encoding/json"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

type Coordinate struct {
	Lats  float64 `form:"latitude"`
	Longs float64 `form:"longitude"`
}

type DataInput struct {
	Source      *Coordinate `form:"source"`
	Destination *Coordinate `form:"destination"`
	Time        time.Time   `form:"time"`
}

type Node struct {
	NodeStation *models.Station
	Time        time.Time
}

type DataOutput struct {
	Start  *Node   `json:"start"`
	End    *Node   `json:"end"`
	Change []*Node `json:"interchanging_points"`
	Wait   int64   `json:"time_to_wait"`
}

type RouteController struct {
	beego.Controller
}

func (r *RouteController) DataEmiter() {
	ip := DataInput{}
	json.Unmarshal(r.Ctx.Input.RequestBody, &ip)
	if ip.Time.IsZero() {
		ip.Time = time.Now()
	}

	startstation := models.GetLocation((*models.Coordinate)(ip.Source))
	destinationstation := models.GetLocation((*models.Coordinate)(ip.Destination))

	waiter := models.GetTime(startstation, &ip.Time)
	startNode := models.NodeCreate(startstation, &ip.Time)
	endNode := models.NodeCreate(destinationstation, &ip.Time)

	op := &DataOutput{Start: (*Node)(startNode), End: (*Node)(endNode), Wait: waiter}

	r.Data["json"] = map[string]DataOutput{"Route": *op}
	r.ServeJSON()
}
