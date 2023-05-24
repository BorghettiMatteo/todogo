package todo

type ToDo struct {
	ID        string    `json:"id"`
	Activity  string    `json:"activity"`
	IsDone    bool      `json:"isdone"`
	StartTime Time.time `json:"time"`
	Owner     string    `json:"owner"`
}
