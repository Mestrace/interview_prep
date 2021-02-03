package p8_4_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestP84(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "P84 Suite")
}
