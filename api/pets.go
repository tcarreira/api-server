package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tcarreira/api-server/pkg/types"
)

func petsCreate(c echo.Context) error {
	pet := new(types.Pet)
	if err := c.Bind(pet); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	if pet.Name == "" || pet.Type == "" {
		return c.JSON(http.StatusBadRequest, "Name and Type are required")
	}
	pet.ID = MyDB.lastPetID
	MyDB.lastPetID++

	MyDB.Pets = append(MyDB.Pets, *pet)
	return c.JSON(http.StatusOK, *pet)
}

func petsList(c echo.Context) error {
	return c.JSON(http.StatusOK, MyDB.Pets)
}

func petsRead(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "ID must be an integer")
	}
	for _, p := range MyDB.Pets {
		if p.ID == id {
			return c.JSON(http.StatusOK, p)
		}
	}
	return c.JSON(http.StatusNotFound, "Pet not found")
}

func petsUpdate(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "ID must be an integer")
	}
	pet := new(types.Pet)
	if err := c.Bind(pet); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	for i, p := range MyDB.Pets {
		if p.ID == id {
			MyDB.Pets[i].Name = pet.Name
			MyDB.Pets[i].Type = pet.Type
			MyDB.Pets[i].Age = pet.Age
			MyDB.Pets[i].Description = pet.Description
			return c.JSON(http.StatusOK, MyDB.Pets[i])
		}
	}
	return c.JSON(http.StatusNotFound, "Pet not found")
}

func petsDelete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "ID must be an integer")
	}
	for i, p := range MyDB.Pets {
		if p.ID == id {
			MyDB.Pets = append(MyDB.Pets[:i], MyDB.Pets[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, "Pet not found")
}
