package main

import (
	"database/sql"
	"fmt"
	"os"

	"weight-tracker-api/pkg/api"
	"weight-tracker-api/pkg/app"
	"weight-tracker-api/pkg/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "This is the startup error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	connectionString := "root:@tcp(localhost:3306)/test"
	db, err := setupDatabase(connectionString)
	if err != nil {
		return err
	}

	// create storage dependency
	storage := repository.NewStorage(db)
	err = storage.RunMigrations(connectionString, db)
	if err != nil {
		return err
	}

	// create router dependency
	router := gin.Default()
	router.Use(cors.Default())

	// create services
	roleService := api.NewRoleService(storage)
	userService := api.NewUserService(storage)

	// start the server
	server := app.NewServer(router, roleService, userService)
	err = server.Run()
	if err != nil {
		return err
	}
	return nil
}

func setupDatabase(connString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	// ping the db to ensure that it is connected
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
