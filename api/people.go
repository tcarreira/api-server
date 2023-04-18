package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tcarreira/api-server/pkg/types"
)

func peopleCreate(c echo.Context) error {
	person := new(types.Person)
	if err := c.Bind(person); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	if person.Name == "" || person.Age <= 0 {
		return c.JSON(http.StatusBadRequest, "Name and Age>0 are required")
	}
	person.ID = MyDB.lastPeopleID
	MyDB.lastPeopleID++

	MyDB.People = append(MyDB.People, *person)
	return c.JSON(http.StatusOK, *person)
}

func peopleList(c echo.Context) error {
	return c.JSON(http.StatusOK, MyDB.People)
}

func peopleRead(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "ID must be an integer")
	}
	for _, p := range MyDB.People {
		if p.ID == id {
			return c.JSON(http.StatusOK, p)
		}
	}
	return c.JSON(http.StatusNotFound, "Person not found")
}

func peopleUpdate(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "ID must be an integer")
	}
	person := new(types.Person)
	if err := c.Bind(person); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	for i, p := range MyDB.People {
		if p.ID == id {
			MyDB.People[i].Name = person.Name
			MyDB.People[i].Age = person.Age
			MyDB.People[i].Description = person.Description
			return c.JSON(http.StatusOK, MyDB.People[i])
		}
	}
	return c.JSON(http.StatusNotFound, "Person not found")
}

func peopleDelete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "ID must be an integer")
	}
	for i, p := range MyDB.People {
		if p.ID == id {
			MyDB.People = append(MyDB.People[:i], MyDB.People[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, "Person not found")
}
