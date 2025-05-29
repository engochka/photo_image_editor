package repository

import "homework-engochka/domain"

type Task interface {
	GetStatus(id string) (*string, error)
	GetResult(id string) (*string, error)
	Post(id string, td domain.Task) error
	Put(id string, td domain.Task) error
}
