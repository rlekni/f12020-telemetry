package clients

import (
	"context"
	"database/sql"
	"fmt"
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

func (client PostgreClient) Connect(ctx context.Context) error {
	logrus.Infoln("Connected to PostgreSQL!")

	return nil
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

func (client PostgreClient) Update(packet interface{}) error {
	logrus.Infoln("Updating Row in PostgreSQL!")

	return fmt.Errorf("NOT IMPLEMENTED FOR PostgreSQL!")
}

func (client PostgreClient) Delete(id string) error {
	info := fmt.Sprintf("Removing Row with id: %s from PostgreSQL!", id)
	logrus.Infoln(info)

	return fmt.Errorf("NOT implemented for PostgreSQL!")
}
