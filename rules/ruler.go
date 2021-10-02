// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules

import "github.com/xanzy/go-gitlab"

type Ruler interface {
	Run(client *gitlab.Client, p *gitlab.Project) bool
	GetName() string
	GetDescription() string
	GetSlug() string
	GetLevel() string
}
