// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules

import "github.com/xanzy/go-gitlab"

type NonFastForwardMerge struct {
	Description string `json:"description"`
	ID          string `json:"ruleId"`
	Level       string `json:"level"`
	Name        string `json:"name"`
}

func (w *NonFastForwardMerge) Run(c *gitlab.Client, p *gitlab.Project) bool {
	return p.MergeMethod != gitlab.FastForwardMerge
}
func (w *NonFastForwardMerge) GetSlug() string {
	return "non-fast-forward-merge"
}

func (w *NonFastForwardMerge) GetLevel() string {
	return LevelPedantic
}

func NewNonFastForwardMerge() Ruler {
	w := &NonFastForwardMerge{
		Name:        "Non Fast-forward Merge",
		Description: "",
	}
	w.ID = w.GetSlug()
	w.Level = w.GetLevel()
	return w
}
