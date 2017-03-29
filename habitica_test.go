package habitica_test

import (
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

	It("returns a client with userID and apiToken", func() {
		c, err := habitica.New("some-user-id", "some-api-token")
		Expect(err).ToNot(HaveOccurred())
		Expect(c).ToNot(BeNil())
	})

	It("configures client with default base url", func() {
		c, err := habitica.New("user", "api")
		Expect(err).ToNot(HaveOccurred())
		Expect(c.BaseURL).To(Equal("https://habitica.com/api/v3"))
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

})
