package main

import (
	"flag"
	"fmt"
)

func main() {
	var addr string
	var db string
	flag.StringVar(&addr, "addr", "127.0.0.1:8000", "Listening address.")
	flag.StringVar(&db, "db", ":memory:", "Sqlite3 database path.")
	flag.Parse()

	repo, err := NewDbRepo(db)
	if err != nil {
		panic(err)
	}

	err = repo.Initialize()
	if err != nil {
		panic(err)
	}

	app := NewApp(repo)
	fmt.Println(app.ListenAndServe(addr))
}
