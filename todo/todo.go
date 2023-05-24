package todo

import "time"

type ToDo struct {
	ID        string    `json:"id"`
	Activity  string    `json:"activity"`
	IsDone    bool      `json:"isdone"`
	StartTime time.Time `json:"time"`
	Owner     string    `json:"owner"`
}
