package main

import (
	"log"

	apihandler "github.com/Stezok/bookhub/internal/deliver/http/handler/api"
	mysqlrepo "github.com/Stezok/bookhub/internal/repository/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/Stezok/bookhub/internal/repository"
	"github.com/Stezok/bookhub/internal/service"
	_ "github.com/go-sql-driver/mysql"
)

var database *sqlx.DB

func init() {
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	// defer cancel()

	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// database = client.Database("BookHub")

	var err error
	database, err = sqlx.Connect("mysql", "root:Sonya<317@(localhost:3306)/BookHub")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	repository := repository.Repository{
		Book: mysqlrepo.NewMySQLBookRepository(database),
	}

	service := service.Service{
		Book: service.NewBookService(repository),
	}

	handler := apihandler.NewAPIHandler(
		service,
		log.Default(),
	)

	r := handler.InitRoutes()
	r.Run()
}
