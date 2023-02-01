package en_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "En Suite")
}
