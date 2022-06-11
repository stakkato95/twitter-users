package main

import (
	_ "github.com/stakkato95/service-engineering-go-lib/logger"
	"github.com/stakkato95/twitter-service-users/app"
	"github.com/stakkato95/twitter-service-users/domain"
	"github.com/stakkato95/twitter-service-users/protoapp"
	"github.com/stakkato95/twitter-service-users/service"
)

func main() {
	repo := domain.NewUserRepo()
	service := service.NewUserService(repo)

	go protoapp.Start(service)
	app.Start(service)
}
