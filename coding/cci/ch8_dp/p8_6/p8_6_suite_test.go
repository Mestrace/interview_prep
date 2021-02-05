package p8_6_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestP86(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "P86 Suite")
}
