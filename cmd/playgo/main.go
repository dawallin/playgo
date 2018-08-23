package main

import (
	"flag"
	"github.com/dawallin/playgo/pkg/domain"
	"github.com/dawallin/playgo/pkg/repository"
	"github.com/dawallin/playgo/pkg/webapi"
)

func main() {

	mongoDbConnectionString := flag.String("db", "127.0.0.1:27017", "mongo db connection string")

	flag.Parse()

	repo := repository.NewRepository(*mongoDbConnectionString)
	domain := domain.NewDomain(repo)

	api := webapi.NewWebapi(domain)

	api.Run(":8080")
}
