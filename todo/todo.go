package todo

import (
	"time"
)

type ToDo struct {
	//rileggere https://gocondor.github.io/docs/validation e capire se fa al caso mio, ossia mettere required e len su tutti gli input che lo necessitano
	Id            int       `json:"id" grom:"primary"`
	Activity      string    `json:"activity"`
	ActivityOwner string    `json:"activityowner"`
	Creation      time.Time `json:"creation"`
	Expiration    time.Time `json:"expiration"`
	IsDone        bool      `json:"isdone"`
}

// passo la reference della struttura così posso omettere la return
func (*ToDo) CreateToDo(activity string, activityowner string, expiration time.Time, isdone bool) ToDo {
	var todo ToDo
	todo.Activity = activity
	todo.ActivityOwner = activityowner
	todo.Expiration = expiration
	todo.IsDone = isdone
	return todo
}
