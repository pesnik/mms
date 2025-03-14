package main

import (
	"database/sql"
	"flag"
	"log"
	"os"

	"github.com/pressly/goose/v3"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	command := flag.String("command", "up", "Migration command: up, down, status")
	dir := flag.String("dir", "migrations", "Directory with migration files")
	flag.Parse()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	switch *command {
	case "up":
		if err := goose.Up(db, *dir); err != nil {
			log.Fatalf("Failed to apply migrations: %v", err)
		}
		log.Println("Migrations applied successfully")
	case "down":
		if err := goose.Down(db, *dir); err != nil {
			log.Fatalf("Failed to rollback migrations: %v", err)
		}
		log.Println("Last migration rolled back successfully")
	case "status":
		if err := goose.Status(db, *dir); err != nil {
			log.Fatalf("Failed to check migration status: %v", err)
		}
	default:
		log.Fatalf("Unknown command: %s. Use 'up', 'down', or 'status'", *command)
	}
}
