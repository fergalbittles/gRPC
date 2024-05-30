package user

import (
	"log"

	"github.com/fergalbittles/grpc/globals"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedUserServiceServer
}

func (s *Server) AddUser(ctx context.Context, request *UserCreateRequest) (*UserCreateResponse, error) {
	log.Printf("| Received request from user: %s", request.User)

	newUser := globals.User{
		FirstName: request.User.FirstName,
		LastName:  request.User.LastName,
		UserName:  request.User.UserName,
		Password:  request.User.Password,
	}
	globals.AppendUsers(newUser)

	return &UserCreateResponse{User: request.User}, nil
}

func (s *Server) ListUsers(ctx context.Context, request *UserListRequest) (*UserListResponse, error) {
	users := globals.ListUsers()

	uu := make([]*User, len(users))
	for i, user := range users {
		uu[i] = &User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			UserName:  user.UserName,
			Password:  user.Password,
		}
	}

	return &UserListResponse{Users: uu}, nil
}
