package migration

import (
	"database/sql"
	"embed"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed scripts
var schema embed.FS

type Migration struct {
	migrate *migrate.Migrate
}

type logger struct{}

func (l *logger) Verbose() bool {
	return true
}

func (l *logger) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func New(db *sql.DB) (*Migration, error) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to prep DB for migration: %w", err)
	}

	source, err := iofs.New(schema, "scripts")
	if err != nil {
		return nil, fmt.Errorf("failed to read schema files: %w", err)
	}
	m, mErr := migrate.NewWithInstance("iofs", source, "postgres", driver)
	if mErr != nil {
		return nil, fmt.Errorf("failed to create db migration instance: %w", mErr)
	}
	m.Log = &logger{}
	return &Migration{
		migrate: m,
	}, err
}

func (m *Migration) Run() error {
	uErr := m.migrate.Up()
	if uErr != nil {
		return fmt.Errorf("failed to migrate database: %w", uErr)
	}
	return nil
}

func (m *Migration) Stop() {
	m.migrate.GracefulStop <- true
}
