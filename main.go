package main

import (
	types "github.com/MorrisMagic/GO-CLI-LIST.git/Types"
)

func main() {
	todos := types.Todos{}
	storage := NewStorage[types.Todos]("todos.json")
	storage.Load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	storage.Save(todos)

}
