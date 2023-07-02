package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/corvetta77/lambda-to-rds/client"
	"github.com/corvetta77/lambda-to-rds/db"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	apiClient := client.NewBleboxClient("http://192.168.55.109", http.DefaultClient)
	measurements, api_err := apiClient.Get()
	if api_err != nil {
		return "lambda - problem with sensor API call", api_err
	}

	dbClient := db.NewDbClient()
	db_err := dbClient.Save(measurements)
	if db_err != nil {
		return "lambda - problem with db call", db_err
	}

	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
