package p8_12_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestP812(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "P812 Suite")
}
