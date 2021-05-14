// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package ruler

import (
	"github.com/globocom/gitlab-lint/rules"
	"github.com/xanzy/go-gitlab"
)

type NonFastForwardMerge struct {
	rules.RulerImpl
}

func (w *NonFastForwardMerge) Run(c *gitlab.Client, p *gitlab.Project) bool {
	return p.MergeMethod != gitlab.FastForwardMerge
}

func NewNonFastForwardMerge() rules.Ruler {
	w := new(NonFastForwardMerge)
	w.ID = "non-fast-forward-merge"
	w.Name = "Non Fast-forward Merge"
	w.Level = rules.LevelPedantic
	return w
}
