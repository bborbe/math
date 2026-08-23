package math

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Double", func() {
	It("doubles positive and negative inputs", func() {
		Expect(Double(2)).To(Equal(4))
		Expect(Double(-3)).To(Equal(-6))
	})
})
