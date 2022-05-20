package team

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
)

type UserDoc struct {
	Email string
	Password string
	teams map[string]interface{}
}

func TeamList(ctx context.Context, app *firebase.App) error {
	ID := ctx.Value("docId")

	client, err := app.Firestore(ctx)
	if err != nil {
		fmt.Println("Error instantiating database: ", err)
	}

	defer client.Close()

	dsnap, err := client.Collection("users").Doc(ID.(string)).Get(ctx);
	if err != nil {
		fmt.Println("Error retrieving list:", err)
	}

	doc := dsnap.Data()
	teamList := doc["teams"]

	for k := range teamList.(map[string]interface{}) {
		fmt.Println(k)
	}

	return nil
}

func AddTeam(ctx context.Context, app *firebase.App) error {
	return nil
}