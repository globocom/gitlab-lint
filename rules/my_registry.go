// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules

var MyRegistry = NewRegistry()

func init() {
	MyRegistry.AddRule(NewEmptyRepository())
	MyRegistry.AddRule(NewGoVendorFolder())
	MyRegistry.AddRule(NewHasOpenIssues())
	MyRegistry.AddRule(NewLastActivity())
	MyRegistry.AddRule(NewNonFastForwardMerge())
	MyRegistry.AddRule(NewWithoutGitlabCI())
	MyRegistry.AddRule(NewWithoutReadme())
	MyRegistry.AddRule(NewWithoutMakefile())
}
