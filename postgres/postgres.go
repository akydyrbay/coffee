package postgres

import "fmt"

const (
	dbUser     = "latte"
	dbPassword = "latte"
	dbHost     = "db"
	dbPort     = "5432"
	dbName     = "frappuccino"
)

func GetDBConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)
}
