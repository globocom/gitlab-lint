// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package main

import (
	"log"

	"github.com/globocom/gitlab-lint/api"
	_ "github.com/globocom/gitlab-lint/config"
	_ "github.com/globocom/gitlab-lint/docs"
)

// @title gitlab-lint API
// @version 0.1.0

// @license.name BSD-3-Clause License
// @license.url https://opensource.org/licenses/BSD-3-Clause

// @host localhost:3000
// @BasePath /api/v1
func main() {
	server, err := api.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	server.Start()
}
