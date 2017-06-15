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
		w.Write(tagResponse)
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
		w.Write(tagResponse)
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
		w.Write(tagResponse)
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

func TestGet_UserTags(t *testing.T) {
	RegisterTestingT(t)
	setup()
	defer teardown()

	request := &http.Request{}
	mux.HandleFunc("/tags", func(w http.ResponseWriter, r *http.Request) {
		request = r
		w.WriteHeader(http.StatusOK)
		w.Write(userTagsResponse)
	})
	resp, err := client.Tags.List(ctx)
	Expect(err).ToNot(HaveOccurred())
	Expect(request.Method).To(Equal(http.MethodGet))

	Expect(resp.Data).ToNot(BeNil())
	tags := resp.Data
	Expect(tags).To(HaveLen(3))
}

func TestPost_ReorderTags(t *testing.T) {
	RegisterTestingT(t)
	setup()
	defer teardown()

	request := &http.Request{}
	receivedReorderTag := habitica.ReorderTag{}
	mux.HandleFunc("/reorder-tags", func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&receivedReorderTag)
		request = r
		w.WriteHeader(http.StatusOK)
		w.Write(reorderTagResponse)
	})

	reorderTag := &habitica.ReorderTag{
		TagID: "c6855fae-ca15-48af-a88b-86d0c65ead47",
		To:    4,
	}
	_, err := client.Tags.Reorder(ctx, reorderTag)
	Expect(err).ToNot(HaveOccurred())
	Expect(request.Method).To(Equal(http.MethodPost))
	Expect(receivedReorderTag.TagID).To(Equal("c6855fae-ca15-48af-a88b-86d0c65ead47"))
}

func TestUpdate_Tag(t *testing.T) {
	RegisterTestingT(t)
	setup()
	defer teardown()

	request := &http.Request{}
	updateTag := &habitica.Tag{Name: "Update Tag"}
	receivedTag := habitica.Tag{}
	mux.HandleFunc("/tags/some-tag-id", func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&receivedTag)
		request = r
		w.WriteHeader(http.StatusOK)
		w.Write(tagResponse)
	})
	_, err := client.Tags.Update(ctx, "some-tag-id", updateTag)
	Expect(err).ToNot(HaveOccurred())
	Expect(request.Method).To(Equal(http.MethodPut))
	Expect(receivedTag.Name).To(Equal("Update Tag"))
}

var tagResponse = []byte(`
{
    "success": true,
    "data": {
        "name": "practicetag",
        "id": "8bc0afbf-ab8e-49a4-982d-67a40557ed1a"
    },
    "notifications": []
}`)

var userTagsResponse = []byte(`
{
    "success": true,
    "data": [
        {
            "name": "Work",
            "id": "3d5d324d-a042-4d5f-872e-0553e228553e"
        },
        {
            "name": "apitester",
            "challenge": "true",
            "id": "f23c12f2-5830-4f15-9c36-e17fd729a812"
        },
        {
            "name": "practicetag",
            "id": "8bc0afbf-ab8e-49a4-982d-67a40557ed1a"
        }
    ],
    "notifications": []
} `)

var reorderTagResponse = []byte(`
{
    "tagId": "c6855fae-ca15-48af-a88b-86d0c65ead47",
    "to": 4
}`)
