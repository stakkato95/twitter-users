package protoapp

import (
	"context"
	"errors"

	pb "github.com/stakkato95/twitter-service-users/proto"
	"github.com/stakkato95/twitter-service-users/service"
)

type defaultUsersServiceServer struct {
	pb.UnimplementedUsersServiceServer
	service service.UserService
}

func NewUsersServiceServer(service service.UserService) pb.UsersServiceServer {
	return &defaultUsersServiceServer{service: service}
}

func (s *defaultUsersServiceServer) CreateUser(ctx context.Context, user *pb.User) (*pb.NewUser, error) {
	entity := ToEntity(user)
	token, createdUser, err := s.service.Create(&entity)
	if err != nil {
		return nil, errors.New(err.Msg)
	}
	return NewUserToDto(createdUser, token), nil
}

func (s *defaultUsersServiceServer) AuthUser(ctx context.Context, user *pb.User) (*pb.Token, error) {
	entity := ToEntity(user)
	token, err := s.service.Authenticate(&entity)
	if err != nil {
		return nil, errors.New(err.Msg)
	}
	return &pb.Token{Token: token}, nil
}

func (s *defaultUsersServiceServer) AuthUserByToken(ctx context.Context, token *pb.Token) (*pb.User, error) {
	user, err := s.service.Authorize(token.Token)
	if err != nil {
		return nil, errors.New(err.Msg)
	}
	return UserToDto(user), nil
}
