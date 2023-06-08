package repository

import (
	"fmt"
	"log"
	"sample/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() error {
	fmt.Println("Start running migration")
	// Directory where your migration files are located
	migrationsDir := "db/migrations"
	dsn := config.SysConfig.PostgresSource

	// Create a new instance of the migration driver
	driver, err := migrate.New("file://"+migrationsDir, dsn)
	if err != nil {
		return err
	}

	// Apply all available migrations
	err = driver.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	fmt.Println("Migrations completed successfully")
	return nil
}
