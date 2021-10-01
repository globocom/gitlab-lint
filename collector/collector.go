// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package main

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/globocom/gitlab-lint/config"
	"github.com/globocom/gitlab-lint/db"
	"github.com/globocom/gitlab-lint/rules"
)

func processProjects(projects map[string]rules.Project) error {
	dbInstance, err := db.NewMongoSession()
	if err != nil {
		log.Errorf("[Collector] Error on create mongo session: %v", err)
		return err
	}

	if _, err := dbInstance.DeleteMany(&rules.Project{}, bson.M{}); err != nil {
		return err
	}

	perInsert := viper.GetInt("db.perInsert")
	temp := make([]interface{}, 0)
	for _, p := range projects {
		temp = append(temp, p)
		if len(temp) >= perInsert {
			log.Debugf("[Collector] Inserting %d projects", perInsert)
			if _, err := dbInstance.InsertMany(&rules.Project{}, temp); err != nil {
				return err
			}
			temp = make([]interface{}, 0)
		}
	}
	log.Debugf("[Collector] Inserting %d projects", len(temp))
	if _, err := dbInstance.InsertMany(&rules.Project{}, temp); err != nil {
		return err
	}

	return nil
}

func processRules(rulesList []rules.Rule) error {
	dbInstance, err := db.NewMongoSession()
	if err != nil {
		log.Errorf("[Collector] Error on create mongo session: %v", err)
		return err
	}

	if _, err := dbInstance.DeleteMany(&rules.Rule{}, bson.M{}); err != nil {
		return err
	}

	perInsert := viper.GetInt("db.perInsert")
	temp := make([]interface{}, 0)
	for _, r := range rulesList {
		temp = append(temp, r)
		if len(temp) >= perInsert {
			log.Debugf("[Collector] Inserting %d rules", perInsert)
			if _, err := dbInstance.InsertMany(rules.Rule{}, temp); err != nil {
				return err
			}
			temp = make([]interface{}, 0)
		}
	}
	log.Debugf("[Collector] Inserting %d rules", len(temp))
	if _, err := dbInstance.InsertMany(rules.Rule{}, temp); err != nil {
		return err
	}

	return nil
}
func processIssues(registry *rules.Registry, git *gitlab.Client) error {
	dbInstance, err := db.NewMongoSession()
	if err != nil {
		log.Errorf("[Collector] Error on create mongo session: %v", err)
		return err
	}
	// iterating over matched rules
	for _, r := range registry.Rules {
		// searching for opened issue for project and rule
		pipeline := bson.M{"$and": bson.A{bson.M{"projectId": r.ProjectID}, bson.M{"ruleId": r.RuleID}, bson.M{"state": "opened"}}}

		issue := &rules.Issue{}
		err := dbInstance.Get(issue, pipeline, &options.FindOneOptions{})

		// if any error
		if err != nil && err != mongo.ErrNoDocuments {
			return err
		}
		// if no opened issue found -> create new issue and add to DB
		if err != nil && err == mongo.ErrNoDocuments {
			createdIssue, _, err := git.Issues.CreateIssue(r.ProjectID, &gitlab.CreateIssueOptions{Title: &r.RuleID})
			if err != nil {
				return err
			}

			if _, err := dbInstance.Insert(&rules.Issue{ProjectID: r.ProjectID, RuleID: r.RuleID, IssueID: createdIssue.ID,
				WebURL: r.WebURL, Title: r.RuleID, Description: "Test", State: createdIssue.State}); err != nil {
				return err
			}
			// Opened issue found
		} else {
			// fetch issue info from gitlab
			gitIssue, _, err := git.Issues.GetIssue(issue.ProjectID, issue.IssueID)
			if err != nil {
				return err
			}
			// if issue was closed but rule still matched
			if gitIssue.State != issue.State && gitIssue.State == "closed" {
				// update old issue on DB
				if _, err := dbInstance.Update(&rules.Issue{}, bson.M{"_id": issue.ID}, bson.M{"$set": bson.M{"state": gitIssue.State}}, &options.UpdateOptions{}); err != nil {
					return err
				}
				desc := "test-reopened"
				// create new issue
				createdIssue, _, err := git.Issues.CreateIssue(r.ProjectID, &gitlab.CreateIssueOptions{Title: &r.RuleID, Description: &desc})
				if err != nil {
					return err
				}
				// add new issue to DB
				if _, err := dbInstance.Insert(&rules.Issue{ProjectID: r.ProjectID, RuleID: r.RuleID, IssueID: createdIssue.ID,
					WebURL: r.WebURL, Title: r.RuleID, Description: "Test-Reopened", State: createdIssue.State}); err != nil {
					return err
				}

			}
		}
	}

	return nil
}

