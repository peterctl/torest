package main

import (
	"fmt"
)

type ToDo struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

func NewToDo(name string, completed bool) *ToDo {
	return &ToDo{
		Id:        RandString(32),
		Name:      name,
		Completed: completed,
	}
}

func (t *ToDo) String() string {
	var c byte
	if t.Completed {
		c = 'X'
	} else {
		c = ' '
	}
	return fmt.Sprintf("[%b] %s", c, t.Name)
}
