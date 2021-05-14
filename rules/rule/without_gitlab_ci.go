// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package ruler

import (
	"github.com/globocom/gitlab-lint/rules"
	"github.com/xanzy/go-gitlab"
)

type WithoutGitlabCI struct {
	rules.RulerImpl
}

func (w *WithoutGitlabCI) Run(c *gitlab.Client, p *gitlab.Project) bool {
	gf := &gitlab.GetFileOptions{
		Ref: gitlab.String(p.DefaultBranch),
	}
	_, _, err := c.RepositoryFiles.GetFile(
		p.PathWithNamespace, ".gitlab-ci.yml", gf,
	)
	// 404
	return err != nil
}

func NewWithoutGitlabCI() rules.Ruler {
	w := new(WithoutGitlabCI)
	w.ID = "without-gitlab-ci"
	w.Name = "Without Gitlab CI"
	w.Level = rules.LevelInfo
	return w
}
