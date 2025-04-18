
# go-rest-api

Welcome to the go-rest-api! This is a basic REST API that allows you to manage items (or artists, depending on your implementation).<br> You can create, retrieve, and delete items using simple HTTP requests.


## Table of Contents :
### Overview

### Getting Started
### Features

### API Endpoints

### Contributing




## Overview
The go-rest-api provides a set of endpoints that allow users to manage items. The API uses standard HTTP methods such as GET, POST, and DELETE, and returns JSON responses with relevant data.
## Getting Started

To use thego-rest-api, you need to have an HTTP client (e.g., Postman, Curl) to send requests. The API is structured to provide an easy-to-use interface with clear responses for various actions.

### Requirements:
Go 1.22+

Internet connection (for API requests)
## Installation

Install my-project with npm

```bash
git clone https://github.com/helouazizi/go-rest-api.git
cd go-rest-api
go run cmd/main.go
```

` http://localhost:8080` 8080 by default fell free to edit the confics
    
## Features
- only go standard pkgs
- simple files structur
- clean and modular code
- Separation of Concerns
- using in memory storage
- using mutex eviting race conditions  
- error handling
- dependency injection to acheive the (IoC)
- logger and config pgk 
- paginations




## API Endpoints
`POST /api/items` to craete item 

`GET /api/items` to get items

`GET /api/items?id=1` to get specific item

`DELETE /api/items?id=1` to delete specific item
## Contributing


Contributions are welcome! If you'd like to contribute to this project, please fork the repository and submit a pull request. You can help by:

- Reporting bugs or issues.

- Adding new features.

- Improving documentation.


## Authors

- [@helouazizi](https://www.github.com/helouazizi)
