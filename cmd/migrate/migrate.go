package main

import (
	"errors"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	// if databaseURL == "" {
	// databaseURL = "mysql://app:password@tcp(db:3306)/halyk_task?parseTime=true"
	// }

	// log.Printf("Database URL: %s", databaseURL)

	m, err := migrate.New(
		"file://cmd/migrate/migrations", // Путь относительно рабочей директории
		databaseURL,
	)
	if err != nil {
		log.Fatal("Failed to create migrate instance:", err)
	}

	defer m.Close()

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migrations successful!")
}
