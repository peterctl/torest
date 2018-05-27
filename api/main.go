package main

import (
	"fmt"
)

func main() {
	app := NewApp(NewMapRepo())
	fmt.Println(app.ListenAndServe(":8000"))
}
