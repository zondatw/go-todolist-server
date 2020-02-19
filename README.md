# Todolist server

## Run

```shell
# run
$ export GOPATH=$(pwd):$GOPATH
$ go run main.go middleware.go route.go

#build
$ export GOPATH=$(pwd):$GOPATH
$ go build main.go middleware.go route.go
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
```
