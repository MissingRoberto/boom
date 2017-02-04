package boom_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBoom(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Boom Suite")
}
