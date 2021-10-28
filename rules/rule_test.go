// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules_test

import (
	. "github.com/globocom/gitlab-lint/rules"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xanzy/go-gitlab"
)

var _ = Describe("Rule Test", func() {
	BeforeEach(func() {
	})

	It("NewRule should return a filled rule", func() {
		// Arrange
		p := &gitlab.Project{
			ID:                1,
			Description:       "project-1",
			PathWithNamespace: "/gitlab/project-1",
			Path:              "/gitlab",
			WebURL:            "http://gitlab.com/",
			Namespace: &gitlab.ProjectNamespace{
				ID:   12345,
				Path: "/project-1",
			},
		}
		r := newMockRule("rule-1", "warning", true)
		// Act
		rule := NewRule(p, r)

		// Assert
		Expect(rule).To(Equal(Rule{
			Description:       "project-1",
			NamespaceID:       12345,
			NamespacePath:     "/project-1",
			Path:              "/gitlab",
			PathWithNamespace: "/gitlab/project-1",
			ProjectID:         1,
			RuleID:            "rule-1",
			Level:             "warning",
			WebURL:            "http://gitlab.com/",
		}))
		Expect(rule.GetCollectionName()).To(Equal("rules"))
		Expect(rule.GetSearchableFields()).To(BeNil())
		Expect(rule.Cast()).To(Equal(&rule))
	})

})
