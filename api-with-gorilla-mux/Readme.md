# Gorilla Mux Router

In the previous section we saw how we can create a simele REST endpoint with net/http. We saw there were some limitations. But In most cases when we don't need complex path matching, net/http works just fine. 

Lets see how gorilla mux addresses the issues we saw with net/http.

## Run the example

```bash
git checkout origin/gorilla-mux-01
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

## Why gorilla mux

In benchmarks gorilla mux is probably one of the slower routers out there. 

The following if froma benchmark result with one parameter.

| lib/framework | Operations K | ns/op | B/op | allocs/op |
|:-------------:|:------------:|:-----:|:----:|:---------:|
|     Beego     |      442     |  2791 |  352 |     3     |
|      Chi      |     1000     |  1006 |  432 |     3     |
|      Echo     |     14662    |  81.9 |   0  |     0     |
|      Gin      |     16683    |  72.3 |   0  |     0     |
|  Gorilla Mux  |      434     |  2943 | 1280 |     10    |
|   HttpRouter  |     23988    |   50  |   0  |     0     |

HttpRouter is around 60x faster. So why are we starting with gorilla mux?

For one, your router will almost never be your bottleneck, in an application where you have file i/o or database operations your performance / speed will depend way more on those things than your router. Also gorilla mux is compliant with net/http. So if you have a project already written with net/http you will have a easier time converting to mux compared to some of the other libraries / frameworks that has a different implementation of Handler.

## Routing based on Http Method

With net/http we were sending all Http Verb request to the route and handling each type in the same function. With mux we can specify the http method for each our route.

```go
  r.HandleFunc("/user", s.getUser).Methods(http.MethodGet)
  r.HandleFunc("/user", s.updateUser).Methods(http.MethodPut)
```

We will also need to create corresoponding methods for each routes. and extract 

With these our api should behave exactly the same.

```bash
curl localhost:7999/user 
```

```bash
curl -X PUT -d '{"username":"mofi","email":"mofi@gmail.com","age":27}' localhost:7999/user
```

