# Echo Web Framwork

In the previous section we saw how we can create a simele REST endpoint with net/http. We saw there were some limitations. But In most cases when we don't need complex path matching, net/http works just fine.

Lets see how Echo solves those problem.

## Run the example

```bash
git checkout origin/echo-03
```

If you are not already in the folder

```bash
cd api-with-echo
```

```bash
go run main.go
```

```bash
curl localhost:7999/api/v1/users
```

## Why Echo

In benchmarks echo is one of the faster ones out there.

The following is from a benchmark result with one parameter.

| lib/framework | Operations K | ns/op | B/op | allocs/op |
|:-------------:|:------------:|:-----:|:----:|:---------:|
|     Beego     |      442     |  2791 |  352 |     3     |
|      Chi      |     1000     |  1006 |  432 |     3     |
|      Echo     |     14662    |  81.9 |   0  |     0     |
|      Gin      |     16683    |  72.3 |   0  |     0     |
|  Gorilla Mux  |      434     |  2943 | 1280 |     10    |
|   HttpRouter  |     23988    |   50  |   0  |     0     |

HttpRouter is around 2x faster. So why are we starting with Echo?

Echo is more full featured compared to some of the routers in the list. It has many built in features that makes building rest api a breeze.

## Routing based on Http Method

With net/http we were sending all Http Verb request to the route and handling each type in the same function. With echo we can specify the http method for each our route.

```go
	e.GET("/users", getAllUsers)
	e.GET("/instructors", getAllInstructors)
	e.GET("/courses", getAllCourses)
```

We can curl for response

```bash
curl localhost:7999/users
```

## Binding Parameters

Echo lets us bind parametes from different sources (path param, query param, request body). We can still reach into the request inside `c` if we want to. But bind is a nice helper that can do type assertions and catch validation errors automatically.

```go
	if err := echo.QueryParamsBinder(c).
		Strings("topics", &topics).
		Int("instructor", &instructor).
		Strings("attendee", &attendees).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "incorrect usage of query param")		
	}
```

We can do the same thing for path Parameters too.

```bash
curl localhost:7999/instructors/1
```

We could also use query parameters like

```bash
curl "localhost:7999/api/v1/courses?topic=go&topic=devops"
```

Query parameters don't have any generic syntax that talks about how query params should be formatted. For example you can find multiple different ways to pass these value

[RFC 3986: Uniform Resource Identifier (URI): Generic Syntax](https://www.rfc-editor.org/rfc/rfc3986) has not no fix guidance on how to do it. You will probably find all these different ways to pass multiple query values for the same key. Its left up to practioners on how to deal with these.

1. example.com/path?name=name1&name=name2
2. example.com/path?name=name1,name2
3. example.com/path?name=[name1,name2]

In our code we implement the first way. Its no more right than any other way. Also its upto us to determine whether we are choosing to take in an array or taking only one of the query param.

For example

```bash
curl "localhost:7999/api/v1/courses?instructor=1&instructor=2"
```

will only take the first query to filter by. 

But

```bash
curl "localhost:7999/api/v1/courses?topic=go&topic=devops"
```

Will take both query and only return courses that has both `go` and `devops` in their topic.

So we are implementing is as a logical `and` operator. 

```bash
curl "localhost:7999/api/v1/users?interest=go&interest=swift&interest=hadoop"
```

will return users who has interest in go, swift and haddop.

It is also reasonable to implement that as a logical `or` depending on your application. We might want to find any course where userId 1, 2 or 3 signed up. The code currently does not do this. You can take it as an excercise to implement this.

> Hint: You can look at the `Contains` function to see what you need to change to make it work as logical or.

As long as we are implementing the behaviour we expect from our application we are fine. 

## Group

So far we have been adding all our routes at the top level of our hostname. i.e. `localhost:7999/<route>`. But there are times when it is desired to have routes that are grouped together based on some criteria. For example we might want all our authentication routes grouped together under `/auth` or we could want to version our api with `/api/v1`. With a sub-router it is possible to apply rules and logic to a group of routes instead of applying these rules individually.

To create a group

```go
    api := e.Group("api/v1")
``` 

We can then treat `api` as if it were a instance of echo and add new routes to it. Any route added to this subrouter will be prefixed with `/api/v1` so the path `/users` become `/api/v1/users`.

We can still test that all our routes work as expected.

```bash
curl localhost:7999/api/v1/instructors
```
