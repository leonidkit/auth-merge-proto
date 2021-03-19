package main

import (
	"auth-merge-api/internal/config"
	"auth-merge-api/internal/rest"
	"auth-merge-api/internal/service"
)

type AuthMergeApp struct {
	config  *config.Config
	service *service.Interface
	api     *rest.API
}

func main() {
	// TODO: AuthMergeApp init
}
