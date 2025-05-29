package http

import (
	"homework-engochka/api/http/types"
	"homework-engochka/usecases"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Task represents an HTTP handler for managing tasks
type Task struct {
	service usecases.Task
}

// NewHandler creates a new instance of Task
func NewHandler(service usecases.Task) *Task {
	return &Task{service: service}
}

// @Summary Get a Status of the Task
// @Description Get an Status by its id
// @Tags object
// @Accept  json
// @Produce json
// @Param task_id query string true "Id of the task"
// @Success 200 {object} types.GetStatusHandlerResponse
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Task not found"
// @Router /status [get]
func (s *Task) getStatusHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateGetStatusHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	stat, err := s.service.GetStatus(req.Id)
	types.ProcessError(w, err, &types.GetStatusHandlerResponse{Status: stat})
}

// @Summary Get a Result of the Task
// @Description Get an Result by its id
// @Tags object
// @Accept  json
// @Produce json
// @Param task_id query string true "Id of the task"
// @Success 200 {object} types.GetResultHandlerResponse
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Task not found"
// @Router /result [get]
func (s *Task) getResultHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateGetResultHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	res, err := s.service.GetResult(req.Id)
	types.ProcessError(w, err, &types.GetResultHandlerResponse{Result: res})
}

// @Summary Create a Task
// @Description Create a new task with the specified uuid64
// @Tags object
// @Accept  json
// @Produce json
// @Param request body domain.Task true "Task data"
// @Success 200 {string} string "Task created successfully"
// @Failure 400 {string} string "Bad request"
// @Router /task [post]
func (s *Task) postHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePostHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	id, err := s.service.Post(*req)
	types.ProcessError(w, err, &types.PostHandlerResponse{Id: id})
}

// WithObjectHandlers registers object-related HTTP handlers
func (s *Task) WithTaskHandlers(r chi.Router) {
	r.Route("/", func(r chi.Router) {
		r.Get("/status", s.getStatusHandler)
		r.Get("/result", s.getResultHandler)
		r.Post("/task", s.postHandler)
	})
}
