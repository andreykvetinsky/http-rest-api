package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL_TEST")
	if databaseURL == "" {
		databaseURL = "host=localhost port=5433 user=postgres password=postgres dbname=restapi_dev sslmode=disable"
	}

	os.Exit(m.Run())
}
