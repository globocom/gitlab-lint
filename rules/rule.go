// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules

import "github.com/xanzy/go-gitlab"

type Rule struct {
	Description       string `json:"description" bson:"description"`
	Level             string `json:"level" bson:"level"`
	NamespaceID       int    `json:"namespaceId" bson:"namespaceId"`
	NamespacePath     string `json:"namespacePath" bson:"namespacePath"`
	Path              string `json:"path" bson:"path"`
	PathWithNamespace string `json:"pathWithNamespace" bson:"pathWithNamespace"`
	ProjectID         int    `json:"projectId" bson:"projectId"`
	RuleID            string `json:"ruleId" bson:"ruleId"`
	WebURL            string `json:"webUrl" bson:"webUrl"`
}

func (r Rule) Cast() Queryable {
	return &r
}

func (r Rule) GetCollectionName() string {
	return "rules"
}

func NewRule(p *gitlab.Project, r Ruler) Rule {
	return Rule{
		Description:       p.Description,
		NamespaceID:       p.Namespace.ID,
		NamespacePath:     p.Namespace.Path,
		Path:              p.Path,
		PathWithNamespace: p.PathWithNamespace,
		ProjectID:         p.ID,
		RuleID:            r.GetSlug(),
		Level:             r.GetLevel(),
		WebURL:            p.WebURL,
	}
}
