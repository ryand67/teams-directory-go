package team

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
)

type User struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Teams []Team `json:"teams"`
}

type Team struct {
	TeamName string `json:"teamName"`
	Members []Member `json:"members"`
}

type Member struct {
	Name string `json:"memberName"`
	Role string `json:"memberRole"`
}

type UserDoc struct {
	Email string
	Password string
	teams map[string]interface{}
}

func GetDoc(ctx context.Context, app *firebase.App) map[string]interface{} {
	ID := ctx.Value("docId")

	client, err := app.Firestore(ctx)
	if err != nil {
		fmt.Println("Error instantiating database: ", err)
	}

	defer client.Close()

	dsnap, err := client.Collection("users").Doc(ID.(string)).Get(ctx);
	if err != nil {
		fmt.Println("Error retrieving document:", err)
	}

	doc := dsnap.Data()
	return doc
}

func TeamList(ctx context.Context, app *firebase.App) {
	doc := GetDoc(ctx, app)

	teamList, ok := doc["Teams"]

	if ok == true {
		for k := range teamList.(map[string]interface{}) {
			fmt.Println(k)
		}
	} else {
		fmt.Println("No teams to list.")
	}
}

func AddTeam(ctx context.Context, app *firebase.App) error {
	ID := ctx.Value("docId")
	
	var newName string
	
	fmt.Println("What's the name of your new team?")
	fmt.Scanln(&newName)

	re := regexp.MustCompile("^[a-zA-Z0-9]*$")
	validName := re.MatchString(newName)

	if validName == false {
		return errors.New("Team name must be alphanumeric!")
	}

	// Grab doc to append values
	doc := GetDoc(ctx, app)
	teamList, ok := doc["Teams"]

	for k := range teamList.(map[string]interface{}) {
		if strings.ToUpper(k) == strings.ToUpper(newName) {
			return errors.New("Team name already exists!")
		}
	}

	if ok == true {
		teamList = append(teamList.([]interface{}), &Team{
			TeamName: newName,
		})
	} else {
		teamList = &Team{
			TeamName: newName,
		}
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		fmt.Println("Error connecting to service.", err)
		return err
	}

	defer client.Close()

	_, err = client.Collection("users").Doc(ID.(string)).Update(ctx, []firestore.Update{
		{
			Path: fmt.Sprint("Teams.", newName),
			Value: teamList,
		},
	})

	return nil
}