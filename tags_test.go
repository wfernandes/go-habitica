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

var tagsResponse = []byte(`
{
    "success": true,
    "data": {
        "name": "practicetag",
        "id": "8bc0afbf-ab8e-49a4-982d-67a40557ed1a"
    },
    "notifications": []
}`)
