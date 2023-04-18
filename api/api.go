package api

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/tcarreira/api-server/pkg/types"
)

type Database struct {
	User         string
	Password     string
	Pets         []types.Pet
	People       []types.Person
	lastPetID    int
	lastPeopleID int
}

var MyDB *Database

// New is a helper function to create a new API server with a new database
func New(e *echo.Echo) *echo.Echo {
	MyDB = NewDB()
	return NewWithDB(e, MyDB)
}

func NewDB() *Database {
	return &Database{
		User:         os.Getenv("API_USER"),
		Password:     os.Getenv("API_PASSWORD"),
		Pets:         []types.Pet{},
		People:       []types.Person{},
		lastPetID:    0,
		lastPeopleID: 0,
	}
}

// NewWithDB is a helper function to create a new API server with a given database (useful for tests)
func NewWithDB(e *echo.Echo, db *Database) *echo.Echo {
	MyDB = db
	e.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, MyDB)
	})

	// Pets
	e.POST("/pets", petsCreate)
	e.GET("/pets", petsList)
	e.GET("/pets/:id", petsRead)
	e.PUT("/pets/:id", petsUpdate)
	e.DELETE("/pets/:id", petsDelete)

	// People
	e.GET("/people", peopleList)
	e.POST("/people", peopleCreate)
	e.GET("/people/:id", peopleRead)
	e.PUT("/people/:id", peopleUpdate)
	e.DELETE("/people/:id", peopleDelete)

	return e
}
