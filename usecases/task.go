package usecases

import "homework-engochka/domain"

type Task interface {
	GetStatus(id string) (*string, error)
	GetResult(id string) (*string, error)
	Post(td domain.Task) (*string, error)
}
