package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

const (
	dbUser     = "latte"
	dbPassword = "latte"
	dbHost     = "db"
	dbPort     = "5432"
	dbName     = "frappuccino"
)

func getDBConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)
}

func main() {
	// Connect to the database
	connStr := getDBConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open a DB connection: %v", err)
	}
	defer db.Close()

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to the DB: %v", err)
	}

	fmt.Println("Connected to the database!")

	// Query the database
	query := "SELECT * from customers"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}
	defer rows.Close()

	// Iterate over the rows
	fmt.Println("Users:")
	for rows.Next() {
		var id int
		var name, email string

		// Scan the row into variables
		if err := rows.Scan(&id, &name, &email); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}

		// Print the result
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
	}

	// Check for errors encountered during iteration
	if err := rows.Err(); err != nil {
		log.Fatalf("Error during row iteration: %v", err)
	}
}

