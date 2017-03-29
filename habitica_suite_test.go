package habitica_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoHabitica(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoHabitica Suite")
}
