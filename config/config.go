// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package config

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	envReplacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(envReplacer)

	viper.SetDefault("debug", false)
	logLevel := log.InfoLevel
	if viper.GetBool("debug") {
		logLevel = log.DebugLevel
	}
	log.SetLevel(logLevel)

	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	viper.SetDefault("port", 8888)

	viper.SetDefault("gitlab.token", "")
	viper.SetDefault("gitlab.apiUrl", "https://gitlab.com/api/v4")
	viper.SetDefault("gitlab.perPage", 20)

	viper.SetDefault("mongodb.endpoint", "mongodb://localhost:27017")
	viper.SetDefault("mongodb.name", "gitlab-lint")

	viper.SetDefault("db.operation.timeout", 3)
	viper.SetDefault("db.perPage", 15)
	viper.SetDefault("db.perInsert", 50)
	viper.SetDefault("db.maxPerPage", 1000)

	viper.SetDefault("sentry.dsn", "")
	viper.SetDefault("sentry.timeout", 3)
}
