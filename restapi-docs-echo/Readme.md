# API Docs With Swagger/Open API

Documentation is a core part of programming. Specially for REST APIs. Using documentation we let our users know how they can use our API. Go is a very readable (opinion) language. While anyone can read our code an have a good idea of how our API works, with a openapi spec we can do much better.

## What is Open API?

The OpenAPI Specification (OAS) provides a consistent means to carry information through each stage of the API lifecycle. It is a specification language for HTTP APIs that defines structure and syntax in a way that is not wedded to the programming language the API is created in. API specifications are typically written in YAML or JSON, allowing for easy sharing and consumption of the specification. 

## Getting Started

```bash
git checkout origin/rest-api-docs-01
```

## Open API Schema

Open API is written in JSON or YAML. For example our schema for the API endpoint for getting a user with an id looks like this:

```yaml
basePath: /
definitions:
  handler.HTTPError:
    properties:
      message: {}
    type: object
  handler.User:
    properties:
      company:
        example: Acme Inc.
        type: string
      email:
        example: johndoe@gmail.com
        type: string
      id:
        example: 1
        type: integer
      interests:
        example:
        - golang
        - python
        items:
          type: string
        type: array
      name:
        example: John Doe
        type: string
    type: object
host: localhost:7999
paths:
  /api/v1/users/{id}:
    get:
      consumes:
      - '*/*'
      description: get user matching given ID.
      parameters:
      - description: user id
        in: path
        name: id
        required: true
        type: integer
      produces:
      - application/json
      responses:
        "200":
          description: OK
          schema:
            $ref: '#/definitions/handler.User'
        "404":
          description: Not Found
          schema:
            $ref: '#/definitions/handler.HTTPError'
      summary: Get user by id
      tags:
      - API
schemes:
- http
swagger: "2.0"
```

## Generating API Docs

Writing these YAMLs by hand would be error prone. So we can make use of tools to generate it from comments for us.

The same schema can be written in Go as:

```go
// API godoc
// @Summary Get user by id
// @Description get user matching given ID.
// @Tags API
// @Accept */*
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} User
// @Failure 404 {object} HTTPError
// @Router /api/v1/users/{id} [get]
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
    // ...
}
```

## Swagger UI

To generate and serve swagger UI we can use [swag]() and [echo-swagger]().

```bash
swag init -g cmd/web/main.go --output docs/
```

Then we run the application as normal

```bash
go run cmd/web/main.go
```

And visit [http://localhost:7999/swagger](http://localhost:7999/swagger)