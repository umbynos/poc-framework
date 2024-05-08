# net/http POC

This POC leverages the `net/http` library. It uses https://github.com/swaggo/http-swagger to generate the documentation

## Install http-swagger
`go install github.com/swaggo/swag/cmd/swag@latest`

## Generate the swagger doc
`swag init --parseInternal true`

## Start the POC
`go run net-http-poc`

## Navigate to the doc
http://localhost:8080/v1/docs/index.html