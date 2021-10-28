// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules_test

import (
	. "github.com/globocom/gitlab-lint/rules"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Project Test", func() {
	BeforeEach(func() {
	})

	It("Project methods should return as expected", func() {
		// Arrange

		// Act
		project := Project{}

		// Assert
		Expect(project.GetCollectionName()).To(Equal("projects"))
		Expect(project.GetSearchableFields()).To(Equal([]string{"name", "weburl", "pathwithnamespace", "description"}))
		Expect(project.Cast()).To(Equal(&project))
	})

})
