// Copyright (c) 2021, Pablo Aguilar
// Licensed under the BSD 3-Clause License

package rules

import (
	log "github.com/sirupsen/logrus"
	"github.com/xanzy/go-gitlab"
)

// GoVendorFolder is a rule to verify if a repository has or not the go vendor folder.
// It look at the repository root searching for the `go.mod` file first since a project without
// that file means that it doesn't use go modules. If there's a `go.mod` it'll search for a
// file called `modules.txt` inside a `vendor` folder. That's the pattern the projects that use
// go modules and vendor its dependencies follows.
// If "go.mod" and "vendor/modules.txt" exist this rule will return `true`.
type GoVendorFolder struct {
	Description string `json:"description"`
	ID          string `json:"ruleId"`
	Level       string `json:"level"`
	Name        string `json:"name"`
}

// NewGoVendorFolder returns an instance of GoVendorFolder with its attributes filled
func NewGoVendorFolder() Ruler {
	v := &GoVendorFolder{
		Name:        "Go Vendor Folder",
		Description: "This rule identifies if a repo has the vendor folder for a project that uses go modules",
	}
	v.ID = v.GetSlug()
	v.Level = v.GetLevel()
	return v
}

func (f *GoVendorFolder) Run(c *gitlab.Client, p *gitlab.Project) bool {
	if p.EmptyRepo {
		return false
	}

	hasGoMod, err := f.searchForGoModFile(p.ID, c)
	if err != nil {
		log.Errorf(`[%s] error searching for "go.mod" file: %s`, f.GetSlug(), err)
		return false
	}

	if !hasGoMod {
		return false
	}

	hasGoVendor, err := f.searchGoVendorModulesFile(p.ID, c)
	if err != nil {
		log.Errorf(`[%s] error searching for "vendor/modules.txt" file: %s`, f.GetSlug(), err)
		return false
	}

	return hasGoVendor
}

func (f *GoVendorFolder) searchForGoModFile(projectID int, c *gitlab.Client) (bool, error) {
	return f.searchForFile(
		projectID,
		"go.mod",
		"",
		c,
	)
}

func (f *GoVendorFolder) searchGoVendorModulesFile(projectID int, c *gitlab.Client) (bool, error) {
	return f.searchForFile(
		projectID,
		"modules.txt",
		"vendor",
		c,
	)
}

func (f *GoVendorFolder) searchForFile(
	projectID int,
	fileName string,
	path string,
	c *gitlab.Client,
) (bool, error) {
	listOpts := gitlab.ListTreeOptions{
		ListOptions: gitlab.ListOptions{
			Page: 1,
		},
		Recursive: gitlab.Bool(false),
		Path:      gitlab.String(path),
	}

	for {
		nodes, resp, err := c.Repositories.ListTree(projectID, &listOpts)
		if err != nil {
			return false, err
		}

		for _, node := range nodes {
			if node.Type == "blob" && node.Name == fileName {
				return true, nil
			}
		}

		if resp.CurrentPage >= resp.TotalPages {
			break
		}

		listOpts.Page = resp.NextPage
	}

	return false, nil
}

func (f *GoVendorFolder) GetSlug() string {
	return "go-vendor-folder"
}

func (f *GoVendorFolder) GetLevel() string {
	return LevelWarning
}

func (e *GoVendorFolder) GetName() string {
	return e.Name
}

func (e *GoVendorFolder) GetDescription() string {
	return e.Description
}
