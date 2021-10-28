// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules_test

import (
	"github.com/globocom/gitlab-lint/rules"
	. "github.com/globocom/gitlab-lint/rules"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xanzy/go-gitlab"
)

var _ = Describe("Registry Test", func() {
	BeforeEach(func() {
	})

	It("AddRule should add rules to RulesFn with its slug as key", func() {
		// Arrange
		registry := NewRegistry()
		rule1 := newMockRule("rule-1", "warning", true)
		rule2 := newMockRule("rule-2", "error", true)
		rule3 := newMockRule("rule-3", "warning", true)

		// Act
		registry.AddRule(rule1)
		registry.AddRule(rule2)
		registry.AddRule(rule3)

		// Assert
		Expect(registry.RulesFn).To(HaveLen(3))
		Expect(registry.RulesFn["rule-1"]).To(Equal(rule1))
		Expect(registry.RulesFn["rule-2"]).To(Equal(rule2))
		Expect(registry.RulesFn["rule-3"]).To(Equal(rule3))
	})

	It("ProcessProject should process the project with the respectives rules", func() {
		// Arrange
		registry := NewRegistry()
		rule1 := newMockRule("rule-1", "warning", true)
		c := &gitlab.Client{}
		p := &gitlab.Project{
			ID:                1,
			Description:       "project-1",
			PathWithNamespace: "/gitlab/project-1",
			Namespace: &gitlab.ProjectNamespace{
				ID:   12345,
				Path: "/project-1",
			},
		}

		// Act
		registry.ProcessProject(c, p, rule1)

		// Assert
		Expect(registry.Rules).To(Equal([]rules.Rule{{
			Description:       "project-1",
			Level:             "warning",
			NamespaceID:       12345,
			NamespacePath:     "/project-1",
			Path:              "",
			PathWithNamespace: "/gitlab/project-1",
			ProjectID:         1,
			RuleID:            "rule-1",
			WebURL:            "",
		}}))
		Expect(registry.Projects["/project-1"]).To(Not(BeNil()))
	})

})
