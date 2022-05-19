package credentials

import (
	"context"
	"fmt"
	"log"
	"net/mail"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

func Login(app *firebase.App, e string, p string) error {
	// ctx := context.Background()
	return nil
}

func SignUp(ctx context.Context, e string, p string, app *firebase.App) (context.Context, error) {
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("Error creating auth util: %v\n", err)
		return nil, err
	}
	
	params := (&auth.UserToCreate{}).Email(e).Password(p)

	u, err := client.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("Error creating user: %v\n", err)
		return nil, err
	}

	log.Printf("Successfully created user: %v\n%v", e, u)
	ctx = context.WithValue(ctx, "user", e)

	return ctx, nil
}

func PromptLogin(app *firebase.App, ctx context.Context) (context.Context, error) {
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

	if valid := emailValid(e); !valid {
		log.Println("Need a valid email!")
		PromptLogin(app, ctx)
	}
	
	var p string
	fmt.Println("What is your password: ")
	fmt.Scan(&p)

	if strings.ToUpper(sOrL) == "S" {
		c, err := SignUp(ctx, e, p, app)
		if err != nil {
			return c, err
		}
	} else if strings.ToUpper(sOrL) == "L" {
		Login(app, e, p)
	} else {
		log.Println("Need a valid response (S/L/Exit)")
		PromptLogin(app, ctx)
	}

	return ctx, nil
}

func emailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}