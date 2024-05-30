package main

import (
	"bufio"
	"cmp"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/fergalbittles/grpc/user"

	"golang.org/x/term"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Attempting to connect to server...")

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(
		":4000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
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

	scanner := bufio.NewScanner(os.Stdin)

	prompt := func(prompt string) string {
		fmt.Print(prompt)
		scanner.Scan()
		return strings.TrimSpace(scanner.Text())
	}

	getpasswd := func() (string, error) {
		fmt.Print("Choose your password: ")
		passwd, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return "", err
		}
		return string(passwd), nil
	}

	// Get the name
	fullName := prompt("Enter your full name: ")

	// Store the name
	var (
		parts     = strings.SplitN(fullName, " ", 2)
		firstName = parts[0]
		lastName  string
	)
	if len(parts) > 1 {
		lastName = parts[1]
	}

	// Get the username
	username := cmp.Or(prompt("Choose your username: "), firstName)
	if username == "" {
		username = firstName
	}

	// Get the password
	password, err := getpasswd()
	if err != nil {
		log.Fatalf("| Error when getting password: %s\n\n", err)
	}

	response, err := u.AddUser(context.Background(), &user.UserCreateRequest{
		User: &user.User{
			FirstName: firstName,
			LastName:  lastName,
			UserName:  username,
			Password:  password,
		},
	})
	if err != nil {
		log.Fatalf("| Error when calling AddUser: %s\n\n", err)
	}

	fmt.Printf("\n%s\n", response.User)
}

func listUsers(u user.UserServiceClient) {
	fmt.Println("\nList Users")
	fmt.Println("++++++++++")

	response, err := u.ListUsers(context.Background(), &user.UserListRequest{})
	if err != nil {
		log.Fatalf("| Error when calling ListUsers: %s\n\n", err)
	}

	for _, user := range response.Users {
		fmt.Println(user)
	}
}
