package rules

import (
	"encoding/base64"
	"strings"

	"github.com/xanzy/go-gitlab"
)

type WithoutDASThund struct {
	Description string `json:"description"`
	ID          string `json:"ruleId"`
	Level       string `json:"level"`
	Name        string `json:"name"`
}

func (w *WithoutDASThund) Run(c *gitlab.Client, p *gitlab.Project) bool {
	gf := &gitlab.GetFileOptions{
		Ref: gitlab.String(p.DefaultBranch),
	}
	file, _, err := c.RepositoryFiles.GetFile(
		p.PathWithNamespace, ".gitlab-ci.yml", gf,
	)
	if err != nil {
		return false
	}

	contents, err := base64.StdEncoding.DecodeString(file.Content)
	if err != nil {
		return false
	}

	return !strings.Contains(string(contents), "dasthund")
}

func (w *WithoutDASThund) GetSlug() string {
	return "without-dasthund"
}

func (w *WithoutDASThund) GetLevel() string {
	return LevelInfo
}

func NewWithoutDASThund() Ruler {
	w := &WithoutDASThund{
		Name:        "Without DASThund",
		Description: "Gitlab CI definitions does not contain the string 'dasthund' that initiate a Dynamic Security Testing",
	}
	w.ID = w.GetSlug()
	w.Level = w.GetLevel()
	return w
}
