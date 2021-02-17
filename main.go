// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package main

import (
	"log"

	"github.com/globocom/gitlab-lint/api"
	_ "github.com/globocom/gitlab-lint/config"
)

func main() {
	server, err := api.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	server.Start()
}
