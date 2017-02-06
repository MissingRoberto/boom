package boom

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

const (
	completeManifestPath             string = "examples/manifest.yml"
	manifestWithoutResourcePoolsPath string = "examples/manifest-without-resource-pools.yml"
)

func TestBoom(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Boom Suite")
}
