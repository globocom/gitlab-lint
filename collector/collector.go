// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package main

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
	"go.mongodb.org/mongo-driver/bson"

	_ "github.com/globocom/gitlab-lint/config"
	"github.com/globocom/gitlab-lint/db"
	"github.com/globocom/gitlab-lint/rules"
)

func processProjects(projects map[string]rules.Project) error {
	dbInstance, err := db.NewMongoSession()
	if err != nil {
		log.Errorf("[Main] Error on create mongo session: %v", err)
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
		log.Errorf("[Main] Error on create mongo session: %v", err)
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

func insertStats(r *rules.Registry, projectsCount int) error {
	dbInstance, err := db.NewMongoSession()
	if err != nil {
		log.Errorf("[Main] Error on create mongo session: %v", err)
		return err
	}

	stats := &rules.Stats{
		CreatedAt:            time.Now().UTC(),
		RegisteredRulesCount: len(r.RulesFn),
		ProjectsCount:        len(r.Projects),
		RulesCount:           len(r.Rules),
		GitlabProjectsCount:  projectsCount,
	}

	log.Debug("[Collector] Inserting collector statistics")
	if _, err := dbInstance.Insert(stats); err != nil {
		return err
	}

	return nil
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

	var gitlabProjectsCount int

	for {
		projects, resp, err := git.Projects.ListProjects(opt)
		if err != nil {
			log.Fatal(err)
		}

		gitlabProjectsCount = resp.TotalItems

		log.Debugf(
			"[Collector] Page %d of %d", resp.CurrentPage, resp.TotalPages,
		)

		for _, project := range projects {
			// Only non-forks
			if project.ForkedFromProject != nil {
				continue
			}

			for _, rulesFn := range rules.MyRegistry.RulesFn {
				rules.MyRegistry.ProcessProject(git, project, rulesFn)
			}
		}

		if resp.CurrentPage >= resp.TotalPages {
			break
		}

		opt.Page = resp.NextPage
	}

	if err := processRules(rules.MyRegistry.Rules); err != nil {
		log.Errorf("[Main] Error on insert rules data: %v", err)
	}

	if err := processProjects(rules.MyRegistry.Projects); err != nil {
		log.Errorf("[Main] Error on insert projects data: %v", err)
	}

	if err := insertStats(rules.MyRegistry, gitlabProjectsCount); err != nil {
		log.Errorf("[Main] Error on insert statistics data: %v", err)
	}
}
