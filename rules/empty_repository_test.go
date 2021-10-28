// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules_test

import (
	. "github.com/globocom/gitlab-lint/rules"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xanzy/go-gitlab"
)

var _ = Describe("Empty Repository Test", func() {
	BeforeEach(func() {
	})

	Context("Test rules", func() {
		It("GetLevel should return error", func() {
			// Arrange

			// Act
			rule := NewEmptyRepository()

			// Assert
			Expect(rule.GetLevel()).To(Equal("error"))
		})

		It("GetSlug should return empty-repository", func() {
			// Arrange

			// Act
			rule := NewEmptyRepository()

			// Assert
			Expect(rule.GetSlug()).To(Equal("empty-repository"))
		})

		It("Run should return true when repo is empty", func() {
			// Arrange
			c := &gitlab.Client{}
			p := &gitlab.Project{
				EmptyRepo: true,
			}
			// Act
			rule := NewEmptyRepository()

			// Assert
			Expect(rule.Run(c, p)).To(Equal(true))
		})

		It("Run should return false when repo is not empty", func() {
			// Arrange
			c := &gitlab.Client{}
			p := &gitlab.Project{
				EmptyRepo: true,
			}
			// Act
			rule := NewEmptyRepository()

			// Assert
			Expect(rule.Run(c, p)).To(Equal(true))
		})
	})
})
