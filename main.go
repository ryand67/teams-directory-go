package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/ryand67/teams-directory-go/pkg/credentials"
	"github.com/ryand67/teams-directory-go/pkg/firebase"
)

func main() {
	// init "global" context
	ctx := context.Background()

	unFlag := flag.String("username", "", "username")
	pwFlag := flag.String("password", "", "password")
	flag.Parse()

	if !credentials.EmailValid(*unFlag) {
		fmt.Println("Provided command line username invalid email format")
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

}