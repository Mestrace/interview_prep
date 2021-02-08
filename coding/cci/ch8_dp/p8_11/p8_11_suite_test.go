package p8_11_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestP811(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "P811 Suite")
}
