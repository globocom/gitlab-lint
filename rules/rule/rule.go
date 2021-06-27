// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package ruler

import (
	"github.com/globocom/gitlab-lint/rules"
)

func init() {
	rules.MyRegistry.AddRule(NewEmptyRepository())
	rules.MyRegistry.AddRule(NewHasOpenIssues())
	rules.MyRegistry.AddRule(NewLastActivity())
	rules.MyRegistry.AddRule(NewNonFastForwardMerge())
	rules.MyRegistry.AddRule(NewWithoutGitlabCI())
	rules.MyRegistry.AddRule(NewWithoutReadme())
}
