// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package ruler

import (
	"github.com/globocom/gitlab-lint/rules"
	"github.com/xanzy/go-gitlab"
)

type HasOpenIssues struct {
	rules.RulerImpl
}

func (h *HasOpenIssues) Run(c *gitlab.Client, p *gitlab.Project) bool {
	return p.OpenIssuesCount > 0
}

func NewHasOpenIssues() rules.Ruler {
	h := new(HasOpenIssues)
	h.ID = "has-open-issues"
	h.Name = "Has Open Issues"
	h.Level = rules.LevelPedantic
	return h
}
