package main

import (
	"context"
	"flag"
	f "fmt"
	"log"
	s "strings"

	"github.com/ryand67/teams-directory-go/credentials"
	"github.com/ryand67/teams-directory-go/firebase"
	"github.com/ryand67/teams-directory-go/team"
)

func main() {
	// init "global" context
	ctx := context.Background()

	unFlag := flag.String("username", "", "username")
	pwFlag := flag.String("password", "", "password")
	flag.Parse()

	if !credentials.EmailValid(*unFlag) {
		f.Println("Provided command line username invalid email format")
	}

	// Creates new firebase app instance
	app := firebase.NewApp(ctx)

	// If username/password handed in via command line
	if *unFlag != "" && *pwFlag != "" && credentials.EmailValid(*unFlag) {
		c, err := credentials.Login(ctx, app, *unFlag, *pwFlag)
		if err != nil {
			log.Fatalf(err.Error())
		}
		ctx = c
	} else {
		// Login or sign up, returns updated context
		c, err := credentials.PromptLogin(app, ctx)
		if err != nil {
			log.Fatalf(err.Error())
		}
		ctx = c
	}

	for {
		f.Print("> ")
		var cmd string
		f.Scanln(&cmd)

		switch s.ToLower(cmd) {
		case "team-list", "tl":
			team.TeamList(ctx, app, "team")
		case "exit", "x", "e":
			log.Fatalf("Program terminated by user.")
		default:
			f.Println("Invalid command, exec 'help' for list of commands.")
		}
	}
}
