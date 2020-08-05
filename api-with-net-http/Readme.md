# REST API with Standard Library net/http

## Run the example
```
git checkout origin/standard-library-net-http-02
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