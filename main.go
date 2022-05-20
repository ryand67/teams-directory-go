package main

import (
	"context"
	"log"
	"os"

	"github.com/ryand67/teams-directory-go/pkg/credentials"
	"github.com/ryand67/teams-directory-go/pkg/firebase"
)

func main() {
	// init "global" context
	ctx := context.Background()

	// Creates new firebase app instance
	app := firebase.NewApp(ctx)

	// Login or sign up, returns updated context
	c, err := credentials.PromptLogin(app, ctx)
	if err != nil {
		log.Fatalf(err.Error())
		os.Exit(1)
	}
	ctx = c
}