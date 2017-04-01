package integration_test

import (
	"context"
	"os"
	"testing"

	"github.com/wfernandes/go-habitica"
)

func TestIntegration_UsersTasks(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration")
	}
	userID := os.Getenv("USER_ID")
	apiToken := os.Getenv("API_TOKEN")

	client, err := habitica.New(userID, apiToken)
	if err != nil {
		t.Errorf("error creating client: %s", err)
	}
	tasks, err := client.Tasks.List(context.Background())
	if err != nil {
		t.Errorf("error retrieving list of tasks: %s", err)
	}
	if tasks == nil {
		t.Error("got nil task response")
	}

	if !tasks.Success && len(tasks.Data) == 0 {
		t.Error("request was not a success")
	}
}
