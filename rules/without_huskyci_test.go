// Copyright (c) 2021, Gustavo Covas
// Licensed under the BSD 3-Clause License

package rules_test

import (
	. "github.com/globocom/gitlab-lint/rules"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Without HuskyCI Test", func() {
	BeforeEach(func() {
	})

	Context("Test rules", func() {
		It("GetLevel should return info", func() {
			// Arrange

			// Act
			rule := NewWithoutHuskyCI()

			// Assert
			Expect(rule.GetLevel()).To(Equal("info"))
		})

		It("GetSlug should return without-husky-ci", func() {
			// Arrange

			// Act
			rule := NewWithoutHuskyCI()

			// Assert
			Expect(rule.GetSlug()).To(Equal("without-husky-ci"))
		})
	})
})
