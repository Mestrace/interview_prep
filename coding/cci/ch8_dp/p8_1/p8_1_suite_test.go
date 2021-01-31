package p8_1_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestP81(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "P81 Suite")
}
