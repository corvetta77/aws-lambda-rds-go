package client_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/corvetta77/lambda-to-rds/client"
)

func TestGettingDataFromSensor(t *testing.T) {
	var apiClient = client.NewBleboxClient("http://192.168.55.109", http.DefaultClient)

	got, err := apiClient.Get()
	if err != nil {
		t.Errorf("get() should return 200 and correct response but got error: %v", err)
		return
	}

	if got[0].Value == 0 {
		t.Errorf("get() should return non-zero value for humidity, got: %v", got[0].Value)
	}
	if got[1].Value == 0 {
		t.Errorf("get() should return non-zero value for temperature, got: %v", got[1].Value)
	}
	fmt.Printf("Got results, humidity=%d, temperature=%d\n", got[0].Value, got[1].Value)
}
