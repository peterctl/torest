package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var CREATE_TABLE = `CREATE TABLE IF NOT EXISTS todos (
	id VARCHAR(32) NOT NULL PRIMARY KEY,
	name TEXT NOT NULL,
	completed BOOL NOT NULL
);`

type DbRepo struct {
	Db *sql.DB
}

func NewDbRepo(dbPath string) (*DbRepo, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	r := &DbRepo{
		Db: db,
	}

	return r, nil
}

func (r *DbRepo) Initialize() error {
	_, err := r.Db.Exec(CREATE_TABLE)
	return err
}

func (r *DbRepo) All() ([]ToDo, error) {
	rows, err := r.Db.Query("SELECT * FROM todos;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var t ToDo
	res := make([]ToDo, 0)
	for rows.Next() {
		rows.Scan(&t.Id, &t.Name, &t.Completed)
		res = append(res, t)
	}

	return res, nil
}

func (r *DbRepo) Find(id string) (*ToDo, error) {
	row, err := r.Db.Query("SELECT * FROM todos WHERE id = ? LIMIT 1;", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var t ToDo
	if row.Next() {
		row.Scan(&t.Id, &t.Name, &t.Completed)
		return &t, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Error: Todo '%s' not found in database.", id))
	}
}

func (r *DbRepo) Save(t *ToDo) error {
	if res, _ := r.Find(t.Id); res == nil {
		return r.Insert(t)
	} else {
		return r.Update(t)
	}
}

func (r *DbRepo) Insert(t *ToDo) error {
	_, err := r.Db.Exec("INSERT INTO todos VALUES (?, ?, ?);",
		t.Id, t.Name, t.Completed)
	return err
}

func (r *DbRepo) Update(t *ToDo) error {
	statement := "UPDATE todos SET name = ?, completed = ? WHERE id = ?;"
	_, err := r.Db.Exec(statement, t.Name, t.Completed, t.Id)
	return err
}

func (r *DbRepo) Delete(t *ToDo) error {
	_, err := r.Db.Exec("DELETE FROM todos WHERE id = ?", t.Id)
	return err
}
