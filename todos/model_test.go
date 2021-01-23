package todos

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/zondaTW/go-todolist-server/lib"
)

func TestQueryTodoTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	deadlineStart, _ := time.Parse("2006-01-02", "2021-01-20")
	deadlineEnd, _ := time.Parse("2006-01-02", "2021-01-21")

	todosMockRows := sqlmock.NewRows([]string{"id", "title", "stared", "deadline_start", "deadline_end", "comment", "completed"}).
		AddRow(1, "Test title", false, deadlineStart, deadlineEnd, "Test comment", false)

	mock.ExpectQuery("^SELECT (.+) FROM todos*").
		WillReturnRows(todosMockRows)

	var todo Todo = Todo{
		ID:            1,
		Title:         "Test title",
		Stared:        false,
		DeadlineStart: lib.Date(deadlineStart),
		DeadlineEnd:   lib.Date(deadlineEnd),
		Comment:       "Test comment",
		Completed:     false,
	}
	todos := make(Todos, 0)
	todos = append(todos, todo)

	newTodos := queryTodoTable(db)
	assert.Equal(t, newTodos, todos, "they should be equal")
}