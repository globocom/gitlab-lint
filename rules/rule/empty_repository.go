// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package ruler

import (
	"github.com/globocom/gitlab-lint/rules"
	"github.com/xanzy/go-gitlab"
)

type EmptyRepository struct {
	rules.RulerImpl
}

func (e *EmptyRepository) Run(c *gitlab.Client, p *gitlab.Project) bool {
	return p.EmptyRepo
}

func NewEmptyRepository() rules.Ruler {
	e := new(EmptyRepository)
	e.ID = "empty-repository"
	e.Name = "Empty Repository"
	e.Level = rules.LevelError
	return e
}
