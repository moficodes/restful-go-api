# Middleware and Security

In this section we will be talking about middleware and security. The topic may seem a little unrelated but implementation wise they go hand in had.

## Try it out

To run the code in this section

```bash
git checkout origin/gorilla-mux-05
```

If you are not already in the folder

```bash
cd api-with-gorilla-mux
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
func Logger(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Println(r.URL)
    next.ServeHTTP(w, r)
  })
}
```

and then wrapping our main router

```go
log.Fatal(http.ListenAndServe(":"+port, Logger(r)))
```

On any request made to our server it will now print out the url of request on every request.

## Chaining Middlewares

Middlewares are useful tools and often one is not enough. If we want to add more of these middlewares we can keep on wrapping our router with our middlewares.

For example

```go
...SomeOtherMiddlerware(OtherMiddlerWare(Logger(r)))
```

There is couple of other ways to do this.

```go
// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.Handler, middlewares ...func(next http.Handler) http.Handler) http.Handler {
  for _, m := range middlewares {
    f = m(f)
  }
  return f
}
```

Then we can use our middlewares like this,

```go
Chain(r, Logger, OtherMiddlerWare, SomeOtherMiddlerware, ...)
```

This in my opinion is a bit cleaner implementation.

In gorilla mux we have another option to do this.

Router has a method  `Use` that takes an array of Middlewares. This is useful to add middlewares to subrouters.