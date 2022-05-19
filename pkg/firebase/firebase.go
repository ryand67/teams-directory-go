package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
)

func NewApp() *firebase.App {
	app, err := firebase.NewApp(context.Background(), nil)	
	if err != nil {
		fmt.Printf("ERR: %s", err)
	}

	return app;
}