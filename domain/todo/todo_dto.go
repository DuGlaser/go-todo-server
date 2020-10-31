package todo

import "errors"

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (t *Todo) Validate() error {
	if t.Title == "" {
		return errors.New("title is empty")
	}

	if t.Description == "" {
		return errors.New("description is empty")
	}

	if t.Status == "" {
		return errors.New("status is empty")
	}

	return nil
}
