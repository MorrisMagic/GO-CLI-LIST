package main

import (
	types "github.com/MorrisMagic/GO-CLI-LIST.git/Types"
)

func main() {
	todos := types.Todos{}
	todos.Add("buy dildo for aziz")
	todos.Add("sleep with yassine")

	todos.Toggle(0)
	todos.Print()

}
