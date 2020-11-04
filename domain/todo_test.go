package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidTodo(t *testing.T) {
	todo := Todo{
		ID:          1,
		Title:       "title",
		Description: "description",
		Status:      "Doing",
	}

	err := todo.Validate()
	if err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, err)
	assert.EqualValues(t, "title", todo.Title)
	assert.EqualValues(t, "description", todo.Description)
	assert.EqualValues(t, Doing, todo.Status)
}

func TestInvalidTitleTodo(t *testing.T) {
	todo := Todo{
		ID:          1,
		Title:       "",
		Description: "description",
		Status:      "Doing",
	}

	err := todo.Validate()
	assert.EqualValues(t, "title is empty", err.Message)
}

func TestInvalidDescriptionTodo(t *testing.T) {
	todo := Todo{
		ID:          1,
		Title:       "title",
		Description: "",
		Status:      "Doing",
	}

	err := todo.Validate()
	assert.EqualValues(t, "description is empty", err.Message)
}

func TestInvalidStatusTodo(t *testing.T) {
	todo := Todo{
		ID:          1,
		Title:       "title",
		Description: "description",
		Status:      "status",
	}

	err := todo.Validate()
	assert.EqualValues(t, "invalid status", err.Message)
}

func TestStatusIsDone(t *testing.T) {
	todo := Todo{
		ID:          1,
		Title:       "title",
		Description: "description",
		Status:      "Done",
	}

	err := todo.Validate()

	assert.Nil(t, err)
	assert.EqualValues(t, Done, todo.Status)
}

func TestTrimSpace(t *testing.T) {
	todo := Todo{
		ID:          1,
		Title:       "  title  ",
		Description: "  description  ",
		Status:      "Done",
	}

	_ = todo.Validate()

	assert.EqualValues(t, "title", todo.Title)
	assert.EqualValues(t, "description", todo.Description)
}
