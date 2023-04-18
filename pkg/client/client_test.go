package client

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/tcarreira/api-server/api"
)

func setupServerWithDB(t *testing.T) (cli *APIClient, db *api.Database, deferFunc func()) {
	db = api.NewDB()
	ts := httptest.NewServer(api.NewWithDB(echo.New(), db))
	deferFunc = ts.Close

	var err error
	cli, err = NewAPIClient(Config{
		Endpoint: ts.URL,
	})
	if err != nil {
		t.Fatal(err)
	}

	return
}
