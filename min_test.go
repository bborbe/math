// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/math"
)

var _ = Describe("Min", func() {
	var input []int
	var result *int
	BeforeEach(func() {
		input = []int{}
	})
	JustBeforeEach(func() {
		result = math.Min(input)
	})
	Context("empty", func() {
		BeforeEach(func() {
			input = []int{}
		})
		It("returns correct min", func() {
			Expect(result).To(BeNil())
		})
	})
	Context("with values", func() {
		BeforeEach(func() {
			input = []int{1, 2, 3}
		})
		It("returns correct min", func() {
			Expect(result).NotTo(BeNil())
			Expect(*result).To(Equal(1))
		})
	})
})
