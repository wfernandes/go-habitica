package habitica_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/wfernandes/go-habitica"

	// . "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestValidation_RequiredArgs(t *testing.T) {
	RegisterTestingT(t)
	_, err := habitica.New("", "some-api-token")
	Expect(err).To(HaveOccurred())

	_, err = habitica.New("some-user-id", "")
	Expect(err).To(HaveOccurred())
}

func TestCreateClient(t *testing.T) {
	RegisterTestingT(t)
	c, err := habitica.New("some-user-id", "some-api-token")
	Expect(err).ToNot(HaveOccurred())
	Expect(c).ToNot(BeNil())
}

func TestConfigure_DefaultBaseUrl(t *testing.T) {
	RegisterTestingT(t)
	c, err := habitica.New("user", "api")
	Expect(err).ToNot(HaveOccurred())
	Expect(c.BaseURL).To(Equal("https://habitica.com/api/v3"))
}

func TestConfigure_BaseUrl(t *testing.T) {
	RegisterTestingT(t)
	c, err := habitica.New(
		"user",
		"api",
		habitica.WithBaseURL("https://somethingelse.com"),
	)
	Expect(err).ToNot(HaveOccurred())
	Expect(c.BaseURL).To(Equal("https://somethingelse.com"))
}

func TestConfigure_DefaultUserAgent(t *testing.T) {
	RegisterTestingT(t)
	c, err := habitica.New("user", "api")
	Expect(err).ToNot(HaveOccurred())
	Expect(c.UserAgent).To(Equal("go-habitica/1"))
}

func TestConfigure_DefaultHttpClient(t *testing.T) {
	RegisterTestingT(t)
	c, err := habitica.New("user", "api")
	Expect(err).ToNot(HaveOccurred())
	Expect(c.Client).ToNot(BeNil())
}

func TestConfigure_HttpClient(t *testing.T) {
	RegisterTestingT(t)
	hc := &http.Client{Timeout: time.Second}
	c, err := habitica.New(
		"user",
		"api",
		habitica.WithHttpClient(hc),
	)
	Expect(err).ToNot(HaveOccurred())
	Expect(c.Client).To(Equal(hc))
}

func TestConfigure_TaskService(t *testing.T) {
	RegisterTestingT(t)
	c, err := habitica.New("user", "api")
	Expect(err).To(BeNil())
	Expect(c.Tasks).ToNot(BeNil())
}

func TestNewRequest_CorrectHeaders(t *testing.T) {
	RegisterTestingT(t)
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
}

func TestNewRequest_ErrorForBadMethod(t *testing.T) {
	RegisterTestingT(t)
	c, err := habitica.New("user", "api")
	Expect(err).ToNot(HaveOccurred())
	_, err = c.NewRequest(" GOT", "")
	Expect(err).To(HaveOccurred())
}
