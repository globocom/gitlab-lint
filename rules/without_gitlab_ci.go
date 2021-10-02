// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules

import "github.com/xanzy/go-gitlab"

type WithoutGitlabCI struct {
	Description string `json:"description"`
	ID          string `json:"ruleId"`
	Level       string `json:"level"`
	Name        string `json:"name"`
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

func (w *WithoutGitlabCI) GetSlug() string {
	return "without-gitlab-ci"
}

func (w *WithoutGitlabCI) GetLevel() string {
	return LevelInfo
}

func (e *WithoutGitlabCI) GetName() string {
	return e.Name
}

func (e *WithoutGitlabCI) GetDescription() string {
	return e.Description
}

func NewWithoutGitlabCI() Ruler {
	w := &WithoutGitlabCI{
		Name:        "Without Gitlab CI",
		Description: "",
	}
	w.ID = w.GetSlug()
	w.Level = w.GetLevel()
	return w
}
