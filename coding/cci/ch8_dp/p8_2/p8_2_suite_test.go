package p8_2_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestP82(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "P82 Suite")
}
