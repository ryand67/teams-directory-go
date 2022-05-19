package main

import (
	"github.com/ryand67/teams-directory-go/pkg/credentials"
	"github.com/ryand67/teams-directory-go/pkg/firebase"
)

func main() {
	// Creates new firebase app instance
	app := firebase.NewApp()

	credentials.PromptLogin(app)
}