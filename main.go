package main

import (
	"context"

	"github.com/ryand67/teams-directory-go/pkg/credentials"
	"github.com/ryand67/teams-directory-go/pkg/firebase"
)

func main() {
	ctx := context.Background()
	// Creates new firebase app instance
	app := firebase.NewApp(ctx)

	credentials.PromptLogin(app, ctx)
}