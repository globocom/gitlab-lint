// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package ruler

import (
	"time"

	"github.com/globocom/gitlab-lint/rules"
	"github.com/xanzy/go-gitlab"
)

type LastActivity struct {
	rules.RulerImpl
}

func (l *LastActivity) Run(c *gitlab.Client, p *gitlab.Project) bool {
	t2 := time.Now()
	days := t2.Sub(*p.LastActivityAt).Hours() / 24
	return days > 365
}

func NewLastActivity() rules.Ruler {
	l := new(LastActivity)
	l.ID = "last-activity-1-year"
	l.Name = "Last Activity > 1 year"
	l.Level = rules.LevelWarning
	return l
}
