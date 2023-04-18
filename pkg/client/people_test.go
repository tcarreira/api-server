package client

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/tcarreira/api-server/api"
	"github.com/tcarreira/api-server/pkg/types"
)

func TestPeopleCreate(t *testing.T) {
	cli, db, deferFunc := setupServerWithDB(t)
	defer deferFunc()

	assert.Equal(t, 0, len(db.People))

	expected1 := types.Person{
		Name:        "Alice",
		Age:         30,
		Description: "A nice person",
	}
	expected2 := types.Person{
		Name:        "Bob",
		Age:         25,
		Description: "A nice bob",
	}
	p1, err := cli.People().Create(expected1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(db.People))
	assert.Equal(t, expected1, *p1)
	assert.Equal(t, expected1, db.People[0])
	p2, err := cli.People().Create(expected2)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(db.People))
	expected2.ID = 1
	assert.Equal(t, expected2, *p2)
	assert.Equal(t, expected2, db.People[1])
}

func TestPeopleList(t *testing.T) {
	cli, db, deferFunc := setupServerWithDB(t)
	defer deferFunc()

	l, err := cli.People().List()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(l))

	db.People = append(db.People, types.Person{ID: 1}, types.Person{ID: 2})
	l, err = cli.People().List()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(l))
}

func TestPeopleGet(t *testing.T) {
	db := api.NewDB()
	ts := httptest.NewServer(api.NewWithDB(echo.New(), db))
	defer ts.Close()

	cli, err := NewAPIClient(Config{
		Endpoint: ts.URL,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = cli.People().Get(1)
	assert.ErrorIs(t, err, ErrorNotFound)

	expected := types.Person{
		ID:          1,
		Name:        "John",
		Age:         30,
		Description: "A nice guy",
	}
	db.People = append(db.People, expected)
	p, err := cli.People().Get(1)
	assert.NoError(t, err)
	assert.Equal(t, expected, *p)
}

func TestPeopleUpdate(t *testing.T) {
	cli, db, deferFunc := setupServerWithDB(t)
	defer deferFunc()

	db.People = append(db.People, types.Person{ID: 1}, types.Person{ID: 2})

	expected := types.Person{
		ID:          1,
		Name:        "John",
		Age:         30,
		Description: "A nice guy",
	}
	p, err := cli.People().Update(expected.ID, expected)
	assert.NoError(t, err)
	assert.Equal(t, expected, *p)
	assert.Equal(t, expected, db.People[0])

	_, err = cli.People().Update(88, expected)
	assert.ErrorIs(t, err, ErrorNotFound)
}

func TestPeopleDelete(t *testing.T) {
	cli, db, deferFunc := setupServerWithDB(t)
	defer deferFunc()

	db.People = append(db.People, types.Person{ID: 1}, types.Person{ID: 2})

	err := cli.People().Delete(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(db.People))
	assert.Equal(t, 2, db.People[0].ID)

	err = cli.People().Delete(1)
	assert.ErrorIs(t, err, ErrorNotFound)
	assert.Equal(t, 1, len(db.People))

	err = cli.People().Delete(2)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(db.People))
}
