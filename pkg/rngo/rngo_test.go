package rngo

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rngo", func() {
	Context("when creating a new Rngo instance", func() {
		It("should go boop!", func() {
			rngo := Boop()
			Expect(rngo).ToNot(BeNil())
		})
	})

})
