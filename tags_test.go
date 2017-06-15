package habitica_test

import (
	"encoding/json"
	"net/http"
	"testing"

	. "github.com/onsi/gomega"
	habitica "github.com/wfernandes/go-habitica"
)

func TestCreate_Tag(t *testing.T) {
	RegisterTestingT(t)
	setup()
	defer teardown()

	request := &http.Request{}
	actualTag := &habitica.Tag{Name: "New Tag"}
	receivedTag := habitica.Tag{}
	mux.HandleFunc("/tags", func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&receivedTag)
		request = r
		w.WriteHeader(http.StatusOK)
		w.Write(tagsResponse)
	})
	_, err := client.Tags.Create(ctx, actualTag)
	Expect(err).ToNot(HaveOccurred())
	Expect(request.Method).To(Equal(http.MethodPost))
	Expect(receivedTag.Name).To(Equal("New Tag"))
}

func TestDelete_Tag(t *testing.T) {
	RegisterTestingT(t)
	setup()
	defer teardown()

	request := &http.Request{}
	mux.HandleFunc("/tags/some-tag-id", func(w http.ResponseWriter, r *http.Request) {
		request = r
		w.WriteHeader(http.StatusOK)
		w.Write(tagsResponse)
	})
	_, err := client.Tags.Delete(ctx, "some-tag-id")
	Expect(err).ToNot(HaveOccurred())
	Expect(request.Method).To(Equal(http.MethodDelete))
}

func TestGet_Tag(t *testing.T) {
	RegisterTestingT(t)
	setup()
	defer teardown()

	request := &http.Request{}
	mux.HandleFunc("/tags/some-tag-id", func(w http.ResponseWriter, r *http.Request) {
		request = r
		w.WriteHeader(http.StatusOK)
		w.Write(tagsResponse)
	})
	resp, err := client.Tags.Get(ctx, "some-tag-id")
	Expect(err).ToNot(HaveOccurred())
	Expect(request.Method).To(Equal(http.MethodGet))
	Expect(request.UserAgent()).To(Equal(habitica.UserAgent))
	Expect(request.Header.Get("x-api-user")).To(Equal("b0413351-405f-416f-8787-947ec1c85199"))
	Expect(request.Header.Get("x-api-key")).To(Equal("api"))
	Expect(request.Header.Get("Content-Type")).To(Equal("application/json"))

	Expect(resp.Data).ToNot(BeNil())
	Expect(resp.Data.Name).To(Equal("practicetag"))
	Expect(resp.Data.ID).ToNot(BeEmpty())
}

var tagsResponse = []byte(`
{
    "success": true,
    "data": {
        "name": "practicetag",
        "id": "8bc0afbf-ab8e-49a4-982d-67a40557ed1a"
    },
    "notifications": []
}`)
