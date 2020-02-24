# Todolist server

## Setup

Create env.ini file:  

```ini
[System]
ip = 0.0.0.0
port = 5000
auth_key = secret key

[Database]
host = test_postgres
port = 5432
user = postgres
password = password
dbname = todo_database
```

## Run

```shell
# run
$ go run main.go middleware.go route.go

#build
$ go build
```

## Table schema

![schema](schema.PNG)  

```sql
CREATE TABLE todos(
    id serial PRIMARY KEY,
    title VARCHAR (80) NOT NULL,
    stared Boolean NOT NULL,
    deadline_start DATE,
    deadline_end DATE,
    comment TEXT,
    completed Boolean NOT NULL
);

CREATE TABLE todo_sort(
    todo_id int UNIQUE NOT NULL REFERENCES todos(id) ON DELETE CASCADE,
    sort_index int UNIQUE NOT NULL
);

CREATE TABLE todo_user(
    id serial PRIMARY KEY,
    username VARCHAR (80) NOT NULL,
    password VARCHAR (80) NOT NULL
);
```

## Docker

Modify the host value to test_postgres which section is Database on env.ini file.  

```shell
# db
$ docker run --name test_postgres -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres

# build and run
$ docker build -t go-todolist-server .
$ docker run --name go-todolist-server -p 5000:5000 --link test_postgres:test_postgres -d go-todolist-server
```
