// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package ruler

import (
	"github.com/globocom/gitlab-lint/rules"
	"github.com/xanzy/go-gitlab"
)

type WithoutReadme struct {
	rules.RulerImpl
}

func (w *WithoutReadme) Run(c *gitlab.Client, p *gitlab.Project) bool {
	return p.ReadmeURL == ""
}

func NewWithoutReadme() rules.Ruler {
	w := new(WithoutReadme)
	w.ID = "without-readme"
	w.Name = "Without Readme"
	w.Level = rules.LevelError
	return w
}
