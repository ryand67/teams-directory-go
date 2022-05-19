package credentials

import (
	"fmt"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
)

func Login() error {
	return nil
}

func SignUp() error {
	return nil
}

func PromptLogin(app *firebase.App) error {
	fmt.Println("Sign-up, Login, or Exit? (S/L/Exit)")
	// Sign up or login
	var sOrL string

	fmt.Scanln(&sOrL)

	if strings.ToUpper(sOrL) == "S" {
		SignUp()
	} else if strings.ToUpper(sOrL) == "L" {
		Login()
	} else if strings.ToUpper(sOrL) == "EXIT" {
		os.Exit(3)
	} else {
		fmt.Println("Need a valid response (S/L)")
		PromptLogin(app)
	}

	return nil
}