package main

import (
	"bufio"
	"context"
	"fmt"
	"grpc/user"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Attempting to connect to server...")

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":4000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("| Could not connect: %s", err)
	}
	defer conn.Close()

	u := user.NewUserServiceClient(conn)

	exit := false
	for {
		if exit {
			fmt.Println("\nExiting System - Goodbye")
			break
		}

		input := userMenu()

		switch input {
		case "1":
			addUser(u)
		case "2":
			listUsers(u)
		case "3":
			exit = true
		default:
			fmt.Println("\nError: Invalid Response")
		}
	}
}

func userMenu() string {
	fmt.Println("\nUser System")
	fmt.Println("+++++++++++")
	fmt.Println("1. Add a User")
	fmt.Println("2. List All Users")
	fmt.Println("3. Exit")
	fmt.Print("\nEnter option number: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	return input
}

func addUser(u user.UserServiceClient) {
	fmt.Println("\nAdd User")
	fmt.Println("++++++++")

	// Get the name
	fmt.Print("Enter your full name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input1 := strings.TrimSpace(scanner.Text())

	// Store the name
	var firstName string
	var lastName string
	fullName := strings.Split(input1, " ")

	if len(fullName) == 1 {
		if fullName[0] == "" {
			fmt.Println("\nFailure to add user: Must enter a name")
			return
		}
		firstName = fullName[0]
		lastName = "N/A"
	} else if len(fullName) > 1 {
		firstName = fullName[0]
		lastName = strings.Join(fullName[1:], " ")
	}

	// Get the username
	fmt.Print("Choose your username: ")
	scanner.Scan()
	username := strings.TrimSpace(scanner.Text())
	if username == "" {
		username = firstName
	}

	// Get the password
	fmt.Print("Choose your password: ")
	scanner.Scan()
	password := strings.TrimSpace(scanner.Text())
	if password == "" {
		password = "password"
	}

	// Send the request
	req := firstName + "|" + lastName + "|" + username + "|" + password
	if strings.Count(req, "|") > 3 {
		fmt.Println("\nFailure to add user: Cannot use \"|\" character")
		return
	}
	request := user.UserRequest{
		Body: req,
	}

	response, err := u.AddUser(context.Background(), &request)
	if err != nil {
		log.Fatalf("| Error when calling AddUser: %s\n\n", err)
	}

	fmt.Printf("\n%s\n", response.Body)
}

func listUsers(u user.UserServiceClient) {
	fmt.Println("\nList Users")
	fmt.Println("++++++++++")

	request := user.UserRequest{
		Body: "List all of the users",
	}

	response, err := u.ListUsers(context.Background(), &request)
	if err != nil {
		log.Fatalf("| Error when calling ListUsers: %s\n\n", err)
	}

	if response.Body == "" {
		fmt.Println("There are no users stored on the system")
		return
	}

	fmt.Print(response.Body)
}
