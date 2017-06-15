package habitica

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	ID        string          `json:"id"`
	UserID    string          `json:"userId"`
	Text      string          `json:"text"`
	Type      string          `json:"type"`
	Notes     string          `json:"notes"`
	Tags      []string        `json:"tags"`
	Completed bool            `json:"completed"`
	Checklist []ChecklistItem `json:"checklist"`
}

type TaskResponse struct {
	Success bool   `json:"success"`
	Data    *Task  `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

type TaskReorderResponse struct {
	Success bool     `json:"success"`
	Data    []string `json:"data,omitempty"`
}

type TasksResponse struct {
	Success bool   `json:"success"`
	Data    []Task `json:"data,omitempty"`
}

type ChecklistItem struct {
	Id        string ` json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

type TaskService struct {
	client *HabiticaClient
}

func newTaskService(h *HabiticaClient) *TaskService {
	return &TaskService{
		client: h,
	}
}

func (t *TaskService) getTaskResponse(ctx context.Context, req *http.Request) (*TaskResponse, error) {
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

func (t *TaskService) Get(ctx context.Context, id string) (*TaskResponse, error) {
	req, err := t.client.NewRequest(http.MethodGet, fmt.Sprintf("tasks/%s", id), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}

	return t.getTaskResponse(ctx, req)
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

	return t.getTaskResponse(ctx, req)
}

func (t *TaskService) Create(ctx context.Context, task *Task) (*TaskResponse, error) {
	req, err := t.client.NewRequest(http.MethodPost, "tasks/user", task)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}

	return t.getTaskResponse(ctx, req)
}

func (t *TaskService) Delete(ctx context.Context, id string) (*TaskResponse, error) {
	req, err := t.client.NewRequest(http.MethodDelete, fmt.Sprintf("tasks/%s", id), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}

	return t.getTaskResponse(ctx, req)
}

func (t *TaskService) AddTag(ctx context.Context, taskID, tagID string) (*TaskResponse, error) {
	req, err := t.client.NewRequest(http.MethodPost, fmt.Sprintf("tasks/%s/tags/%s", taskID, tagID), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}

	return t.getTaskResponse(ctx, req)
}

func (t *TaskService) DeleteTag(ctx context.Context, taskID, tagID string) (*TaskResponse, error) {
	req, err := t.client.NewRequest(http.MethodDelete, fmt.Sprintf("tasks/%s/tags/%s", taskID, tagID), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}

	return t.getTaskResponse(ctx, req)
}

func (t *TaskService) AddChecklistItem(ctx context.Context, taskID string, item *ChecklistItem) (*TaskResponse, error) {
	req, err := t.client.NewRequest(http.MethodPost, fmt.Sprintf("tasks/%s/checklist", taskID), item)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}

	return t.getTaskResponse(ctx, req)
}

func (t *TaskService) UpdateChecklistItem(ctx context.Context, taskID string, item *ChecklistItem) (*TaskResponse, error) {
	req, err := t.client.NewRequest(http.MethodPut, fmt.Sprintf("tasks/%s/checklist/%s", taskID, item.Id), item)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}

	return t.getTaskResponse(ctx, req)
}

func (t *TaskService) DeleteChecklistItem(ctx context.Context, taskID, itemID string) (*TaskResponse, error) {
	req, err := t.client.NewRequest(http.MethodDelete, fmt.Sprintf("tasks/%s/checklist/%s", taskID, itemID), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}

	return t.getTaskResponse(ctx, req)
}

func (t *TaskService) ClearCompletedTodos(ctx context.Context) (*TaskResponse, error) {
	req, err := t.client.NewRequest(http.MethodPost, "tasks/clearcompletedtodos", nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}

	return t.getTaskResponse(ctx, req)
}

func (t *TaskService) MoveToPosition(ctx context.Context, taskID string, position int) (*TaskReorderResponse, error) {
	req, err := t.client.NewRequest(http.MethodPost, fmt.Sprintf("tasks/%s/move/to/%d", taskID, position), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}
	resp, err := t.client.Do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("unable to perform request: %s", err)
	}
	defer resp.Body.Close()
	taskReorderResp := &TaskReorderResponse{}
	err = json.NewDecoder(resp.Body).Decode(taskReorderResp)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response body: %s", err)
	}

	return taskReorderResp, err
}
