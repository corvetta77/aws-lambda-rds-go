package db_test

import (
	"testing"

	"github.com/corvetta77/lambda-to-rds/db"
	"github.com/corvetta77/lambda-to-rds/model"
)

func TestGettingDataFromSensor(t *testing.T) {
	var dbClient = db.NewDbClient()
	err := dbClient.Save(testdata("sasa"))
	if err != nil {
		t.Errorf("db() should do ping wo error=%s", err)
	}
}

func testdata(accountId string) []model.Sensor {
	return []model.Sensor{{
		Id:           123,
		Type:         &accountId,
		Value:        11,
		State:        1,
		Trend:        2,
		ElapsedTimeS: 1,
	}}
}
