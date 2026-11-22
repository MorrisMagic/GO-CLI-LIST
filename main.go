package main

import (
	types "github.com/MorrisMagic/GO-CLI-LIST.git/Types"
)

func main() {
	todos := types.Todos{}
	storage := NewStorage[types.Todos]("todos.json")
	storage.Load(&todos)
	todos.Add("buy dildo for aziz")
	todos.Add("sleep with yassine")

	todos.Toggle(0)
	todos.Print()
	storage.Save(todos)

}
