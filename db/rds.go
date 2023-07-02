package db

import (
	"database/sql"
	"fmt"

	"github.com/corvetta77/lambda-to-rds/model"
)

type dbClient struct {
	connectionString string
}

type DbClient interface {
	Save([]model.Sensor) error
}

func NewDbClient() DbClient {
	dbEndpoint := ""
	dbUser := ""
	dbPassword := ""
	db_name := ""

	return &dbClient{
		connectionString: fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbEndpoint, db_name),
	}
}

func (client *dbClient) Save(measurements []model.Sensor) error {

	db, err := sql.Open("mysql", client.connectionString)
	if err != nil {
		fmt.Println("error: %s", err)
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}

	defer db.Close()

	return nil
}
