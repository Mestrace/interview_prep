package p8_9_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestP89(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "P89 Suite")
}
