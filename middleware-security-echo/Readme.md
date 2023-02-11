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

## JWT

JSON Web Tokens are an open, industry standard [RFC 7519](https://tools.ietf.org/html/rfc7519) method for representing claims securely between two parties. It is very easy to verify JWT tokens in go.

In this example we will be validating a JWT token that we generate in [jwt.io](jwt.io) website. With a payload (feel free to use any name or even any other payload here)

```json
{
  "name": "John Doe"
}
```

And for secret we use a string `very-secret` (goes without saying this is a secret so generate a longer more random string for your application). this will generate us a jwt token. If you dont want to go throught the trouble to generate this yourself, you can use this.

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UifQ.wSzHi09b5o8aSjDHjlGxED9Cg-_-8T6lTWZjs6_Netg
```

We write a new function called `JWTAuth` which is a middleware. In this we check for the Header with key `Authorization`. There is no rule that says token should be sent in this manner. But this in convention and many apps will expect to get the token in this header. So its best practice to keep it there.

We get the claim and attach it to the request context as extra data so we can get it in our handler when needed.

In our `handlerFunc` we get the value from context and respond back with the users name.

```bash
curl -H "Authorization: bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UifQ.wSzHi09b5o8aSjDHjlGxED9Cg-_-8T6lTWZjs6_Netg" http://localhost:7999/auth/test
```

Our server should respond back with
```json
{
  "data": "John Doe"
}
```

The `jwtCustomClaims` struct embeds `jwt.RegisteredClaims` struct. Which makes `jwtCustomClaims` an implementer of `Claims` interface. Struct embedding is interesting read more about it in [gobyexample](https://gobyexample.com/struct-embedding). 

```go
type jwtCustomClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}
```

This custom claim is being set because we setup echojwt with a config

```go
config := echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return new(jwtCustomClaims)
	},
	SigningKey: []byte("very-secret"),
}
```

In our `HandlerFunc` we grab the token value set onto our `echo.Context` with key `user`. This is the default value. We could have changed it with the config above if we wanted to. If you want to see how this got set you can take a look at the [implementation here](https://github.com/labstack/echo-jwt/blob/95b0b607987a3bed484870b2052ef212763742a4/jwt.go#L233)

```go
func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	return c.JSON(http.StatusOK, Message{Data: name})
}
```
