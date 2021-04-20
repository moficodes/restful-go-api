# Middleware and Security

In this section we will be talking about middleware and security. The topic may seem a little unrelated but implementation wise they go hand in had.

## Try it out

To run the code in this section

```bash
git checkout origin/middleware-security-01
```

If you are not already in the folder

```bash
cd middleware-security
```

```bash
go run main.go
```

```bash
curl localhost:7999
```

## Middleware

Middleware is a function that wraps our handler. Thats all.

```go
func Middleware(h handler) handler
```

This simple implementation has alot of power. In go functions can be passed in to other functions as a parameter.

Say we want to add a log to every request we are serving that prints out the URL of the request.

We can do that by creating our own middleware like so.

```go
func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println(c.Request().URL)
		return next(c)
	}
}
```

and then using it at the root `echo` router

```go
e := echo.New()
e.Use(Logger)
```

On any request made to our server it will now print out the url of request.

## Chaining Middleware

`echo.Use` takes in a slice of middlewares we want to use and apply them in reverse order.

We can also do it manually ourselves

```go
func Chain(h echo.HandlerFunc, middleware ...func(echo.HandlerFunc) echo.HandlerFunc) echo.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}
```

This is less flexible compared to what echo provides out of the box with `echo.Use`. 

## Echo Middlewares

Echo has a list of middlewares built in from the middleware package. This includes CORS, CSRF, JWT, Jaeger, Prometheus and many more. The logger middlerware we used in the last section is also a middleware from echo. You can find the full list at [echo docs](https://echo.labstack.com/middleware/)