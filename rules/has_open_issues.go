// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules

import (
	"github.com/xanzy/go-gitlab"
)

type HasOpenIssues struct {
	Description string `json:"description"`
	ID          string `json:"ruleId"`
	Level       string `json:"level"`
	Name        string `json:"name"`
}

func (h *HasOpenIssues) Run(c *gitlab.Client, p *gitlab.Project) bool {
	return p.OpenIssuesCount > 0
}

func (h *HasOpenIssues) GetSlug() string {
	return "has-open-issues"
}

func (h *HasOpenIssues) GetLevel() string {
	return LevelPedantic
}

func (e *HasOpenIssues) GetName() string {
	return e.Name
}

func (e *HasOpenIssues) GetDescription() string {
	return e.Description
}

func NewHasOpenIssues() Ruler {
	h := &HasOpenIssues{
		Name:        "Has Open Issues",
		Description: "",
	}
	h.ID = h.GetSlug()
	h.Level = h.GetLevel()
	return h
}
