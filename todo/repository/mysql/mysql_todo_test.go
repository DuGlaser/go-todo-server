package mysql_test

import (
	"context"
	"log"
	"net/http"
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

	stmt, restErr := DBConn.PrepareContext(ctx, queryDropTodo)
	if restErr != nil {
		return restErr
	}
	defer stmt.Close()

	_, restErr = stmt.ExecContext(ctx)
	if restErr != nil {
		return restErr
	}
	return nil
}

func Up() error {
	ctx := context.Background()

	stmt, restErr := DBConn.PrepareContext(ctx, queryUpTodo)
	if restErr != nil {
		return restErr
	}
	defer stmt.Close()

	_, restErr = stmt.ExecContext(ctx)
	if restErr != nil {
		return restErr
	}
	return nil
}

func TestMain(m *testing.M) {
	testDB := db.NewTodoConfig().DB.Test
	pool, restErr := dockertest.NewPool("")
	if restErr != nil {
		log.Fatalf("Could not connect to docker : %s", restErr)
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

	resource, restErr := pool.RunWithOptions(&opts)
	if restErr != nil {
		log.Fatalf("Could not start resource: %s", restErr.Error())
	}

	defer DBConn.Close()

	restErr = Drop()
	if restErr != nil {
		panic(restErr)
	}

	restErr = Up()
	if restErr != nil {
		panic(restErr)
	}

	code := m.Run()

	if restErr := pool.Purge(resource); restErr != nil {
		log.Fatalf("Could not purge resource: %s", restErr)
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
	restErr := repo.Store(TestTodo)

	assert.Nil(t, restErr)
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

	if restErr := repo.Store(TestTodo); restErr != nil {
		t.Fatalf(restErr.Message)
	}

	todo, restErr := repo.GetByID(TestTodo.ID)
	if restErr != nil {
		t.Fatalf(restErr.Message)
	}

	assert.Nil(t, restErr)
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

	todos, restErr := repo.GetAll()
	if restErr != nil {
		t.Fatalf(restErr.Message)
	}

	assert.Nil(t, restErr)
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
	restErr := repo.Update(TestTodo)

	assert.Nil(t, restErr)

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

	todo, restErr := repo.GetByID(TestTodo.ID)
	assert.EqualValues(t, restErr.Status, http.StatusNotFound)
	assert.EqualValues(t, domain.Todo{}, todo)
}
