package habitica

import "errors"

const (
	baseURL = "https://habitica.com/api/v3"
)

type HabitClient struct {
	BaseURL string
}

type ClientOpt func(*HabitClient)

func New(userID, apiToken string, opts ...ClientOpt) (*HabitClient, error) {
	if len(userID) == 0 || len(apiToken) == 0 {
		return nil, errors.New("needs valid user id and api token")
	}

	client := &HabitClient{
		BaseURL: baseURL,
	}

	for _, o := range opts {
		o(client)
	}

	return client, nil
}

func WithBaseURL(url string) func(*HabitClient) {
	return func(h *HabitClient) {
		h.BaseURL = url
	}
}
