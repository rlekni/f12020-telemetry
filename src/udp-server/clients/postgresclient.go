package clients

import (
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
)

const (
	host          = "postgres"
	port          = 5432
	user          = "admin"
	password      = "password123"
	dbname        = "f1telemetry"
	testInsertSQL = `
	INSERT INTO packets (type)
	VALUES ($1)
	RETURNING id`
)

type PostgreClient struct {
	Database *sql.DB
}

func NewPostgreConnection() PostgreClient {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logrus.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		logrus.Fatalln(err)
	}
	// defer db.Close()

	fmt.Println("Connected to PostgreSQL!")

	return PostgreClient{
		Database: db,
	}
}

func (client PostgreClient) Insert(packet interface{}) {
	var id int
	err := client.Database.QueryRow(testInsertSQL, "test").Scan(&id)
	if err != nil {
		logrus.Fatalln(err)
	}

	fmt.Println("New record ID is:", id)
}
