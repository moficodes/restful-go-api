# Database

Most Rest API represents some state. That state is usually stored in Database. Go has great support for all kinds of database. 

Depending on how data is stored and accessed databases can be of many types.

1. Key Value (redis, memcached, etcd)
2. Wide Column (apache hbase)
3. Document Store (MongoDB, Firestore, CouchDB)
4. Relational (Mysql, Postgres, CockroachDB)
5. Graph (neo4j, janus graph)
6. Search Engine (elastic, algolia)
7. Multi Modal (fauna, cosmos)

Each database type has their use cases. For rest api, usually we see the use of Document Store or Relational Database. 

## Which type of database to use?

The database type to use depends on the need. But to be more precise it depends on the structure of the data. One thing that gets floated around quite alot is structured vs unstructured data. Unstructured data usually is not very useful. For that matter data is not super useful unless we put some meaning behind it. If we have data that we don't know the exact representation of Document Store are super useful. We can get any old json and store it. But to be able to query that data we need to have some level of structure. So when we hear Document Store can store unstructured data it doesn't mean our application can actually do anything useful with it unless we have some idea of what that data might be. 

Rest deals with states of things. Often that state is related to other things. We can represent the state of these things in a table format using relational databases. Relational databases are probably the most widely used databases in the world. In performance they are generally faster than document store for complex query. 

Key Value databases are probably most notably used for cachning rest api. There are a number of other uses for them. But for rest api they are usually not the most suitable.

## Go + Postgres

We will be using postgres as our database of choice. The example does not do anything specific for postgres, so chances are it will work with mysql variant of databases as well. 

For using postgres and go we have a few choice of libraries. 

* [pq](https://github.com/lib/pq)
* [pgx](github.com/jackc/pgx)
* [sql](https://golang.org/pkg/database/sql/)
* [gorm](https://github.com/go-gorm/gorm)
* [go-pg](https://github.com/go-pg/pg)

gorm and go-pg are orms. Whether or not orms are good or bad are beyond the scope of this workshop. I personally switch between pgx and standard library sql package. pq is no longer being actively developed and is now suggesting to use pgx in its place. 

## Pooling connection

By default go rest api server serves requests concurrently. So if more than one request comes to our server that requires database access we could either create a new connection every request or try to reuse requests. Creating a new connection might lead to problems because we can hit the limit of connections in our database. Trying to reuse connection seems like a good idea but we would have to manage locks to make sure access to connection from concurrent go routines are safe. Luckily `pgx` has `pgxpool` that manages a concurrency safe connection pool for us. 

## Setup local database env

With docker we can setup a ephemeral postgres instance

```bash
docker run --name postgresql-container -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres
```

This uses the `postgres` docker container and sets up a database on port `5432` with password = password. 

> Suffice to say this is not something you want to do for production use. 

We have a database init script that can setup some tables and some initial data. 

You can either run the sql queries using some tool like [Dbeaver](https://dbeaver.io/) or [DataGrip](https://www.jetbrains.com/datagrip/). If you have any Jet Brains producs like Goland or Intellij you already have a database tool built in. You can also use command line tool `psql` 



This will prompt you for password which is password. It will then load up our data. 

## Run the example

```bash
git checkout origin/rest-api-database-01
```

If you are not already in the folder

```bash
cd rest-api-database
```

To load the data using psql.

```bash
psql -h localhost -p 5432 -d postgres -f database/migration/dump.sql -U postgres
```

```bash
go run cmd/web/main.go
```

```bash
curl localhost:7999/api/v1/users/1
```

