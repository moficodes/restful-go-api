# Echo Web Framwork

In the previous section we saw how we can create a simele REST endpoint with net/http. We saw there were some limitations. But In most cases when we don't need complex path matching, net/http works just fine.

Lets see how Echo solves those problem.

## Run the example

```bash
git checkout origin/echo-01
```

If you are not already in the folder

```bash
cd api-with-echo
```

```bash
go run main.go
```

```bash
curl localhost:7999
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
	e.GET("/", s.home)
	e.GET("/user", s.getUser)
	e.PUT("/user", s.updateUser)
```

We will also need to create corresoponding methods for each routes.

With these our api should behave exactly the same.

```bash
curl localhost:7999/user
```

We can update the user

```bash
curl -X PUT -H 'Content-Type: application/json' -d '{"username":"mofi","email":"mofi@gmail.com","age":27}' localhost:7999/user
```

Then see the update was successful

```bash
curl localhost:7999/user
```