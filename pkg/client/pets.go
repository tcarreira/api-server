package client

import (
	"encoding/json"
	"strconv"

	"github.com/tcarreira/api-server/pkg/types"
)

type petsClient struct {
	client *APIClient
}

func (c *APIClient) Pet() *petsClient {
	return &petsClient{c}
}

func (c *petsClient) Create(p types.Pet) (*types.Pet, error) {
	data, err := c.client.DoPOST("/pets", p)
	if err != nil {
		return nil, err
	}
	var pet types.Pet
	err = json.Unmarshal(data, &pet)
	if err != nil {
		return nil, err
	}
	return &pet, nil
}

func (c *petsClient) Get(id int) (*types.Pet, error) {
	data, err := c.client.DoGET("/pets/" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	var pet types.Pet
	err = json.Unmarshal(data, &pet)
	if err != nil {
		return nil, err
	}
	return &pet, nil
}

func (c *petsClient) List() ([]*types.Pet, error) {
	data, err := c.client.DoGET("/pets")
	if err != nil {
		return nil, err
	}
	var pets []*types.Pet
	err = json.Unmarshal(data, &pets)
	if err != nil {
		return nil, err
	}
	return pets, nil
}

func (c *petsClient) Update(id int, p types.Pet) (*types.Pet, error) {
	data, err := c.client.DoPUT("/pets/"+strconv.Itoa(id), p)
	if err != nil {
		return nil, err
	}
	var pet types.Pet
	err = json.Unmarshal(data, &pet)
	if err != nil {
		return nil, err
	}
	return &pet, nil
}

func (c *petsClient) Delete(id int) error {
	return c.client.DoDELETE("/pets/" + strconv.Itoa(id))
}
