package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	types "github.com/MorrisMagic/GO-CLI-LIST.git/Types"
)

type cmdFlags struct {
	Add    string
	Delete int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *cmdFlags {
	cf := cmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify titles")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index")
	flag.IntVar(&cf.Delete, "del", -1, "Specify a todo by index to delete a todo")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo by index to toggle")
	flag.BoolVar(&cf.List, "list", false, "list all todos")

	flag.Parse()
	return &cf

}

func (cf *cmdFlags) Execute(todos *types.Todos) {
	switch {
	case cf.List:
		todos.Print()
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error , invalid format for edit . Please use id:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error : invalid index for edit")
			os.Exit(1)
		}

		todos.Edit(index, parts[1])
	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle)
	case cf.Delete != -1:
		todos.Delete(cf.Delete)
	default:
		fmt.Println("Invalid command")
	}
}
