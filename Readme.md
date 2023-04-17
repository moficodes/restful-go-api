# RESTful Go API

This repo holds the code content for the RESTful GO API live trainging session.

This content will continue to grow as I learn more and get your feedback.

If any section is not 100% clear open an issue. If you see anything that you can fix, create a PR.

## Structure

In order to gradually build up the conent for optimal learning, I have decided to make use of git branches. Each branch
name will have the format `<Topic>-01..n`.

Each topic will be in their own folder and will be a complete go project.

At each level of the workshop the branch should be working code. If it is not I will mention it.

## Sections

1. Standard Library net/http
    - [Getting Started](https://github.com/moficodes/restful-go-api/tree/standard-library-net-http-01/api-with-net-http#run-the-example)
    - [Custom Handler Type](https://github.com/moficodes/restful-go-api/tree/standard-library-net-http-02/api-with-net-http#why-a-struct)
    - [JSON Response](https://github.com/moficodes/restful-go-api/tree/standard-library-net-http-03/api-with-net-http#json)
    - [HTTP Verbs](https://github.com/moficodes/restful-go-api/tree/standard-library-net-http-04/api-with-net-http#http-verbs)
    - [Request Body](https://github.com/moficodes/restful-go-api/tree/standard-library-net-http-05/api-with-net-http#rest-routes)
    - [Handler vs HandlerFunc vs *HandlerMethod](https://github.com/moficodes/restful-go-api/tree/standard-library-net-http-06/api-with-net-http#handler-vs-handlerfunc-vs-handlermethod)
    - [Path Parameter](https://github.com/moficodes/restful-go-api/tree/standard-library-net-http-07/api-with-net-http#path-parameter)

2. a. Gorilla Mux
    - [Why Gorilla Mux](https://github.com/moficodes/restful-go-api/tree/gorilla-mux-01/api-with-gorilla-mux#why-gorilla-mux)
    - [Path Parameter](https://github.com/moficodes/restful-go-api/tree/gorilla-mux-02/api-with-gorilla-mux#path-params)
    - [Query Parameter](https://github.com/moficodes/restful-go-api/tree/gorilla-mux-03/api-with-gorilla-mux#query-parameters)
    - [Match Query](https://github.com/moficodes/restful-go-api/tree/gorilla-mux-04/api-with-gorilla-mux#match-query)
    - [Sub Router](https://github.com/moficodes/restful-go-api/tree/gorilla-mux-05/api-with-gorilla-mux#sub-router)

   b. Echo
    - [Why Echo](https://github.com/moficodes/restful-go-api/tree/echo-01/api-with-echo#why-echo)
    - [Binding Parameters](https://github.com/moficodes/restful-go-api/tree/echo-02/api-with-echo#binding-parameters)
    - [Sub Router / Groups](https://github.com/moficodes/restful-go-api/tree/echo-03/api-with-echo#group)

3. a. Middleware and Security with Gorilla Mux
    - [Middleware](https://github.com/moficodes/restful-go-api/tree/middleware-security-01/middleware-security#middleware)
    - [Chaining Middlewares](https://github.com/moficodes/restful-go-api/tree/middleware-security-02/middleware-security#chaining-middlewares)
    - [Mux Handlers](https://github.com/moficodes/restful-go-api/tree/middleware-security-03/middleware-security#mux-handlers)
    - [JWT Auth](https://github.com/moficodes/restful-go-api/tree/middleware-security-04/middleware-security#jwt-authentication)

   b. Middleware and Security with Echo
    - [Middleware](https://github.com/moficodes/restful-go-api/tree/middleware-echo-01/middleware-security-echo#middleware)
    - [Chaining & Echo Middlewares](https://github.com/moficodes/restful-go-api/tree/middleware-echo-02/middleware-security-echo#chaining-middleware)
    - [JWT Auth](https://github.com/moficodes/restful-go-api/tree/middleware-echo-03/middleware-security-echo#jwt)
    - [JWT Auth with EchoJWT](https://github.com/moficodes/restful-go-api/tree/middleware-echo-04/middleware-security-echo#jwt)
4. Project Structure
    - [Common Project Structures for Go application](https://github.com/moficodes/restful-go-api/tree/project-structure-01/project-structure)

5. Testing and Benchmarking
    - [Unit Testing](https://github.com/moficodes/restful-go-api/tree/testing-benchmarking-01/testing-benchmark)
    - [Unit Testing With Echo](https://github.com/moficodes/restful-go-api/tree/testing-benchmarking-echo-01/testing-benchmark-echo)

6. Database
    - [Postgres Database](https://github.com/moficodes/restful-go-api/tree/rest-api-database-01/rest-api-database#go--postgres)
    - [Testing Database](https://github.com/moficodes/restful-go-api/tree/rest-api-database-02/rest-api-database#testing)

7. Application Delivery
    - [Docker Compose](https://github.com/moficodes/restful-go-api/tree/containers-01/rest-api-container)
    - [Kubernetes](https://github.com/moficodes/restful-go-api/tree/containers-02/rest-api-container)

8. Docs Generation
    - [Swagger](https://github.com/moficodes/restful-go-api/tree/rest-api-docs-01/restapi-docs-echo)