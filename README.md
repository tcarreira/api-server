# api server

![Go Version](https://img.shields.io/github/go-mod/go-version/tcarreira/api-server)
[![Actions Status](https://github.com/tcarreira/api-server/workflows/Go/badge.svg)](https://github.com/tcarreira/api-server/actions)
[![Coverage](https://tcarreira.github.io/api-server/coverage/badge.svg)](https://tcarreira.github.io/api-server/coverage/coverage.html)
[![GitHub](https://img.shields.io/github/license/tcarreira/api-server)](https://github.com/tcarreira/api-server/blob/main/LICENSE)

A simple demo api server + api Client for CRUD demos (fully in-memory, no dependencies)

# Server

Download the latest version at https://github.com/tcarreira/api-server/releases/latest

or

Install it from source

```sh
go install github.com/tcarreira/api-server@latest
api-server --version
API_PORT=8888 api-server  # start server, listening on a non-default port
```

# Client

## Install

```
go get "github.com/tcarreira/api-server/pkg/client"
```

## Use

```go
package main

import (
	"fmt"

	"github.com/tcarreira/api-server/pkg/client"
	"github.com/tcarreira/api-server/pkg/types"
)

func main() {
	client, err := client.NewAPIClient(client.Config{Endpoint: "http://localhost:8888"})
	if err != nil {
		panic(err)
	}

	// Using the People API
	p := &types.Person{
		Name:        "John Doe",
		Age:         30,
		Description: "Nice person",
	}
	err = client.People().Create(p)
	if err != nil {
		panic(err)
	}

	_, err = client.People().Get(p.ID)
	_, err = client.People().List()
	err = client.People().Delete(p.ID)
	err = client.People().Update(p.ID, p)

	// Using the Pets API
	pets, err := client.Pet().List()
	if err != nil {
		panic(err)
	}
	fmt.Println(pets)
}
```

# Types

- `Pet`
- `Person`


# Endpoints

| endpoint | method | description |
| --- | --- | --- |
| `/` | GET | "ok" |
| `/status` | GET | server config info |
| `/version` | GET | info about version (built) |
| `/people` | POST | Create a person |
| `/people` | GET | List all people |
| `/people/:id` | GET | Get a person by ID |
| `/people/:id` | PUT | Update a person by ID |
| `/people/:id` | DELETE | Delete a person by ID |
| `/pets` | POST | Create a pet |
| `/pets` | GET | List all pets |
| `/pets/:id` | GET | Get a pet by ID |
| `/pets/:id` | PUT | Update a pet by ID |
| `/pets/:id` | DELETE | Delete a pet by ID |
