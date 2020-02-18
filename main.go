package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	_ "github.com/lib/pq"
)

type Todo struct {
	ID            int    `json:"id" form:"id"`
	Title         string `json:"title" form:"title"`
	Stared        bool   `json:"stared" form:"stared"`
	DeadlineStart Date   `json:"deadline_start" form:"deadline_start"`
	DeadlineEnd   Date   `json:"deadline_end" form:"deadline_end"`
	Comment       string `json:"comment" form:"comment"`
	Completed     bool   `json:"completed" form:"completed"`
}

type Date time.Time

func (date Date) Value() (driver.Value, error) {
	return []byte(time.Time(date).Format(`"2006-01-02"`)), nil
}

func (date *Date) Scan(v interface{}) error {
	switch s := v.(type) {
	case time.Time:
		*date = Date(s)
	case []byte:
		newTime, err := time.Parse("2006-01-02", string(s))
		if err != nil {
			return err
		}
		*date = Date(newTime)
	case string:
		newTime, err := time.Parse("2006-01-02", s)
		if err != nil {
			return err
		}
		*date = Date(newTime)
	default:
		return fmt.Errorf("date: Unsupport scanning type %T", v)
	}
	return nil
}

func (date *Date) UnmarshalJSON(input []byte) error {
	newTime, err := time.Parse(`"2006-01-02"`, string(input))
	if err != nil {
		return err
	}

	*date = Date(newTime)
	return nil
}

func (date Date) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(date).Format(`"2006-01-02"`)), nil
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

var db *sql.DB

func main() {
	cfg, err := ini.Load("env.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	var (
		host     = cfg.Section("Database").Key("host").String()
		port     = cfg.Section("Database").Key("port").Value()
		user     = cfg.Section("Database").Key("user").String()
		password = cfg.Section("Database").Key("password").String()
		dbname   = cfg.Section("Database").Key("dbname").String()
	)

	connectionString := fmt.Sprintf(
		"host=%s port=%s "+
			"user=%s password=%s "+
			"dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := gin.Default()
	router.Use(CORSMiddleware)
	api := router.Group("/api")
	{
		api.GET("/todos", todosGET)
		api.POST("/todos", todosPOST)
		api.GET("/todos/:id", todoGET)
		api.PUT("/todos/:id", todoPUT)
		api.DELETE("/todos/:id", todoDELETE)
	}
	router.Run(":5000")
}

func CORSMiddleware(context *gin.Context) {
	context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	context.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE")

	if context.Request.Method == "OPTIONS" {
		context.AbortWithStatus(204)
		return
	}

	context.Next()
}

func todosGET(context *gin.Context) {
	todos := queryTodoTable(db)
	context.JSON(http.StatusOK, todos)
}

func todosPOST(context *gin.Context) {
	var todo Todo
	var newTodo Todo
	if err := context.ShouldBindJSON(&todo); err == nil {
		newTodo = addTodo(db, todo)
		context.JSON(http.StatusOK, newTodo)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func todoGET(context *gin.Context) {
	id := context.Param("id")
	todo := getTodo(db, id)
	context.JSON(http.StatusOK, todo)
}

func todoPUT(context *gin.Context) {
	id := context.Param("id")
	var todo Todo
	if err := context.ShouldBindJSON(&todo); err == nil {
		newTodo := updateTodo(db, id, todo)
		context.JSON(http.StatusOK, newTodo)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func todoDELETE(context *gin.Context) {
	id := context.Param("id")
	deleteTodo(db, id)
	context.JSON(http.StatusNoContent, gin.H{})
}
