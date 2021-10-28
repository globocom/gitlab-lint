// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules_test

import (
	. "github.com/globocom/gitlab-lint/rules"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xanzy/go-gitlab"
)

var _ = Describe("Fast Forward Merge Test", func() {
	BeforeEach(func() {
	})

	Context("Test rules", func() {
		It("GetLevel should return pedantic", func() {
			// Arrange

			// Act
			rule := NewNonFastForwardMerge()

			// Assert
			Expect(rule.GetLevel()).To(Equal("pedantic"))
		})

		It("GetSlug should return non-fast-forward-merge", func() {
			// Arrange

			// Act
			rule := NewNonFastForwardMerge()

			// Assert
			Expect(rule.GetSlug()).To(Equal("non-fast-forward-merge"))
		})

		It("Run should return true when merge method is not ff", func() {
			// Arrange
			c := &gitlab.Client{}
			p := &gitlab.Project{
				MergeMethod: "merge",
			}
			// Act
			rule := NewNonFastForwardMerge()

			// Assert
			Expect(rule.Run(c, p)).To(Equal(true))
		})

		It("Run should return false when merge method is ff", func() {
			// Arrange
			c := &gitlab.Client{}
			p := &gitlab.Project{
				MergeMethod: "ff",
			}
			// Act
			rule := NewNonFastForwardMerge()

			// Assert
			Expect(rule.Run(c, p)).To(Equal(false))
		})
	})
})
