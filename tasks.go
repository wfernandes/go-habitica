package habitica

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	Text      string `json:"text"`
	Type      string `json:"type"`
	Notes     string `json:"notes"`
	Completed bool   `json:"completed"`
}

type TaskResponse struct {
	Success bool   `json:"success"`
	Data    Task   `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

type TasksResponse struct {
	Success bool   `json:"success"`
	Data    []Task `json:"data,omitempty"`
}

type TaskService struct {
	client *HabiticaClient
}

func newTaskService(h *HabiticaClient) *TaskService {
	return &TaskService{
		client: h,
	}
}

func (t *TaskService) Get(ctx context.Context, id string) (*TaskResponse, error) {
	req, err := t.client.NewRequest(http.MethodGet, fmt.Sprintf("tasks/%s", id), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}

	resp, err := t.client.Do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("unable to perform request: %s", err)
	}
	defer resp.Body.Close()

	var taskResp TaskResponse
	err = json.NewDecoder(resp.Body).Decode(&taskResp)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response body: %s", err)
	}

	return &taskResp, err
}

func (t *TaskService) List(ctx context.Context) (*TasksResponse, error) {
	req, err := t.client.NewRequest(http.MethodGet, "tasks/user", nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}

	resp, err := t.client.Do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("unable to perform request: %s", err)
	}
	defer resp.Body.Close()

	var tasksResp TasksResponse
	err = json.NewDecoder(resp.Body).Decode(&tasksResp)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response body: %s", err)
	}

	return &tasksResp, err
}

func (t *TaskService) Update(ctx context.Context, id string, task *Task) (*TaskResponse, error) {
	req, err := t.client.NewRequest(http.MethodPut, fmt.Sprintf("tasks/%s", id), task)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}
	resp, err := t.client.Do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("unable to perform request: %s", err)
	}
	defer resp.Body.Close()

	taskResp := &TaskResponse{}
	err = json.NewDecoder(resp.Body).Decode(taskResp)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response body: %s", err)
	}

	return taskResp, err
}

func (t *TaskService) Create(ctx context.Context, task *Task) (*TaskResponse, error) {
	req, err := t.client.NewRequest(http.MethodPost, "tasks/user", task)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}
	resp, err := t.client.Do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("unable to perform request: %s", err)
	}
	defer resp.Body.Close()

	taskResp := &TaskResponse{}
	err = json.NewDecoder(resp.Body).Decode(taskResp)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response body: %s", err)
	}

	return taskResp, err
}

func (t *TaskService) Delete(ctx context.Context, id string) (*TaskResponse, error) {
	req, err := t.client.NewRequest(http.MethodDelete, fmt.Sprintf("tasks/%s", id), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}
	resp, err := t.client.Do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("unable to perform request: %s", err)
	}
	defer resp.Body.Close()

	taskResp := &TaskResponse{}
	err = json.NewDecoder(resp.Body).Decode(taskResp)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response body: %s", err)
	}

	return taskResp, err
}
