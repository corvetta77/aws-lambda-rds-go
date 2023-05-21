package client

import (
	"encoding/json"
	"io"
	"log"

	"github.com/corvetta77/lambda-to-rds/model"
)

func (client *bleboxClient) Get() ([]model.Sensor, error) {
	response, err := client.http.Get(client.url + "/state")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("fetch(): could not read response of GET request: %s\n", err)
		return nil, err
	}

	var responseData model.ResponseData
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		log.Printf("fetch(): could not Unmarshall response to type model.ResponseAccountData: %s\n", err)
		return nil, err
	}
	return responseData.MultiSensor.Sensors, nil

}
