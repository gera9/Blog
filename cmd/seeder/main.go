package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	// Default connection string, change as needed or set DATABASE_URL environment variable
	postgresConnStr := os.Getenv("POSTGRES_URL")

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, postgresConnStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close(ctx)

	// Path to the SQL file
	sqlPath := "internal/repositories/testdata/init-db.sql"
	sqlBytes, err := os.ReadFile(sqlPath)
	if err != nil {
		log.Fatalf("Unable to read SQL file: %v", err)
	}

	_, err = conn.Exec(ctx, string(sqlBytes))
	if err != nil {
		log.Fatalf("Unable to execute SQL: %v", err)
	}

	log.Println("Database seeded successfully")
}
