package util

import f "fmt"

type Command struct {
	Name string
	Command []string
	Desc string
	Category string
}

var commands []Command = []Command{
	{
		Name: "Add Team",
		Command: []string{"add-team", "at"},
		Desc: "Adds a team to your account",
		Category: "teams",
	},
	{
		Name: "Team List",
		Command: []string{"team-list", "tl"},
		Desc: "Lists all available teams",
		Category: "teams",
	},
	{
		Name: "Help",
		Command: []string{"help", "doc", "d", "h"},
		Desc: "Prints all available commands, takes a category argument",
		Category: "util",
	},
	{
		Name: "Exit",
		Command: []string{"exit", "x", "e"},
		Desc: "Exits the program",
		Category: "util",
	},
}

func Help(cat string) {
	f.Println("NAME -- COMMANDS -- DESC -- CATEGORY")
	for _, v := range commands {
		if cat != "" && cat == v.Category {
			f.Println(v.Name, " -- ", v.Command, " -- ", v.Desc, " -- ", v.Category)
		} else if cat == "" {
			f.Println(v.Name, " -- ", v.Command, " -- ", v.Desc, " -- ", v.Category)
		} 	
	}
}