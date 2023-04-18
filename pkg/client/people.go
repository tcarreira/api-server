package client

import (
	"encoding/json"
	"strconv"

	"github.com/tcarreira/api-server/pkg/types"
)

type peopleClient struct {
	client *APIClient
}

func (c *APIClient) People() *peopleClient {
	return &peopleClient{c}
}

func (c *peopleClient) Create(p *types.Person) error {
	data, err := c.client.DoPOST("/people", p)
	if err != nil {
		return err
	}
	var person types.Person
	err = json.Unmarshal(data, &person)
	if err != nil {
		return err
	}
	p.ID = person.ID // update the ID with server info
	return nil
}

func (c *peopleClient) Get(id int) (*types.Person, error) {
	data, err := c.client.DoGET("/people/" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	var person types.Person
	err = json.Unmarshal(data, &person)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func (c *peopleClient) List() ([]*types.Person, error) {
	data, err := c.client.DoGET("/people")
	if err != nil {
		return nil, err
	}
	var people []*types.Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		return nil, err
	}
	return people, nil
}

func (c *peopleClient) Update(id int, p *types.Person) error {
	data, err := c.client.DoPUT("/people/"+strconv.Itoa(id), p)
	if err != nil {
		return err
	}
	var person types.Person
	err = json.Unmarshal(data, &person)
	if err != nil {
		return err
	}
	p.ID = person.ID // update the ID with server info
	return nil
}

func (c *peopleClient) Delete(id int) error {
	return c.client.DoDELETE("/people/" + strconv.Itoa(id))
}
