package todo

import (
	"errors"
	"fmt"

	"github.com/DuGlaser/go-todo-server/db"
)

const (
	queryGetAllTodo  = "SELECT id, title, description, status FROM todos;"
	queryGetByIdTodo = "SELECT id, title, description, status FROM todos WHERE id=?;"
	queryInsertTodo  = "INSERT INTO todos(title, description, status) VALUES(?,?,?);"
	queryUpdateTodo  = "UPDATE todos SET title=?, description=?, status=? WHERE id=?;"
	queryDeleteTodo  = "DELETE FROM todos WHERE id=?;"
)

func (t *Todo) GetAll() (Todos, error) {
	stmt, err := db.TodoDB.GetClient().Prepare(queryGetAllTodo)
	if err != nil {
		return nil, errors.New("error when trying to get todos")
	}
	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error when trying to get todos")
	}

	todos := make(Todos, 0)

	for rows.Next() {
		var todo Todo

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

func (t *Todo) GetByID() error {
	stmt, err := db.TodoDB.GetClient().Prepare(queryGetByIdTodo)
	if err != nil {
		return errors.New("error when trying to get todo by id")
	}
	defer stmt.Close()

	row := stmt.QueryRow(
		t.ID,
	)
	err = row.Scan(
		&t.ID,
		&t.Title,
		&t.Description,
		&t.Status,
	)
	if err != nil {
		return errors.New("error when trying to save todo by id")
	}

	return nil
}

func (t *Todo) Save() error {
	stmt, err := db.TodoDB.GetClient().Prepare(queryInsertTodo)
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

func (t *Todo) Update() error {
	stmt, err := db.TodoDB.GetClient().Prepare(queryUpdateTodo)
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

func (t *Todo) Delete() error {
	stmt, err := db.TodoDB.GetClient().Prepare(queryDeleteTodo)
	if err != nil {
		fmt.Println(err)
		return errors.New("error when trying to delete todo")
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		t.ID,
	)
	if err != nil {
		fmt.Println(err)
		return errors.New("error when trying to delete todo")
	}

	return nil
}
