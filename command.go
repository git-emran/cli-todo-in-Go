package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Edit   string
	Del    int
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new Todo Specify Title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit the todo by index and specify a new Title")

	flag.IntVar(&cf.Del, "del", -1, "Specify the Todo by index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify todo by index to toggle completion")

	flag.BoolVar(&cf.List, "list", false, "List all Todo items")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: Invalid format for Edit. Please use index:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: Invalid index for edit")
			os.Exit(1)
		}

		todos.edit(index, parts[1])

	case cf.Toggle != 1:
		todos.delete(cf.Del)

	default:
		fmt.Println("Invalid Ccommand")
	}
}
