// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules_test

import (
	. "github.com/globocom/gitlab-lint/rules"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Without Gitlab-CI Test", func() {
	BeforeEach(func() {
	})

	Context("Test rules", func() {
		It("GetLevel should return info", func() {
			// Arrange

			// Act
			rule := NewWithoutGitlabCI()

			// Assert
			Expect(rule.GetLevel()).To(Equal("info"))
		})

		It("GetSlug should return without-gitlab-ci", func() {
			// Arrange

			// Act
			rule := NewWithoutGitlabCI()

			// Assert
			Expect(rule.GetSlug()).To(Equal("without-gitlab-ci"))
		})
	})
})
