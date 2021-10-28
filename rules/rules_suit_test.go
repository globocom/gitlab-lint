package rules_test

import (
	"github.com/globocom/gitlab-lint/rules"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xanzy/go-gitlab"

	"testing"
)

func TestPrettifiers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rules Suite")
}

//Mocks
type MockRule struct {
	slug      string
	level     string
	runResult bool
}

func (m MockRule) Run(client *gitlab.Client, p *gitlab.Project) bool { return m.runResult }
func (m MockRule) GetSlug() string                                   { return m.slug }
func (m MockRule) GetLevel() string                                  { return m.level }

func newMockRule(slug string, level string, runResult bool) rules.Ruler {
	return MockRule{slug: slug, level: level, runResult: runResult}
}
