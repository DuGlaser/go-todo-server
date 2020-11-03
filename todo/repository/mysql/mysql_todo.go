package mysql

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/DuGlaser/go-todo-server/domain"
)

type mysqlTodoRepository struct {
	Conn *sql.DB
}

const (
	queryGetAllTodo  = "SELECT id, title, description, status FROM todos;"
	queryGetByIdTodo = "SELECT id, title, description, status FROM todos WHERE id=?;"
	queryInsertTodo  = "INSERT INTO todos(title, description, status) VALUES(?,?,?);"
	queryUpdateTodo  = "UPDATE todos SET title=?, description=?, status=? WHERE id=?;"
	queryDeleteTodo  = "DELETE FROM todos WHERE id=?;"
)

func NewMysqlTodoRepository(Conn *sql.DB) domain.TodoRepository {
	return &mysqlTodoRepository{Conn}
}

func (m *mysqlTodoRepository) GetAll() (domain.Todos, error) {
	stmt, err := m.Conn.Prepare(queryGetAllTodo)
	if err != nil {
		return nil, errors.New("error when trying to get todos")
	}
	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error when trying to get todos")
	}

	todos := make(domain.Todos, 0)

	for rows.Next() {
		var todo domain.Todo

		scanErr := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.Status,
		)
		if scanErr != nil {
			fmt.Println(err)
			return nil, errors.New("error when trying to get todos")
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (m *mysqlTodoRepository) GetByID(id int64) (domain.Todo, error) {
	var t domain.Todo
	stmt, err := m.Conn.Prepare(queryGetByIdTodo)
	if err != nil {
		return t, errors.New("error when trying to get todo by id")
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	err = row.Scan(
		&t.ID,
		&t.Title,
		&t.Description,
		&t.Status,
	)
	if err != nil {
		return t, errors.New("error when trying to get todo by id")
	}

	return t, nil
}

func (m *mysqlTodoRepository) Store(t *domain.Todo) error {
	stmt, err := m.Conn.Prepare(queryInsertTodo)
	if err != nil {
		return errors.New("error when trying to save todo")
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		t.Title,
		t.Description,
		t.Status,
	)
	if err != nil {
		return errors.New("error when trying to save todo")
	}

	todoID, err := result.LastInsertId()
	if err != nil {
		return errors.New("error when trying to save todo")
	}

	t.ID = todoID

	return nil
}

func (m *mysqlTodoRepository) Update(t *domain.Todo) error {
	stmt, err := m.Conn.Prepare(queryUpdateTodo)
	if err != nil {
		return errors.New("error when trying to update todo")
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		t.Title,
		t.Description,
		t.Status,
		t.ID,
	)
	if err != nil {
		return errors.New("error when trying to update todo")
	}

	return nil
}

func (m *mysqlTodoRepository) Delete(id int64) error {
	stmt, err := m.Conn.Prepare(queryDeleteTodo)
	if err != nil {
		fmt.Println(err)
		return errors.New("error when trying to delete todo")
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
		return errors.New("error when trying to delete todo")
	}

	return nil
}
