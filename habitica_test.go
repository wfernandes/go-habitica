package habitica_test

import (
	"fmt"
	"net/http"
	"time"

	"github.com/wfernandes/go-habitica"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Habitica", func() {
	It("returns an error for empty userID or apiToken", func() {
		_, err := habitica.New("", "some-api-token")
		Expect(err).To(HaveOccurred())

		_, err = habitica.New("some-user-id", "")
		Expect(err).To(HaveOccurred())

	})

	It("accepts userID and apiToken", func() {
		c, err := habitica.New("some-user-id", "some-api-token")
		Expect(err).ToNot(HaveOccurred())
		Expect(c).ToNot(BeNil())
	})

	Context("configuration", func() {
		It("configures client with default base url", func() {
			c, err := habitica.New("user", "api")
			Expect(err).ToNot(HaveOccurred())
			Expect(c.BaseURL).To(Equal("https://habitica.com/api/v3"))
		})

		It("configures client with user agent", func() {
			c, err := habitica.New("user", "api")
			Expect(err).ToNot(HaveOccurred())
			Expect(c.UserAgent).To(Equal("go-habitica/1"))
		})

		It("allows for configurable base url", func() {
			c, err := habitica.New(
				"user",
				"api",
				habitica.WithBaseURL("https://somethingelse.com"),
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(c.BaseURL).To(Equal("https://somethingelse.com"))
		})

		It("configures client with default http client", func() {
			c, err := habitica.New("user", "api")
			Expect(err).ToNot(HaveOccurred())
			Expect(c.Client).ToNot(BeNil())
		})

		It("allows for configurable http client", func() {
			hc := &http.Client{Timeout: time.Second}
			c, err := habitica.New(
				"user",
				"api",
				habitica.WithHttpClient(hc),
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(c.Client).To(Equal(hc))
		})
	})

	Context("new request", func() {
		It("returns a request with appropriate headers", func() {
			c, err := habitica.New("user", "api")
			Expect(err).ToNot(HaveOccurred())
			urlPath := "tasks/group/some-group-id"
			request, err := c.NewRequest(http.MethodGet, urlPath)
			Expect(err).ToNot(HaveOccurred())
			Expect(request.Method).To(Equal(http.MethodGet))
			Expect(request.UserAgent()).To(Equal(habitica.UserAgent))
			Expect(request.URL.String()).To(Equal(fmt.Sprintf("%s/%s", c.BaseURL, urlPath)))
			Expect(request.Header.Get("x-api-user")).To(Equal("user"))
			Expect(request.Header.Get("x-api-key")).To(Equal("api"))
			Expect(request.Header.Get("Content-Type")).To(Equal("application/json"))
		})

		It("returns an error if unable to create request", func() {
			c, err := habitica.New("user", "api")
			Expect(err).ToNot(HaveOccurred())
			_, err = c.NewRequest(" GOT", "")
			Expect(err).To(HaveOccurred())
		})
	})

	Context("task service", func() {
		It("configures a task service", func() {
			c, err := habitica.New("user", "api")
			Expect(err).To(BeNil())
			Expect(c.Tasks).ToNot(BeNil())
		})
	})
})
