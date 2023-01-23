# go-backend-template
A template for a GoLang backend API using Echo and MongoDB.

Based on @bmdavis419 [the-better-backend](https://youtu.be/6C-2R92L01Q).

## Getting Started

### Prerequisites

- [GoLang](https://golang.org/doc/install)
- [MongoDB](https://docs.mongodb.com/manual/installation/)

### Installing

0. Install extra packages: 
    ```go install github.com/cosmtrek/air@latest```
    ```go install github.com/swaggo/swag/cmd/swag@latest```
1. Clone the repo
2. Create your own .env file
3. ```make dev```
4. view docs at http://localhost:8080/swagger

### Scripts

- ```make dev``` - runs the server in development mode
- ```make swagger``` - generates the swagger docs
