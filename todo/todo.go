package todo

import (
	"time"
)

type ToDo struct {
	ID            string    `json:"id"`
	Activity      string    `json:"activity"`
	ActivityOwner string    `json:"activityowner"`
	Creation      time.Time `json:"creation"`
	Expiration    time.Time `json:"expiration"`
	IsDone        bool      `json:"isdone"`
}
