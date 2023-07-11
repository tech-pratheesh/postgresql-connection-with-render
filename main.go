package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	host := os.Getenv("PSQL_DB_HOST")
	port := os.Getenv("PSQL_DB_PORT")
	user := os.Getenv("PSQL_DB_USER")
	pass := os.Getenv("PSQL_DB_PASS")
	database := os.Getenv("PSQL_DB_DATABASE")
	sslmode := os.Getenv("PSQL_DB_SCHEMA")

	// Set up a connection string with the database credentials
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%v sslmode=%s", user, pass, database, host, port, sslmode)

	// Open a connection to the database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to PostgreSQL database!")

	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "success")
	})
	e.Start(":8080")
}
