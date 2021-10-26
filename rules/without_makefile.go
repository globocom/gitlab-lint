// Copyright (c) 2021, Paulo Ricardo Koch
// Licensed under the BSD 3-Clause License

package rules

import (
	"github.com/xanzy/go-gitlab"
)

type WithoutMakefile struct {
	Description string `json:"description"`
	ID          string `json:"ruleId"`
	Level       string `json:"level"`
	Name        string `json:"name"`
}

func (w *WithoutMakefile) Run(c *gitlab.Client, p *gitlab.Project) bool {
	ref := gitlab.GetFileOptions{Ref: gitlab.String(p.DefaultBranch)}
	file, _, _ := c.RepositoryFiles.GetFile(p.ID, "Makefile", &ref)

	return file == nil
}

func (w *WithoutMakefile) GetSlug() string {
	return "without-makefile"
}

func (w *WithoutMakefile) GetLevel() string {
	return LevelInfo
}

func NewWithoutMakefile() Ruler {
	w := &WithoutMakefile{
		Name:        "Without Makefile",
		Description: "",
	}
	w.ID = w.GetSlug()
	w.Level = w.GetLevel()
	return w
}
