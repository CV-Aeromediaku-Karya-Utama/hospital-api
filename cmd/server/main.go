package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"os"

	"hospital-api/pkg/api"
	"hospital-api/pkg/app"
	"hospital-api/pkg/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := godotenv.Load(); err != nil {
		//log.Fatal("Error loading .env file")
		fmt.Println("Error loading .env file")
	}
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "This is the startup error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("SSL_MODE"))

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

	// create services
	roleService := api.NewRoleService(storage)
	userService := api.NewUserService(storage)
	authService := api.NewAuthService(storage)
	permissionService := api.NewPermissionService(storage)

	// start the server
	server := app.NewServer(
		router,
		authService,
		roleService,
		userService,
		permissionService,
	)
	err = server.Run()
	if err != nil {
		return err
	}
	return nil
}

func setupDatabase(connString string) (*sql.DB, error) {
	db, err := sql.Open(os.Getenv("DB_DRIVER"), connString)
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
