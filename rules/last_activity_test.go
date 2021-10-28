// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules_test

import (
	"time"

	. "github.com/globocom/gitlab-lint/rules"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xanzy/go-gitlab"
)

var _ = Describe("Last Activity Test", func() {
	BeforeEach(func() {
	})

	Context("Test rules", func() {
		It("GetLevel should return warning", func() {
			// Arrange

			// Act
			rule := NewLastActivity()

			// Assert
			Expect(rule.GetLevel()).To(Equal("warning"))
		})

		It("GetSlug should return last-activity-1-year", func() {
			// Arrange

			// Act
			rule := NewLastActivity()

			// Assert
			Expect(rule.GetSlug()).To(Equal("last-activity-1-year"))
		})

		It("Run should return false when the repo has last activity greater then 1 year", func() {
			// Arrange
			lastActivity := time.Now().Add(-1 * 24 * 366 * time.Hour)
			c := &gitlab.Client{}
			p := &gitlab.Project{
				LastActivityAt: &lastActivity,
			}
			// Act
			rule := NewLastActivity()

			// Assert
			Expect(rule.Run(c, p)).To(Equal(true))
		})

		It("Run should return false when the repo has last activity greater then 1 year", func() {
			// Arrange
			lastActivity := time.Now().Add(-1 * 24 * 364 * time.Hour)
			c := &gitlab.Client{}
			p := &gitlab.Project{
				LastActivityAt: &lastActivity,
			}
			// Act
			rule := NewLastActivity()

			// Assert
			Expect(rule.Run(c, p)).To(Equal(false))
		})
	})
})
