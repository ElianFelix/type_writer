package services

import (
	"context"
	"type_writer_api/providers"
	"type_writer_api/structures"
)

type UserServiceInterface interface {
	GetUsers(ctx context.Context) ([]*structures.UserResp, error)
	GetUserByIdOrUsername(ctx context.Context, userId int, username string) (*structures.UserResp, error)
	CreateUser(ctx context.Context, userInfo structures.UserReq) (*structures.UserResp, error)
	UpdateUser(ctx context.Context, userInfo structures.UserReq) (*structures.UserResp, error)
	DeleteUser(ctx context.Context, userId int) (bool, error)
}

type UserService struct {
	UserProvider *providers.UserProvider
}

func (u *UserService) GetUsers(ctx context.Context) ([]*structures.UserResp, error) {
	var usersResp []*structures.UserResp

	users, err  := u.UserProvider.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		sanitizedUser := structures.UserResp{
			Id: user.Id,
			UserType: user.UserType,
			Username: user.Username,
			Name: user.Name,
			Email: user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		usersResp = append(usersResp, &sanitizedUser)
	}

	return usersResp, nil
}

func NewUserService(userProvider *providers.UserProvider) *UserService {
	return &UserService{
		UserProvider: userProvider,
	}
}
