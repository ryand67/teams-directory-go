package util

import f "fmt"

type Command struct {
	Name string
	Command []string
	Desc string
}

var commands []Command = []Command{
	{
		Name: "Add Team",
		Command: []string{"add-team", "at"},
		Desc: "Adds a team to your account",
	},
	{
		Name: "Team List",
		Command: []string{"team-list", "tl"},
		Desc: "Lists all available teams",
	},
	{
		Name: "Help",
		Command: []string{"help", "doc", "d", "h"},
		Desc: "Prints all available commands",
	},
	{
		Name: "Exit",
		Command: []string{"exit", "x", "e"},
		Desc: "Exits the program",
	},
}

func Help() {
	f.Println("NAME -- COMMANDS -- DESC")
	for _, v := range commands {
		f.Println(v.Name, "--", v.Command, "--", v.Desc)
	}
}