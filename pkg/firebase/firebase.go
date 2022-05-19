package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func NewApp() *firebase.App {
	app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsFile("../../auth.json"))	
	if err != nil {
		fmt.Printf("ERR: %s", err)
	}

	return app;
}