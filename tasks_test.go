package habitica_test

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/wfernandes/go-habitica"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	mux    *http.ServeMux
	ts     *httptest.Server
	client *habitica.HabitClient
)

var _ = Describe("Tasks", func() {

	BeforeEach(func() {
		var err error
		mux = http.NewServeMux()
		ts = httptest.NewServer(mux)
		client, err = habitica.New(
			"userid",
			"api",
			habitica.WithBaseURL(ts.URL),
		)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		ts.Close()
	})

	It("returns a task for get", func() {
		mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {

		})

		task, err := client.Tasks.Get(context.Background())
		Expect(err).ToNot(HaveOccurred())
		Expect(task).ToNot(BeNil())
	})

})
