package repository

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"inventory-api/pkg/api"
	"log"
	"path/filepath"
	"runtime"
)

type Storage interface {
	RunMigrations(connectionString string, db *sql.DB) error
	api.UserRepository
	api.RoleRepository
	api.AuthRepository
	api.ProductCategoryRepository
	api.ProductRepository
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{db: db}
}

func (s *storage) RunMigrations(connectionString string, db *sql.DB) error {
	if connectionString == "" {
		return errors.New("repository: the connString was empty")
	}
	// get base path
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "../..")

	path := fmt.Sprint(basePath, "/pkg/repository/migrations/")
	migrationsPath := fmt.Sprintf("file:%s", path)
	driver, _ := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgresql", driver)

	if err != nil {
		log.Fatal(err)
	}
	// Migrate all the way up ...
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	return nil
}
