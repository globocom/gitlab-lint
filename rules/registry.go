// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules

import "github.com/xanzy/go-gitlab"

type Registry struct {
	Projects map[string]Project
	Rules    []Rule
	RulesFn  map[string]Ruler
}

func (r *Registry) AddRule(ruler Ruler) {
	if _, ok := r.RulesFn[ruler.GetSlug()]; ok {
		return
	}
	r.RulesFn[ruler.GetSlug()] = ruler
}

func (r *Registry) ProcessProject(c *gitlab.Client, p *gitlab.Project, ruler Ruler) bool {
	result := ruler.Run(c, p)
	if !result {
		return false
	}
	rule := NewRule(p, ruler)
	r.Rules = append(r.Rules, rule)

	if _, ok := r.Projects[p.PathWithNamespace]; !ok {
		newRules := map[string]int{rule.Level: 1}
		project := Project{Project: p, Rules: newRules}
		r.Projects[project.PathWithNamespace] = project
		return true
	}

	projects := r.Projects[p.PathWithNamespace]
	projects.Rules[rule.Level] += 1

	return true
}
