package credentials

import (
	"context"
	"fmt"
	"log"
	"net/mail"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/iterator"
)

type User struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx context.Context, app *firebase.App, e string, p string) (context.Context, error) {
	client, err := app.Firestore(ctx)
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	defer client.Close()

	var userDoc map[string]interface{}
	iter := client.Collection("users").Where("Email", "==", e).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if doc != nil {
			ctx = context.WithValue(ctx, "docId", doc.Ref.ID)
			userDoc = doc.Data()
		}
	}

	if userDoc["Password"] != p {
		fmt.Println("Wrong username and/or password")
		panic("Wrong username and/or password")
	}

	ctx = context.WithValue(ctx, "user", userDoc["Email"])
	return ctx, nil
}

func SignUp(ctx context.Context, e string, p string, app *firebase.App) (context.Context, error) {
	client, err := app.Firestore(ctx)
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	defer client.Close()

	iter := client.Collection("users").Where("Email", "==", e).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if doc != nil {
			fmt.Println("User exists with that email")
			panic("User exists with that email")
		}
	}
	
	_, _, err = client.Collection("users").Add(ctx, &User{
		Email: e,
		Password: p,
	})
	if err != nil {
		log.Fatalf("Failed adding user: %v", err)
	}

	log.Printf("Successfully created user: %v\n", e)
	ctx = context.WithValue(ctx, "user", e)

	return ctx, nil
}

func PromptLogin(app *firebase.App, ctx context.Context) (context.Context, error) {
	// catches panics
	defer func() {
		if r := recover(); r != nil {
			PromptLogin(app, ctx)
		}
	}()

	fmt.Println("Sign-up, Login, or Exit? (S/L/Exit)")

	// Sign up or login
	var sOrL string
	fmt.Scanln(&sOrL)

	if strings.ToUpper(sOrL) == "EXIT" || strings.ToUpper(sOrL) == "E" {
		os.Exit(3)
	}

	var e string
	fmt.Println("What is your email: ")
	fmt.Scan(&e)
	e = strings.TrimSpace(e)

	if valid := EmailValid(e); !valid {
		log.Println("Need a valid email.")
		PromptLogin(app, ctx)
	}
	
	var p string
	fmt.Println("What is your password: ")
	fmt.Scan(&p)

	if strings.ToUpper(sOrL) == "S" {
		c, err := SignUp(ctx, e, p, app)
		ctx = c
		if err != nil {
			return c, err
		}
	} else if strings.ToUpper(sOrL) == "L" {
		c, err := Login(ctx, app, e, p)
		ctx = c
		if err != nil {
			return c, err
		}
	} else {
		log.Println("Need a valid response (S/L/Exit)")
		PromptLogin(app, ctx)
	}

	return ctx, nil
}

func EmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}