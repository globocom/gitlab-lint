// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules

type Level struct {
	Description string `json:"description" bson:"description"`
	ID          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Total       int32  `json:"total" bson:"total"`
}

const LevelError = "error"
const LevelExperimental = "experimental"
const LevelInfo = "info"
const LevelPedantic = "pedantic"
const LevelWarning = "warning"

var AllSeverities = map[string]Level{
	LevelError: {
		Description: "",
		ID:          LevelError,
		Name:        "Error",
	},
	LevelExperimental: {
		Description: "",
		ID:          LevelExperimental,
		Name:        "Experimental",
	},
	LevelInfo: {
		Description: "",
		ID:          LevelInfo,
		Name:        "Info",
	},
	LevelPedantic: {
		Description: "",
		ID:          LevelPedantic,
		Name:        "Pedantic",
	},
	LevelWarning: {
		Description: "",
		ID:          LevelWarning,
		Name:        "Warning",
	},
}
