package client

import (
	"net/http/httptest"
	"testing"

	"github.com/tcarreira/api-server/api"
	"github.com/tcarreira/api-server/pkg/types"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPetCreate(t *testing.T) {
	cli, db, deferFunc := setupServerWithDB(t)
	defer deferFunc()

	assert.Equal(t, 0, len(db.Pets))

	expected1 := types.Pet{
		Name:        "chien",
		Type:        "mouse",
		Description: "A nice dog",
	}
	expected2 := types.Pet{
		Name:        "aufauf",
		Type:        "snake",
		Age:         2,
		Description: "A nice cat",
	}

	err := cli.Pet().Create(&expected1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(db.Pets))
	assert.Equal(t, expected1, db.Pets[0])

	assert.Equal(t, 0, expected2.ID)
	err = cli.Pet().Create(&expected2)
	assert.NoError(t, err)
	assert.Equal(t, 1, expected2.ID)
	assert.Equal(t, 2, len(db.Pets))
	assert.Equal(t, expected2, db.Pets[1])
}

func TestPetList(t *testing.T) {
	cli, db, deferFunc := setupServerWithDB(t)
	defer deferFunc()

	l, err := cli.Pet().List()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(l))

	db.Pets = append(db.Pets, types.Pet{ID: 1}, types.Pet{ID: 2})
	l, err = cli.Pet().List()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(l))
}

func TestPetGet(t *testing.T) {
	db := api.NewDB()
	ts := httptest.NewServer(api.NewWithDB(echo.New(), db))
	defer ts.Close()

	cli, err := NewAPIClient(Config{
		Endpoint: ts.URL,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = cli.Pet().Get(1)
	assert.ErrorIs(t, err, ErrorNotFound)

	expected := types.Pet{
		ID:          1,
		Name:        "John",
		Age:         30,
		Description: "A nice guy",
	}
	db.Pets = append(db.Pets, expected)
	p, err := cli.Pet().Get(1)
	assert.NoError(t, err)
	assert.Equal(t, expected, *p)
}

func TestPetUpdate(t *testing.T) {
	cli, db, deferFunc := setupServerWithDB(t)
	defer deferFunc()

	db.Pets = append(db.Pets, types.Pet{ID: 1}, types.Pet{ID: 2})

	expected := types.Pet{
		ID:          1,
		Name:        "Buddy",
		Type:        "dog",
		Age:         30,
		Description: "A nice bud",
	}
	err := cli.Pet().Update(expected.ID, &expected)
	assert.NoError(t, err)
	assert.Equal(t, expected, db.Pets[0])

	err = cli.Pet().Update(88, &expected)
	assert.ErrorIs(t, err, ErrorNotFound)
}

func TestPetDelete(t *testing.T) {
	cli, db, deferFunc := setupServerWithDB(t)
	defer deferFunc()

	db.Pets = append(db.Pets, types.Pet{ID: 1}, types.Pet{ID: 2})

	err := cli.Pet().Delete(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(db.Pets))
	assert.Equal(t, 2, db.Pets[0].ID)

	err = cli.Pet().Delete(1)
	assert.ErrorIs(t, err, ErrorNotFound)
	assert.Equal(t, 1, len(db.Pets))

	err = cli.Pet().Delete(2)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(db.Pets))
}
