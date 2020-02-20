# Todolist server

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
