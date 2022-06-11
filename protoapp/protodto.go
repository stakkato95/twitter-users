package protoapp

import (
	"github.com/stakkato95/twitter-service-users/domain"
	pb "github.com/stakkato95/twitter-service-users/proto"
)

func ToEntity(u *pb.User) domain.User {
	return domain.User{
		Id:       u.Id,
		Username: u.Username,
		Password: u.Password,
	}
}

func NewUserToDto(u *domain.User, token string) *pb.NewUser {
	return &pb.NewUser{
		User: UserToDto(u),
		Token: &pb.Token{
			Token: token,
		},
	}
}

func UserToDto(u *domain.User) *pb.User {
	return &pb.User{
		Id:       u.Id,
		Username: u.Username,
		Password: u.Password,
	}
}
