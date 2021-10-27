// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules

import "github.com/xanzy/go-gitlab"

type WithoutReadme struct {
	Description string `json:"description"`
	ID          string `json:"ruleId"`
	Level       string `json:"level"`
	Name        string `json:"name"`
}

func (w *WithoutReadme) Run(c *gitlab.Client, p *gitlab.Project) bool {
	return p.ReadmeURL == ""
}
func (w *WithoutReadme) GetSlug() string {
	return "without-readme"
}

func (w *WithoutReadme) GetLevel() string {
	return LevelError
}

func (e *WithoutReadme) GetName() string {
	return e.Name
}

func (e *WithoutReadme) GetDescription() string {
	return e.Description
}

func NewWithoutReadme() Ruler {
	w := &WithoutReadme{
		Name:        "Without Readme",
		Description: "",
	}
	w.ID = w.GetSlug()
	w.Level = w.GetLevel()
	return w
}
