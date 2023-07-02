package client

import (
	"net/http"

	"github.com/corvetta77/lambda-to-rds/model"
)

type bleboxClient struct {
	http *http.Client
	url  string
}

type BleboxClient interface {
	Get() ([]model.Sensor, error)
}

func NewBleboxClient(url string, httpClient *http.Client) BleboxClient {
	return &bleboxClient{
		http: httpClient,
		url:  url,
	}
}
