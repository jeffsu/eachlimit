package eachlimit_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEachlimit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Eachlimit Suite")
}
