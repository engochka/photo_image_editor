package types

import (
	"encoding/json"
	"fmt"
	"homework-engochka/domain"
	"net/http"
)

type GetStatusHandlerRequest struct {
	Id string `json:"task_id"`
}

func CreateGetStatusHandlerRequest(r *http.Request) (*GetStatusHandlerRequest, error) {
	id := r.URL.Query().Get("task_id")
	if id == "" {
		return nil, fmt.Errorf("missing id")
	}
	return &GetStatusHandlerRequest{Id: id}, nil
}

type GetResultHandlerRequest struct {
	Id string `json:"task_id"`
}

func CreateGetResultHandlerRequest(r *http.Request) (*GetResultHandlerRequest, error) {
	id := r.URL.Query().Get("task_id")
	if id == "" {
		return nil, fmt.Errorf("missing id")
	}
	return &GetResultHandlerRequest{Id: id}, nil
}

type PostHandlerRequest struct {
	domain.Task
}

func CreatePostHandlerRequest(r *http.Request) (*domain.Task, error) {
	var req domain.Task
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("error while decoding json: %v", err)
	}
	return &req, nil
}

type GetStatusHandlerResponse struct {
	Status *string `json:"status"`
}

type GetResultHandlerResponse struct {
	Result *string `json:"result"`
}

type PostHandlerResponse struct {
	Id *string `json:"task_id"`
}
