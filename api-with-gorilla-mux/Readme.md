# Gorilla Mux Router

In the previous section we saw how we can create a simele REST endpoint with net/http. We saw there were some limitations. But In most cases when we don't need complex path matching, net/http works just fine. 

Lets see how gorilla mux addresses the issues we saw with net/http.

## Run the example

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
  r.HandleFunc("/users", getAllUsers).Methods(http.MethodGet)

  r.HandleFunc("/courses", getCoursesWithInstructorAndAttendee).
    Queries("instructor", "{instructor:[0-9]+}", "attendee", "{attendee:[0-9]+}").
    Methods(http.MethodGet)

  r.HandleFunc("/courses", getAllCourses).Methods(http.MethodGet)
  r.HandleFunc("/instructors", getAllInstructors).Methods(http.MethodGet)
```

We will also need to create corresoponding methods for each routes.

With these our api should behave exactly the same.

```bash
curl localhost:7999/api/v1/users
```

## Path Params

With net/http we saw a simple example of how we can do path params. We gorilla mux this becomes significantly easier (this is also where the libraries and framework differ in implementations).

For this part we start with a more complex example. Lets imagine a application that is keeping track of courses, instructors and attendees to these courses.

We have HandlerFunc for returing all users, instructors and courses. As well as getting individual user, course or instructor. And thats the route we are interested in here.

```go
  api.HandleFunc("/users/{id}", getUserByID).Methods(http.MethodGet)
```

In this route we have a path param `{id}` which we will access in the `getUserByID` function. 

```go
  pathParams := mux.Vars(r)
```

Thats all we have to do to get access to all the path parameters in `map[string]string`. Although depending on the use case that might not be enough. For example in our use case we expect the param to be an int. Its simple to convert the data to the right format. 

```go
if val, ok := pathParams["id"]; ok {
  id, err = strconv.Atoi(val)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(`{"error": "need a valid id"}`))
    return
  }
}
```

Sending the error message and setting the status code helps the consumer of this api to take appropriate actions.

Once we hace access to the id we can search our list for the appropirate resource. In our case its looping through an array. In most cases it would be querying some sort of database.

## Query Parameters

We can get access the the query map by calling `r.URL.Query()`. This returns a object of type `url.Value` which has the underlying type of `map[string][]string`. It is a map of array of string. Query parameters can be repeated thats why its an array.

We can also use `r.FormValue(key)` to get a query. 
>This method can get both query param and form value in request body.

For our example say we want to filter our all lists with some criterias like users with interests, instructors with expertise and courses with topics.

Lets look at `getAllUsers` function.

```go
  query := r.URL.Query()
```

With this we get access the the `map[string][]string`. Some applications mandate only one query per key. For that we can use `query.Get(key)` to get the first instance of the key in the query map. For our example we are allowing user to do multiple query of the same key. So we access the array.

```go
  interests, ok := query["interest"]
```

In go on map data access we get a second value which says if the value was there or not. If there is no query with the key we just return the whole result. In this example we are treating the queries as an `and` relation. If we query for `NodeJS` and `AI` all the users with interests in both of the topic will be returned. It is also possible to treat this like a or relation albeit not that common.

## Match Query

In the previous example we had a query parameter that was optional. But what if we want to go to a different route based on a query parameter?

Yes we can.

```go
r.HandleFunc("/courses", getCoursesWithInstructorAndAttendee).
  Queries("instructor", "{instructor:[0-9]+}", "attendee", "{attendee:[0-9]+}").
  Methods(http.MethodGet)
```

This route will only match request with query `instructor` and `attendee` and the value type integer. Anything else will match the other route with the same path. Or return 404 if nothing else matches.

>The route matching is top to bottom read operation. If you have more specific route you should put them above otherwise a generic route might catch and respond.

## Sub Router

So far we have been adding all our routes at the top level of our hostname. i.e. `localhost:7999/<route>`. But there are times when it is desired to have routes that are grouped together based on some criteria. For example we might want all our authentication routes grouped together under `/auth` or we could want to version our api with `/api/v1`. With a sub-router it is possible to apply rules and logic to a group of routes instead of applying these rules individually.

To create a sub-router 

```go
  api := r.PathPrefix("/api/v1").Subrouter()
```

We can then treat `api` as if it were a Router and add new routes to it. Any route added to this subrouter will be prefixed with `/api/v1` so the path `/users` become `/api/v1/users`. 

We can still test that all our routes work as expected.

```bash
curl localhost:7999/api/v1/instructors
```
