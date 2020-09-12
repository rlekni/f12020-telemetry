package clients

import (
	"database/sql"
	"fmt"
	"main/helpers"
	"os"
	"strconv"
)

const (
	testInsertSQL = `
	INSERT INTO packets (type)
	VALUES ($1)
	RETURNING id`
)

type PostgreClient struct {
	Database *sql.DB
}

func NewPostgreConnection() PostgreClient {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DATABASE")
	port, err := strconv.ParseInt(os.Getenv("POSTGRES_PORT"), 10, 64)
	helpers.ThrowIfError(err)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	helpers.ThrowIfError(err)

	err = db.Ping()
	helpers.ThrowIfError(err)
	// defer db.Close()

	fmt.Println("Connected to PostgreSQL!")

	return PostgreClient{
		Database: db,
	}
}

func (client PostgreClient) Insert(packet interface{}) {
	var id int
	err := client.Database.QueryRow(testInsertSQL, "test").Scan(&id)
	helpers.ThrowIfError(err)

	fmt.Println("New record ID is:", id)
}
