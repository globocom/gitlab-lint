// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules_test

import (
	. "github.com/globocom/gitlab-lint/rules"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xanzy/go-gitlab"
)

var _ = Describe("Without Readme Test", func() {
	BeforeEach(func() {
	})

	Context("Test rules", func() {
		It("GetLevel should return info", func() {
			// Arrange

			// Act
			rule := NewWithoutReadme()

			// Assert
			Expect(rule.GetLevel()).To(Equal("error"))
		})

		It("GetSlug should return without-readme", func() {
			// Arrange

			// Act
			rule := NewWithoutReadme()

			// Assert
			Expect(rule.GetSlug()).To(Equal("without-readme"))
		})

		It("Run should return true when readme exists", func() {
			// Arrange
			c := &gitlab.Client{}
			p := &gitlab.Project{
				ReadmeURL: "http://gitlab.com/project/readme.md",
			}
			// Act
			rule := NewWithoutReadme()

			// Assert
			Expect(rule.Run(c, p)).To(Equal(false))
		})

		It("Run should return true when readme does not exist", func() {
			// Arrange
			c := &gitlab.Client{}
			p := &gitlab.Project{
				ReadmeURL: "",
			}
			// Act
			rule := NewWithoutReadme()

			// Assert
			Expect(rule.Run(c, p)).To(Equal(true))
		})
	})
})
