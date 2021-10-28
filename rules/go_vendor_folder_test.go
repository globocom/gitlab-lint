// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules_test

import (
	. "github.com/globocom/gitlab-lint/rules"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xanzy/go-gitlab"
)

var _ = Describe("Go Vendor Folder Test", func() {
	BeforeEach(func() {
	})

	Context("Test rules", func() {
		It("GetLevel should return warning", func() {
			// Arrange

			// Act
			rule := NewGoVendorFolder()

			// Assert
			Expect(rule.GetLevel()).To(Equal("warning"))
		})

		It("GetSlug should return go-vendor-folder", func() {
			// Arrange

			// Act
			rule := NewGoVendorFolder()

			// Assert
			Expect(rule.GetSlug()).To(Equal("go-vendor-folder"))
		})

		It("Run should return false when the repo is empty", func() {
			// Arrange
			c := &gitlab.Client{}
			p := &gitlab.Project{
				EmptyRepo: true,
			}
			// Act
			rule := NewGoVendorFolder()

			// Assert
			Expect(rule.Run(c, p)).To(Equal(false))
		})
	})
})
