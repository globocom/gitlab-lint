package rules

import "go.mongodb.org/mongo-driver/bson/primitive"

type Issue struct {
	ID          primitive.ObjectID `json:"_id"bson:"_id,omitempty"`
	ProjectID   int                `json:"projectId" bson:"projectId"`
	RuleID      string             `json:"ruleId" bson:"ruleId"`
	IssueID     int                `json:"issueId" bson:"issueId"`
	WebURL      string             `json:"webUrl" bson:"webUrl"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
}

type Issues []Issue

func (i Issue) Cast() Queryable {
	return &i
}

func (i Issue) GetCollectionName() string {
	return "issues"
}

func (i Issue) GetSearchableFields() []string {
	return []string{"title", "projectid", "ruleid", "state"}
}
