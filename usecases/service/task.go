package service

import (
	"homework-engochka/domain"
	"homework-engochka/repository"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	repo repository.Task
}

func NewTask(repo repository.Task) *Task {
	return &Task{
		repo: repo,
	}
}

func (t *Task) GetStatus(id string) (*string, error) {
	return t.repo.GetStatus(id)
}

func (t *Task) GetResult(id string) (*string, error) {
	return t.repo.GetResult(id)
}

func (t *Task) Post(td domain.Task) (*string, error) {
	uuid := uuid.New()
	Id := uuid.String()
	td.Status = "in_progress"
	td.Result = "begin"
	go t.scam(Id, td)
	return &Id, t.repo.Post(Id, td)
}

func (t *Task) scam(Id string, td domain.Task) {
	time.Sleep(66 * time.Second)
	td.Status = "ready"
	td.Result = "finish"
	t.repo.Put(Id, td)
}
