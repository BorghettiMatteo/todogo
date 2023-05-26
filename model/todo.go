package model

import (
	"time"
)

type ToDo struct {
	//rileggere https://gocondor.github.io/docs/validation e capire se fa al caso mio, ossia mettere required e len su tutti gli input che lo necessitano
	Id            int       `json:"id" gorm:"primaryKey"`
	Activity      string    `json:"activity"`
	ActivityOwner string    `json:"activityowner"`
	Creation      time.Time `json:"creation" gorm:"DEFAULT:CURRENT_TIMESTAMP"`
	Expiration    time.Time `json:"expiration"`
	IsDone        bool      `json:"isdone" gorm:"DEFAULT:FALSE"`
}

// passo la reference della struttura così posso omettere la return
func (*ToDo) CreateToDo(activity string, activityowner string, expiration time.Time) ToDo {
	var todo ToDo
	todo.Activity = activity
	todo.ActivityOwner = activityowner
	todo.Expiration = expiration
	return todo
}

// questo override è necessario per specificare il puntamento alla tabella da dovere andare a scriverci sopra

func (t *ToDo) TableName() string {
	// custom table name, this is default
	return "public.todo"
}
