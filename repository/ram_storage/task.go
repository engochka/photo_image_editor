package ram_storage

import (
	"errors"
	"homework-engochka/domain"
	"homework-engochka/repository"
)

type Task struct {
	data map[string]domain.Task
}

func NewTask() *Task {
	return &Task{
		data: make(map[string]domain.Task),
	}
}

func (t *Task) Post(id string, td domain.Task) error {
	if _, exists := t.data[id]; exists {
		return errors.New("id already exists")
	}
	t.data[id] = td
	return nil
}

func (t *Task) Put(id string, td domain.Task) error {
	_, exists := t.data[id]
	if !exists {
		return repository.ErrNotFound
	}
	t.data[id] = td
	return nil
}

func (t *Task) GetStatus(id string) (*string, error) {
	task, exists := t.data[id]
	if !exists {
		return nil, repository.ErrNotFound
	}
	return &task.Status, nil

}

func (t *Task) GetResult(id string) (*string, error) {
	task, exists := t.data[id]
	if !exists {
		return nil, repository.ErrNotFound
	}
	return &task.Result, nil
}
