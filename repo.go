package main

type ToDoRepo interface {
	All() ([]ToDo, error)
	Find(id string) (*ToDo, error)
	Save(t *ToDo) error
	Delete(t *ToDo) error
}
