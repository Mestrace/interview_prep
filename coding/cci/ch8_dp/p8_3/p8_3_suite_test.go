package p8_3_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestP83(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "P83 Suite")
}
