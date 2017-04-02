package habitica

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL   = "https://habitica.com/api/v3"
	UserAgent = "go-habitica/1" // 1 is the version
)

type HabiticaClient struct {
	userID    string
	apiToken  string
	BaseURL   string
	UserAgent string
	Client    *http.Client

	Tasks *TaskService
}

type ClientOpt func(*HabiticaClient)

func New(userID, apiToken string, opts ...ClientOpt) (*HabiticaClient, error) {
	if len(userID) == 0 || len(apiToken) == 0 {
		return nil, errors.New("needs valid user id and api token")
	}

	h := &HabiticaClient{
		userID:    userID,
		apiToken:  apiToken,
		BaseURL:   baseURL,
		UserAgent: UserAgent,
		Client:    http.DefaultClient,
	}

	for _, o := range opts {
		o(h)
	}

	h.Tasks = newTaskService(h)

	return h, nil
}

func WithBaseURL(baseUrl string) func(*HabiticaClient) {
	return func(h *HabiticaClient) {
		h.BaseURL = baseUrl
	}
}

func WithHttpClient(c *http.Client) func(*HabiticaClient) {
	return func(h *HabiticaClient) {
		h.Client = c
	}
}

func (h *HabiticaClient) NewRequest(method, urlPath string, body interface{}) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s", h.BaseURL, urlPath)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-user", h.userID)
	req.Header.Set("x-api-key", h.apiToken)

	return req, nil
}

func (h *HabiticaClient) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	return h.Client.Do(req)
}
