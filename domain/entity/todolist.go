package entity

type TodoList struct {
	ID     int64  `json:"id" db:"id"`
	Task   string `json:"task" db:"task"`
	Assign int64  `json:"assign" db:"assign"`
}
