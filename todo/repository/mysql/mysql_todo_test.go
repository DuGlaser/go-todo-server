package mysql_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/DuGlaser/go-todo-server/db"
	"github.com/DuGlaser/go-todo-server/domain"
	"github.com/DuGlaser/go-todo-server/todo/repository/mysql"
	"github.com/ory/dockertest"
	"github.com/stretchr/testify/assert"
)

const (
	queryDropTodo = "DROP TABLE IF EXISTS todos"
	queryUpTodo   = "CREATE TABLE IF NOT EXISTS todos(" +
		"id int PRIMARY KEY NOT NULL AUTO_INCREMENT ," +
		"title VARCHAR(50)," +
		"description VARCHAR(50)," +
		"status VARCHAR(15)" +
		")"
)

var (
	DBConn = db.NewTestDB()
)

func Drop() error {
	ctx := context.Background()

	stmt, err := DBConn.PrepareContext(ctx, queryDropTodo)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

func Up() error {
	ctx := context.Background()

	stmt, err := DBConn.PrepareContext(ctx, queryUpTodo)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

func TestMain(m *testing.M) {
	testDB := db.NewTodoConfig().DB.Test
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker : %s", err)
	}

	opts := dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "5.7",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=" + testDB.Password,
			"MYSQL_PASSWORD=" + testDB.Password,
			"MYSQL_USER=" + testDB.Username,
		},
		ExposedPorts: []string{testDB.Port},
	}

	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err.Error())
	}

	defer DBConn.Close()

	err = Drop()
	if err != nil {
		panic(err)
	}

	err = Up()
	if err != nil {
		panic(err)
	}

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestStore(t *testing.T) {
	TestTodo := &domain.Todo{
		Title:       "test_title",
		Description: "test_description",
		Status:      domain.Doing,
	}
	repo := mysql.NewMysqlTodoRepository(DBConn)
	err := repo.Store(TestTodo)

	assert.NoError(t, err)
	assert.NotNil(t, TestTodo.ID)
}

func TestGetByID(t *testing.T) {
	TestTodo := &domain.Todo{
		Title:       "test_title",
		Description: "test_description",
		Status:      domain.Doing,
	}
	repo := mysql.NewMysqlTodoRepository(DBConn)
	Drop()
	Up()

	err := repo.Store(TestTodo)
	if err != nil {
		t.Fatalf(err.Error())
	}

	todo, err := repo.GetByID(TestTodo.ID)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.NoError(t, err)
	assert.EqualValues(t, *TestTodo, todo)
}

func TestGetAll(t *testing.T) {
	repo := mysql.NewMysqlTodoRepository(DBConn)
	Drop()
	Up()

	TestTodo1 := &domain.Todo{
		Title:       "test_title",
		Description: "test_description",
		Status:      domain.Doing,
	}

	TestTodo2 := &domain.Todo{
		Title:       "test_title",
		Description: "test_description",
		Status:      domain.Doing,
	}

	_ = repo.Store(TestTodo1)
	_ = repo.Store(TestTodo2)

	TestTodos := make(domain.Todos, 0)
	TestTodos = append(TestTodos, *TestTodo1)
	TestTodos = append(TestTodos, *TestTodo2)

	todos, err := repo.GetAll()
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.NoError(t, err)
	assert.EqualValues(t, TestTodos, todos)
}

func TestUpdate(t *testing.T) {
	repo := mysql.NewMysqlTodoRepository(DBConn)
	Drop()
	Up()

	TestTodo := &domain.Todo{
		Title:       "test_title",
		Description: "test_description",
		Status:      domain.Doing,
	}
	_ = repo.Store(TestTodo)

	TestTodo.Title = "test_title_updated"
	TestTodo.Description = "test_description_updated"
	err := repo.Update(TestTodo)

	assert.NoError(t, err)

	todo, _ := repo.GetByID(TestTodo.ID)

	assert.EqualValues(t, *TestTodo, todo)
}

func TestDelete(t *testing.T) {
	repo := mysql.NewMysqlTodoRepository(DBConn)
	Drop()
	Up()

	TestTodo := &domain.Todo{
		Title:       "test_title",
		Description: "test_description",
		Status:      domain.Doing,
	}
	_ = repo.Store(TestTodo)

	todo, _ := repo.GetByID(TestTodo.ID)
	assert.EqualValues(t, *TestTodo, todo)

	repo.Delete(TestTodo.ID)

	todo, err := repo.GetByID(TestTodo.ID)
	assert.NotNil(t, err)
	assert.EqualValues(t, domain.Todo{}, todo)
}
