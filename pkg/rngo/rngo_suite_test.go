package rngo_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRNGO(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RNGO Suite")
}
