package main

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

func main() {
	var username string
	var email string
	var password string
	var repPassword string
	fmt.Print("Enter your username: ")
	fmt.Fscan(os.Stdin, &username)

	fmt.Print("Enter your email: ")
	fmt.Fscan(os.Stdin, &email)

	var checkedEmail bool = checkEmail(email)

	for !checkedEmail {
		fmt.Print("Your email is incorrect!\n")
		fmt.Print("Enter your email: ")
		fmt.Fscan(os.Stdin, &email)
		checkedEmail = checkEmail(email)
	}

	fmt.Printf("√ '%s' is a valid email\n", email)

	fmt.Print("Enter your password: ")
	fmt.Fscan(os.Stdin, &password)

	fmt.Print("Repeat your password: ")
	fmt.Fscan(os.Stdin, &repPassword)

	var checkedPassword = checkPassword(password, repPassword)

	for !checkedPassword {
		fmt.Print("Passwords are not equal!\n")
		fmt.Print("Repeat your password: ")
		fmt.Fscan(os.Stdin, &repPassword)
		checkedPassword = checkPassword(password, repPassword)
	}

	time.Sleep(1 * time.Second)

	fmt.Print("\n\n\nCatching your data")
	time.Sleep(300 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(300 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(300 * time.Millisecond)
	fmt.Print(".")

	time.Sleep(5 * time.Second)

	fmt.Println("\n\n\nYour username is: " + username)
	fmt.Println("Your email is: " + email)
	fmt.Println("Your password is: " + password)

	fmt.Print("\nNow we have your data. Wait for your naked photos in the Internet)")
}

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func checkEmail(email string) bool {
	pattern := `^\w+@\w+\.\w+$`
	matched, err := regexp.Match(pattern, []byte(email))
	check(err)
	return matched
}

func checkPassword(password1 string, password2 string) bool {
	if password1 == password2 {
		return true
	} else {
		return false
	}
}
