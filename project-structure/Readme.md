# Project Structure

In this chapter we will discuss project structure. In go project structure does not make any difference in performance or functionality. It is a matter of preference and organization of code.

## Flat Structure

As the name suggest in this structure the code is laid out in a single folder. This structure is great to start out as everything is accessible everywhere. I recommend starting with this.

```
rest-api
├── Readme.yaml
├── data
│   ├── courses.json
│   ├── instructors.json
│   └── users.json
├── go.mod
├── go.sum
├── handlers.go
├── main.go
├── middlewares.go
└── types.go
```

Our Course API was already in a flat structure. But better organized version can be seen [here](https://github.com/moficodes/restful-go-api/tree/project-structure-01/project-structure/flat-structure). 

## MVC Structure

This structure follows the MVC pattern to separate out Model, Controller and Data layer. If you are coming from a MVC structure project from languages like C#, Java etc this might be more comfortable to start with.

```
rest-api
├── Readme.md
├── model
	└── user.go
├── controller
	└── user.go
├── go.mod
├── go.sum
└── main.go
```

## Pkg Structure

This is a very common pattern in the go ecosystem. 

```
rest-api
pkg-structure
├── Readme.md
├── cmd
│   └── web
│       └── main.go
├── data
│   ├── courses.json
│   ├── instructors.json
│   └── users.json
├── go.mod
├── go.sum
├── internal
│   ├── handler
│   │   ├── authenticated.go
│   │   ├── course.go
│   │   ├── instructor.go
│   │   ├── server.go
│   │   ├── types.go
│   │   ├── user.go
│   │   └── utils.go
│   └── middleware
│       ├── jwt.go
│       └── types.go
└── pkg
    └── middleware
        └── logger.go
```

Our Course API is converted in the pkg structure [here](https://github.com/moficodes/restful-go-api/tree/project-structure-01/project-structure/pkg-structure). 

_What to put in pkg vs internals?_ 

> The way I reason about it is, the things in pkg could be used for other projects just by copying the code or directly importing as a module. Internal has logic that only pertain to this application. 

I have using this structure for my project recently. One great benefit is that, this structure promotes the idea of multiple binary from the same code base. Inside my `cmd` folder I could easily create `cli` binary that uses the same underlying code to create a CLI tool. Not to say it is impossible to do so in the other structure. But this structure suits that usecase perfectly.

_Which one should I choose?_

> Like most things in technology the answer is "It depends". If you are starting a new project you are pretty much free to structure your code in any way you like. 

