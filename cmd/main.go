package main

import (
	"context"
	"fmt"
	"frappuccinno/postgres"
	"os"

	"github.com/jackc/pgx/v5"
)

// const (
// 	dbUser     = "latte"
// 	dbPassword = "latte"
// 	dbHost     = "db"
// 	dbPort     = "5432"
// 	dbName     = "frappuccino"
// )

// func getDBConnectionString() string {
// 	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
// 		dbUser, dbPassword, dbHost, dbPort, dbName)
// }

// id VARCHAR(255) PRIMARY KEY,
// name VARCHAR(255) NOT NULL,
// email VARCHAR(255) UNIQUE,
// phone VARCHAR(50),
func main() {
	// Connect to the database
	connStr := postgres.getDBConnectionString()
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())
	err = conn.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	var id string
	var name string
	var email string
	var phone string

	err = conn.QueryRow(context.Background(), "select id, name, email, phone from customers").Scan(&id, &name, &email, &phone)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}
	fmt.Printf("id: %s, name: %s, email: %s, phone: %s\n", id, name, email, phone)
}
