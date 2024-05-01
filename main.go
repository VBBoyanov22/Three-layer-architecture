package main

// Import the necessary packages
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define the file where user data is stored
const userDataFile = "user_data.txt"

func main() {
	// Loop to display menu options
	for {
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			login()
		case 2:
			register()
		case 3:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please choose again.")
		}
	}
}

// Function to handle user login
func login() {
	fmt.Println("Login")
	fmt.Print("Enter username: ")
	username := getUserInput() // Get username input
	fmt.Print("Enter password: ")
	password := getUserInput() // Get password input

	// Check if user exists in user_data.txt
	file, err := os.Open(userDataFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		if data[0] == username && data[1] == password {
			fmt.Println("Login successful!")
			return
		}
	}

	fmt.Println("Invalid username or password.")
}

// Function to handle user registration
func register() {
	fmt.Println("Register")
	fmt.Print("Enter username: ")
	username := getUserInput() // Get username input
	fmt.Print("Enter password: ")
	password := getUserInput() // Get password input

	// Open the user data file to append new user information
	file, err := os.OpenFile(userDataFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Write the new user information to user_data.txt
	_, err = fmt.Fprintf(file, "%s %s\n", username, password)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Registration successful!")
}

// Function to get user input from stdin
func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
