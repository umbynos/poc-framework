# openapi-generator POC

This POC leverages the `net/http` library but it uses a different approach: code-generation

## Convert from Swagger 2.0 to OpenAPI 3.0
https://stackoverflow.com/questions/59749513/how-to-convert-openapi-2-0-to-openapi-3-0
Paste your OpenAPI 2.0 definition into https://editor.swagger.io and select Edit > Convert to OpenAPI 3 from the menu.

I've taken the swagger from ../net-http-poc/docs/swagger.yaml and converted it into what you can find inside ./api.yaml file

## Install openapi-generator
`brew install openapi-generator`

## Doc
https://openapi-generator.tech/docs/generators/go-server

## Generate the server boilerplate
`openapi-generator generate -i api.yaml -g go-server`

## Start the POC
`go run openapi-generator-poc`

## Navigate to the doc
<!-- TODO -->