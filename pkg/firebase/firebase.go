package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func NewApp(ctx context.Context) *firebase.App {
	// Create new firebase app instance and return it
	app, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: "team-directory-c8996"}, option.WithCredentialsFile("auth.json"))	
	if err != nil {
		fmt.Printf("ERR: %s", err)
	}

	return app
}