package rules_test

import (
	. "github.com/globocom/gitlab-lint/rules"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Without DASThund Test", func() {
	BeforeEach(func() {
	})

	Context("Test rules", func() {
		It("GetLevel should return info", func() {
			rule := NewWithoutDASThund()

			Expect(rule.GetLevel()).To(Equal("info"))
		})

		It("GetSlug should return without-dasthund", func() {
			rule := NewWithoutDASThund()

			Expect(rule.GetSlug()).To(Equal("without-dasthund"))
		})
	})
})
