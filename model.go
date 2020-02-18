package main

import "database/sql"

type Todo struct {
	ID            int    `json:"id" form:"id"`
	Title         string `json:"title" form:"title"`
	Stared        bool   `json:"stared" form:"stared"`
	DeadlineStart Date   `json:"deadline_start" form:"deadline_start"`
	DeadlineEnd   Date   `json:"deadline_end" form:"deadline_end"`
	Comment       string `json:"comment" form:"comment"`
	Completed     bool   `json:"completed" form:"completed"`
}

type Todos []Todo

func queryTodoTable(db *sql.DB) Todos {
	sqlStatement := "SELECT id, title, stared, deadline_start, deadline_end, comment, completed FROM todos;"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	todos := make(Todos, 0)

	for rows.Next() {
		var todo Todo
		switch err := rows.Scan(
			&todo.ID, &todo.Title, &todo.Stared,
			&todo.DeadlineStart, &todo.DeadlineEnd,
			&todo.Comment, &todo.Completed,
		); err {
		case nil:
			todos = append(todos, todo)
		default:
			if err != nil {
				panic(err)
			}
		}
	}

	return todos
}

func addTodo(db *sql.DB, todo Todo) Todo {
	var newTodo Todo
	sqlStatement := `INSERT INTO todos (title, stared, deadline_start, deadline_end, comment, completed) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, title, stared, deadline_start, deadline_end, comment, completed;`
	err := db.QueryRow(
		sqlStatement, todo.Title, todo.Stared,
		todo.DeadlineStart, todo.DeadlineEnd,
		todo.Comment, todo.Completed,
	).Scan(
		&newTodo.ID, &newTodo.Title, &newTodo.Stared,
		&newTodo.DeadlineStart, &newTodo.DeadlineEnd,
		&newTodo.Comment, &newTodo.Completed)
	if err != nil {
		panic(err)
	}
	return newTodo
}

func getTodo(db *sql.DB, id string) Todo {
	var todo Todo
	sqlStatement := `SELECT id, title, stared, deadline_start, deadline_end, comment, completed FROM todos
		where id = $1;`
	err := db.QueryRow(sqlStatement, id).Scan(
		&todo.ID, &todo.Title, &todo.Stared,
		&todo.DeadlineStart, &todo.DeadlineEnd,
		&todo.Comment, &todo.Completed)
	if err != nil {
		panic(err)
	}

	return todo
}

func updateTodo(db *sql.DB, id string, todo Todo) Todo {
	var newTodo Todo
	sqlStatement := `UPDATE todos SET title = $2, stared = $3,
			deadline_start = $4, deadline_end = $5, comment = $6, completed = $7
		WHERE id = $1 
		RETURNING id, title, stared, deadline_start, deadline_end, comment, completed;`
	err := db.QueryRow(
		sqlStatement, id, todo.Title, todo.Stared,
		todo.DeadlineStart, todo.DeadlineEnd,
		todo.Comment, todo.Completed,
	).Scan(
		&newTodo.ID, &newTodo.Title, &newTodo.Stared,
		&newTodo.DeadlineStart, &newTodo.DeadlineEnd,
		&newTodo.Comment, &newTodo.Completed)
	if err != nil {
		panic(err)
	}
	return newTodo
}

func deleteTodo(db *sql.DB, id string) {
	sqlStatement := "DELETE FROM todos WHERE id = $1;"
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}
}
