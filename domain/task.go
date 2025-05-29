package domain

type Task struct {
	Task   string `json:"task"`
	Status string `json:"-"`
	Result string `json:"-"`
}
