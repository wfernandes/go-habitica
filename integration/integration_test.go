package integration_test

import (
	"context"
	"os"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/wfernandes/go-habitica"
)

func TestIntegration_CRUDTask(t *testing.T) {
	RegisterTestingT(t)
	if testing.Short() {
		t.Skip("skipping integration")
	}

	userID := os.Getenv("USER_ID")
	apiToken := os.Getenv("API_TOKEN")

	client, err := habitica.New(userID, apiToken)
	Expect(err).ToNot(HaveOccurred(), "error creating client")

	// create a task
	ctx := context.Background()
	taskResponse := &habitica.TaskResponse{}
	task := &habitica.Task{
		Text:  "This is a test task",
		Notes: "Integration test task notes",
		Type:  "todo",
	}
	taskResponse, err = client.Tasks.Create(ctx, task)
	Expect(err).ToNot(HaveOccurred())
	Expect(taskResponse.Success).To(BeTrue(), "create did not succeed")
	task = &taskResponse.Data
	Expect(task.ID).ToNot(BeEmpty(), "task ID was empty")
	Expect(task.UserID).To(Equal(userID))

	//update the task
	task.Text = "Task has been updated"
	taskResponse, err = client.Tasks.Update(ctx, task.ID, task)
	Expect(err).ToNot(HaveOccurred())
	Expect(taskResponse.Success).To(BeTrue(), "update did not succeed")

	//get task
	_, err = client.Tasks.Get(ctx, task.ID)
	Expect(err).ToNot(HaveOccurred())
	Expect(task.Text).To(Equal("Task has been updated"))
	Expect(task.UserID).To(Equal(userID))

	//delete the task
	taskResponse, err = client.Tasks.Delete(ctx, task.ID)
	Expect(err).ToNot(HaveOccurred())
	Expect(taskResponse.Success).To(BeTrue(), "delete did not succeed")

	// confirm that the task has been deleted
	taskResponse, err = client.Tasks.Get(ctx, task.ID)
	Expect(err).ToNot(HaveOccurred())
	Expect(taskResponse.Success).To(BeFalse())
	Expect(taskResponse.Error).To(Equal("NotFound"))
	Expect(taskResponse.Message).To(Equal("Task not found."))

}

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
