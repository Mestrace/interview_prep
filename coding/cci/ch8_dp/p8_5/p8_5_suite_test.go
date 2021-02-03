package p8_5_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestP85(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "P85 Suite")
}
