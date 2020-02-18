# Todolist server

## Table schema

CREATE TABLE todos(
    id serial PRIMARY KEY,
    title VARCHAR (80) NOT NULL,
    stared Boolean NOT NULL,
    deadline_start DATE,
    deadline_end DATE,
    comment TEXT,
    completed Boolean NOT NULL
);