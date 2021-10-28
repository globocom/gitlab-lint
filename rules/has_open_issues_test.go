// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules_test

import (
	. "github.com/globocom/gitlab-lint/rules"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xanzy/go-gitlab"
)

var _ = Describe("Has Open Issues Test", func() {
	BeforeEach(func() {
	})

	Context("Test rules", func() {
		It("GetLevel should return warning", func() {
			// Arrange

			// Act
			rule := NewHasOpenIssues()

			// Assert
			Expect(rule.GetLevel()).To(Equal("pedantic"))
		})

		It("GetSlug should return has-open-issues", func() {
			// Arrange

			// Act
			rule := NewHasOpenIssues()

			// Assert
			Expect(rule.GetSlug()).To(Equal("has-open-issues"))
		})

		It("Run should return true when the repo has at least one open issue", func() {
			// Arrange
			c := &gitlab.Client{}
			p := &gitlab.Project{
				OpenIssuesCount: 1,
			}
			// Act
			rule := NewHasOpenIssues()

			// Assert
			Expect(rule.Run(c, p)).To(Equal(true))
		})

		It("Run should return false when the repo has no open issues", func() {
			// Arrange
			c := &gitlab.Client{}
			p := &gitlab.Project{
				OpenIssuesCount: 0,
			}
			// Act
			rule := NewHasOpenIssues()

			// Assert
			Expect(rule.Run(c, p)).To(Equal(false))
		})
	})
})
