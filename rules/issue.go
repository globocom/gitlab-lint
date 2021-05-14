// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules

import "time"

type Issue struct {
	Closed        bool      `json:"closed" bson:"closed"`
	CreatedAt     time.Time `json:"createdAt" bson:"createdAt"`
	GitlabIssueID int       `json:"gitlabIssueID" bson:"gitlabIssueID"`
	ModifiedAt    time.Time `json:"modifiedAt" bson:"modifiedAt"`
	ProjectID     int       `json:"projectId" bson:"projectId"`
	RuleID        string    `json:"ruleId" bson:"ruleId"`
}

func (i Issue) Cast() Queryable {
	return &i
}

func (i Issue) GetCollectionName() string {
	return "issues"
}

func (i Issue) GetSearchableFields() []string {
	return nil
}

func NewIssue(projectID int, ruleID string, gitlabIssueID int) Issue {
	return Issue{
		Closed:        false,
		CreatedAt:     time.Now().UTC(),
		GitlabIssueID: gitlabIssueID,
		ModifiedAt:    time.Now().UTC(),
		ProjectID:     projectID,
		RuleID:        ruleID,
	}
}
