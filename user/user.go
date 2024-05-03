package user

import (
	"grpc/globals"
	"log"
	"strconv"
	"strings"

	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) AddUser(ctx context.Context, request *UserRequest) (*UserResponse, error) {
	log.Printf("| Received request from user: %s", request.Body)

	userDetails := strings.Split(request.Body, "|")
	newUser := globals.User{
		ID:        globals.UseNextID(),
		FirstName: userDetails[0],
		LastName:  userDetails[1],
		UserName:  userDetails[2],
		Password:  userDetails[3],
	}
	globals.AllUsers = append(globals.AllUsers, newUser)

	if newUser.LastName == "N/A" {
		res := "Confirmed: \"" + newUser.FirstName + "\" has been added to the system"
		return &UserResponse{Body: res}, nil
	}

	res := "Confirmed: \"" + newUser.FirstName + " " + newUser.LastName + "\" has been added to the system"
	return &UserResponse{Body: res}, nil
}

func (s *Server) ListUsers(ctx context.Context, request *UserRequest) (*UserResponse, error) {
	log.Printf("| Received request from user: %s", request.Body)

	var res string
	for _, user := range globals.AllUsers {
		res += "\nID: " + strconv.Itoa(user.ID) + "\n"
		res += "First Name: " + user.FirstName + "\n"
		res += "Last Name: " + user.LastName + "\n"
		res += "Username: " + user.UserName + "\n"
		res += "Password: " + user.Password[0:1] + "****" + "\n"
	}

	return &UserResponse{Body: res}, nil
}
