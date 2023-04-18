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
