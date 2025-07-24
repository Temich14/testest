package main

import (
	"awesomeProject2/internal/Infrastructure/data/pg"
	"awesomeProject2/internal/adapters"
	"awesomeProject2/internal/adapters/handlers"
	"awesomeProject2/internal/app/server"
	"awesomeProject2/internal/core/usecases/create"
	"awesomeProject2/internal/core/usecases/get"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"os"
)

func main() {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	repo := pg.NewRepo(db)

	createUc := create.NewCreateUserHandler(repo)

	getUc := get.NewGetByIDHandler(repo)

	handler := handlers.NewGetHandler(getUc)

	httpServer := server.NewServer(*handler)

	go func() {
		err := httpServer.Run()
		if err != nil {

		}
	}()

	grcpSrv := adapters.NewMyGrpcServer(&createUc, getUc)

	go grcpSrv.Run()

	select {}

}
