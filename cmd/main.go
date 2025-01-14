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
	query := "SELECT id, name, email, phone from customers"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}
	defer rows.Close()

	// Iterate over the rows

	fmt.Println("Query results:")
	for rows.Next() {
		var id string
		var name string
		var email string
		var phone string
		if err := rows.Scan(&id, &name, &email, &phone); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
			fmt.Println(err)

		}
		fmt.Println(id, name, email, phone)
	}

}
