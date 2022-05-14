package credentials

import "fmt"

func Login() error {
	return nil
}

func SignUp() error {
	return nil
}

func PromptLogin() {
	fmt.Println("Sign-up or Login? (S/L)")
	var sOrL string

	fmt.Scanln(&sOrL)

	if sOrL == "S" || sOrL == "s" {
		SignUp()
	} else if sOrL == "L" || sOrL == "l" {
		Login()
	} else {
		fmt.Println("Need a valid response (S/L)")
		PromptLogin()
	}
}