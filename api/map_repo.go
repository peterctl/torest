package main

import (
	"errors"
	"fmt"
)

type MapToDoRepo struct {
	todos map[string]ToDo
}

func NewMapRepo() *MapToDoRepo {
	return &MapToDoRepo{
		todos: make(map[string]ToDo, 100),
	}
}

func (r *MapToDoRepo) All() (res []ToDo, err error) {
	res = make([]ToDo, len(r.todos))
	i := 0
	for _, todo := range r.todos {
		res[i] = todo
		i++
	}
	err = nil
	return
}

func (r *MapToDoRepo) Find(id string) (*ToDo, error) {
	todo, ok := r.todos[id]
	if ok {
		return &todo, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Could not find todo %s", id))
	}
}

func (r *MapToDoRepo) Save(t *ToDo) error {
	r.todos[t.Id] = *t
	return nil
}

func (r *MapToDoRepo) Delete(t *ToDo) error {
	_, ok := r.todos[t.Id]
	if ok {
		delete(r.todos, t.Id)
		return nil
	}
	return errors.New(fmt.Sprintf("Could not find todo %s", t.Id))
}