func insertStats(r *rules.Registry) error {
	dbInstance, err := db.NewMongoSession()
	if err != nil {
		log.Errorf("[Collector] Error on create mongo session: %v", err)
		return err
	}

	projectsCount, err := dbInstance.Count(&rules.Project{}, db.FindFilter{})
	if err != nil {
		return err
	}

	pipeline := []bson.M{{"$group": bson.M{
		"_id":   "$level",
		"count": bson.M{"$sum": 1},
	}}}

	levelsData, err := dbInstance.Aggregate(&rules.Rule{}, pipeline)
	if err != nil {
		return err
	}

	levelsCount := map[string]int32{}
	for _, level := range levelsData {
		levelsCount[level["_id"].(string)] = level["count"].(int32)
	}

	stats := &rules.Stats{
		CreatedAt:            time.Now().UTC(),
		GitlabProjectsCount:  projectsCount,
		LevelsCount:          levelsCount,
		ProjectsCount:        len(r.Projects),
		RegisteredRulesCount: len(r.RulesFn),
		RulesCount:           len(r.Rules),
	}

	log.Debug("[Collector] Inserting collector statistics")
	if _, err := dbInstance.Insert(stats); err != nil {
		return err
	}

	return nil
}

func worker(projects []*gitlab.Project, git *gitlab.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, project := range projects {
		// Only non-forks
		if project.ForkedFromProject != nil {
			continue
		}

		for _, rulesFn := range rules.MyRegistry.RulesFn {
			rules.MyRegistry.ProcessProject(git, project, rulesFn)
		}
	}
}

func main() {
	git, err := gitlab.NewClient(
		viper.GetString("gitlab.token"),
		gitlab.WithBaseURL(viper.GetString("gitlab.apiUrl")),
	)
	if err != nil {
		log.Fatalf("Failed to create new client: %v", err)
	}

	opt := &gitlab.ListProjectsOptions{
		Archived: gitlab.Bool(false),
		ListOptions: gitlab.ListOptions{
			PerPage: viper.GetInt("gitlab.perPage"),
			Page:    1,
		},
		OrderBy:          gitlab.String("path"),
		Search:           gitlab.String(""),
		SearchNamespaces: gitlab.Bool(true),
		Sort:             gitlab.String("asc"),
		Statistics:       gitlab.Bool(true),
	}

	var wg sync.WaitGroup

	for {
		projects, resp, err := git.Projects.ListProjects(opt)
		if err != nil {
			log.Fatal(err)
		}
		log.Debugf(
			"[Collector] Page %d of %d", resp.CurrentPage, resp.TotalPages,
		)

		wg.Add(1)
		go worker(projects, git, &wg)

		if resp.CurrentPage >= resp.TotalPages {
			break
		}
		opt.Page = resp.NextPage
	}
	wg.Wait()

	if err := processRules(rules.MyRegistry.Rules); err != nil {
		log.Errorf("[Collector] Error on insert rules data: %v", err)
	}

	if err := processProjects(rules.MyRegistry.Projects); err != nil {
		log.Errorf("[Collector] Error on insert projects data: %v", err)
	}

	if err := insertStats(rules.MyRegistry); err != nil {
		log.Errorf("[Collector] Error on insert statistics data: %v", err)
	}
	if err := processIssues(rules.MyRegistry, git); err != nil {
		log.Errorf("[Collector] Error on processing issue: %v", err)
	}
}
