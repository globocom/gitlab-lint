// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules_test

import (
	. "github.com/globocom/gitlab-lint/rules"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License
var _ = Describe("Stats Test", func() {
	BeforeEach(func() {
	})

	It("Stats methods should return as expected", func() {
		// Arrange

		// Act
		stats := Stats{}

		// Assert
		Expect(stats.GetCollectionName()).To(Equal("statistics"))
		Expect(stats.GetSearchableFields()).To(BeNil())
		Expect(stats.Cast()).To(Equal(&stats))
	})

})
