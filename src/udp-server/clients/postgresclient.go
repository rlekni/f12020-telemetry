package clients

import (
	"context"
	"database/sql"
	"main/helpers"

	"github.com/sirupsen/logrus"
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

func (client PostgreClient) Disconnect(ctx context.Context) error {
	logrus.Warningln("Closing PostgreSQL connection!")
	return client.Database.Close()
}

func (client PostgreClient) Insert(ctx context.Context, packetType string, packet interface{}) error {
	var id int
	err := client.Database.QueryRow(testInsertSQL, "test").Scan(&id)
	helpers.LogIfError(err)
	return err
}
