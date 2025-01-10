// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/math"
)

var _ = Describe("Diff", func() {
	It("float64 / float64", func() {
		result := math.Diff(float64(6), float64(2))
		Expect(result).NotTo(BeNil())
		Expect(*result).To(Equal(float64(3)))
	})
	It("float64 / int", func() {
		result := math.Diff(float64(6), 2)
		Expect(result).NotTo(BeNil())
		Expect(*result).To(Equal(float64(3)))
	})
})
