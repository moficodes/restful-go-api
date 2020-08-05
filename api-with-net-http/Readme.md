# REST API with Standard Library net/http

## Run the example
```
git checkout origin/standard-library-net-http-07
```

If you are not already in the folder
```
cd api-with-net-http
```

```
go run main.go
```

```
curl localhost:7999
```

You should see output `hello world` printed.

## Why a Struct
In our previous example we implemented handler interface on a `anything` type. That worked. So why we are using a `server` struct?
The benefit of using a struct becomes apparent when we build more complex servers. And want to have other methods / types in the struct for instance logger.

## Handlerfunc

The [HandlerFunc](https://golang.org/pkg/net/http/#HandlerFunc) type is an adapter to allow the use of ordinary functions as HTTP handlers. If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.

We can pass in any function / method that has the same signature as `ServeHTTP` as a handler.

## JSON
For net/http response we can technically write any content we want. But in most cases our data is not read by humans but consumed by other services / applications. So having structured data is a big part of REST. In the olden days `XML` was the go to format for REST but now a days most REST API use `JSON` as the format for the data.

Go has great support for `JSON` out of the box.

Once you run this version of the application you can test the output

```bash
curl -v localhost:7999/user
```

Your output would looks somehting like this

```shell
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 7999 (#0)
> GET /user HTTP/1.1
> Host: localhost:7999
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Wed, 05 Aug 2020 05:46:31 GMT
< Content-Length: 64
< 
{"username":"moficodes","email":"moficodes@gmail.com","age":27}
* Connection #0 to host localhost left intact
* Closing connection 0
```

## HTTP Verbs
In the previous example we had a method called `getUser`. It is pretty clear it is a GET operation. But if you were to make a request like

```bash
curl -x POST "localhost:7999/user
```

You would still see the same output. 

This might seem weird since we had named our function `getUser` we probably was expecting to only do get in that route. But naming our function has no effect on the HTTP verb we allow. 

```go
http.HandleFunc("/user", s.getUser)
```

This is the line of code that registers that route. And all it says is anytime a request comes we will be serving that using the `s.getUser` HandlerFunc.

net/http does not have direct support for http verbs like GET, POST, PUT, DELETE etc. But it is easily implementable. This is also a selling poing for other library / framework that is more developer friendly in letting us control the allowed http methods.

## REST Routes
In REST the same route can mean different thing based on the HTTP method of the request.

For example our `/user` route can be a retrieve on `GET` and update on `PUT`

To test out put

```
curl -X PUT -d '{"username":"mofi","email":"moficodes@ibm.com","age":27}' localhost:7999/user
```

Here we are updating the username to `mofi` and email to `moficodes@ibm.com`. 

We should see `{"update": "ok"}` Print.

We can also try sending a malformed JSON string. This will return a `405` Bad Request.

## Handler vs HandlerFunc vs *HandlerMethod

> The HandlerMethod is in * because that is not a real term go uses. I am using it to differentiate between a HandlerFunc attached to a type vs a regular function that has the same type as ServeHTTP function.

The main difference between all of these is developer experience. Creating a new Handler for each of our route is probably not going to be fun. HandlerFunc or methods on structs make is really flexible to build our routes. Choose whichever fits the problem at hand best. For our `/user` endpoint we needed access to a resource that was attached to our `server` struct. So using a `HandlerMethod` was ideal. But if we had a different route that did not have any dependency like that I might have opted for a regular function.

## Path Parameter
Matching complex routes using the standard library net/http package is quite difficult. We just added a new route `/base64/` that returns base64 representation anything sent after the `/`.

```bash
curl localhost:7999/base64/hello-world
```

We should see output `aGVsbG8td29ybGQ=%`

Doing anything more complex would be too cumbersome. This is another reason many developers will reach for a library/framework. And if your application has needs for dynamic routes net/http might not be the best solution.

For example matching `/user/1/blog/4/comment/2` is a very common REST pattern for routes. And with just net/http it will be alot of work to implement something that does this.

Luckily we have many great libraries and frameworks at our disposal.

We will move on to gorilla/mux now.

>The example used here does not follow REST Philosophy. We can have a longer discussion about why something is or isn't restful. The function is more of an action than a representation of some entity.