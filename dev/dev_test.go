package dev_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// . "github/desteves/stackreference-mock/dev"
)

var _ = Describe("Dev", func() {
	Context("When the x set\n", func() {
		It("Should return the x", func() {
			Expect(1).To(Equal(1))
		})
	})
})
