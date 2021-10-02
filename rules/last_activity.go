// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules

import (
	"time"

	"github.com/xanzy/go-gitlab"
)

type LastActivity struct {
	Description string `json:"description"`
	ID          string `json:"ruleId"`
	Level       string `json:"level"`
	Name        string `json:"name"`
}

func (l *LastActivity) Run(c *gitlab.Client, p *gitlab.Project) bool {
	t2 := time.Now()
	days := t2.Sub(*p.LastActivityAt).Hours() / 24
	return days > 365
}

func (l *LastActivity) GetSlug() string {
	return "last-activity-1-year"
}

func (l *LastActivity) GetLevel() string {
	return LevelWarning
}

func (e *LastActivity) GetName() string {
	return e.Name
}

func (e *LastActivity) GetDescription() string {
	return e.Description
}

func NewLastActivity() Ruler {
	l := &LastActivity{
		Name:        "Last Activity > 1 year",
		Description: "",
	}
	l.ID = l.GetSlug()
	l.Level = l.GetLevel()
	return l
}
