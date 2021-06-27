// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules

import (
	"sync"

	"github.com/xanzy/go-gitlab"
)

var MyRegistry = &Registry{
	Projects: map[string]Project{},
	Rules:    []Rule{},
	RulesFn:  map[string]Ruler{},
}

type Registry struct {
	mu       sync.Mutex
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

func (r *Registry) ProcessProject(c *gitlab.Client, p *gitlab.Project, ruler Ruler) {
	result := ruler.Run(c, p)

	r.mu.Lock()
	defer r.mu.Unlock()

	rule := NewRule(p, ruler)
	if result {
		r.Rules = append(r.Rules, rule)
	}

	if _, ok := r.Projects[p.PathWithNamespace]; !ok {
		newRules := map[string]int{rule.Level: 0}
		project := Project{Project: p, Rules: newRules}
		r.Projects[project.PathWithNamespace] = project
	}

	inc := 0
	if result {
		inc = 1
	}
	r.Projects[p.PathWithNamespace].Rules[rule.Level] += inc

}
