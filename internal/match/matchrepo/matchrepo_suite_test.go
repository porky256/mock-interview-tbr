package matchrepo_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMatchrepo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Matchrepo Suite")
}
