package habitica

import "context"

type Task struct{}

type TaskResponse struct {
	Success bool `json:"success"`
	Data    Task `json:"data"`
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
	// req, err := t.client.NewRequest()
	// // TODO: Handle err
	// _ = err
	// resp, err := t.client.Do(ctx, req)
	// // TODO: Handle err and resp
	// _ = err
	// _ = resp
	return &TaskResponse{Success: true}, nil
}
