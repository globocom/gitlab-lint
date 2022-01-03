// Copyright (c) 2021, Gustavo Covas
// Licensed under the BSD 3-Clause License

package rules

import (
	"encoding/base64"
	"strings"

	"github.com/xanzy/go-gitlab"
)

type WithoutHuskyCI struct {
	Description string `json:"description"`
	ID          string `json:"ruleId"`
	Level       string `json:"level"`
	Name        string `json:"name"`
}

func (w *WithoutHuskyCI) Run(c *gitlab.Client, p *gitlab.Project) bool {
	gf := &gitlab.GetFileOptions{
		Ref: gitlab.String(p.DefaultBranch),
	}
	file, _, err := c.RepositoryFiles.GetFile(
		p.PathWithNamespace, ".gitlab-ci.yml", gf,
	)
	if err != nil {
		return false
	}

	contents, err := base64.StdEncoding.DecodeString(file.Content)
	if err != nil {
		return false
	}

	return !strings.Contains(string(contents), "huskyCI")
}

func (w *WithoutHuskyCI) GetSlug() string {
	return "without-husky-ci"
}

func (w *WithoutHuskyCI) GetLevel() string {
	return LevelInfo
}

func NewWithoutHuskyCI() Ruler {
	w := &WithoutHuskyCI{
		Name:        "Without Husky CI",
		Description: "Gitlab CI definitions does not contain the string 'huskyCI'",
	}
	w.ID = w.GetSlug()
	w.Level = w.GetLevel()
	return w
}
