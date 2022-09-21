package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"

	"hospital-api/pkg/api"
	"hospital-api/pkg/app"
	"hospital-api/pkg/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "This is the startup error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	connectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))

	db, err := setupDatabase(connectionString)
	if err != nil {
		return err
	}

	gormDM, err := setupGormDatabase(connectionString)
	if err != nil {
		return err
	}

	// create storage dependency
	storage := repository.NewStorage(db, gormDM)
	//if err = storage.RunMigrations(connectionString, db); err != nil {
	//	return err
	//}
	if err = storage.RunGormMigrations(gormDM); err != nil {
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
	db, err := sql.Open("postgres", connString)
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

func setupGormDatabase(connString string) (*gorm.DB, error) {
	gormDB, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return gormDB, nil
}
