package users_service

import (
	"context"
	"log/slog"
	"type_writer_api/helpers"
	"type_writer_api/providers/users"
	"type_writer_api/structures"
)

type UsersServiceInterface interface {
	GetUsers(ctx context.Context) ([]*structures.UserResp, error)
	GetUserByIdOrUsername(ctx context.Context, userId *int, username *string) (*structures.UserResp, error)
	CreateUser(ctx context.Context, userInfo structures.UserReq) (*structures.UserResp, error)
	UpdateUser(ctx context.Context, userInfo structures.UserReq, userId int) (*structures.UserResp, error)
	DeleteUser(ctx context.Context, userId int) (bool, error)
}

type UsersService struct {
	UsersProvider users_provider.UsersProviderInterface
}

func (u *UsersService) GetUsers(ctx context.Context) ([]*structures.UserResp, error) {
	var result []*structures.UserResp

	users, err := u.UsersProvider.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		result = append(result, structures.ConvertUserToResponse(user))
	}

	return result, nil
}

func (u *UsersService) GetUserByIdOrUsername(ctx context.Context, userId *int, username *string) (*structures.UserResp, error) {
	user, err := u.UsersProvider.GetUserByIdOrUsername(ctx, userId, username)
	if err != nil {
		return nil, err
	}

	result := structures.ConvertUserToResponse(user)
	return result, nil
}

func (u *UsersService) CreateUser(ctx context.Context, userInfo structures.UserReq) (*structures.UserResp, error) {
	userToCreate := structures.ConvertRequestToUser(&userInfo)

	hashedPassword, err := helpers.HashPassword(userInfo.Password)
	if err != nil {
		slog.ErrorContext(ctx, "failed password hashing", "error", err)
		return nil, err
	}
	userToCreate.PasswdHash = hashedPassword

	createdUser, err := u.UsersProvider.CreateUser(ctx, *userToCreate)
	if err != nil {
		slog.ErrorContext(ctx, "failed to create user", "error", err)
		return nil, err
	}

	result := structures.ConvertUserToResponse(createdUser)
	return result, nil
}

func (u *UsersService) UpdateUser(ctx context.Context, userInfo structures.UserReq, userId int) (*structures.UserResp, error) {
	existingUser, err := u.UsersProvider.GetUserByIdOrUsername(ctx, &userId, nil)

	if err != nil {
		slog.ErrorContext(ctx, "failed to update user", "error", err)
		return nil, err
	}

	if userInfo.UserType != "" {
		existingUser.UserType = userInfo.UserType
	}
	if userInfo.Username != "" {
		existingUser.Username = userInfo.Username
	}
	if userInfo.Name != "" {
		existingUser.Name = userInfo.Name
	}
	if userInfo.Email != "" {
		existingUser.Email = userInfo.Email
	}
	if userInfo.Password != "" {
		hashedPassword, err := helpers.HashPassword(userInfo.Password)
		if err != nil {
			slog.ErrorContext(ctx, "failed password hashing", "error", err)
			return nil, err
		}
		existingUser.PasswdHash = hashedPassword
	}

	updatedUser, err := u.UsersProvider.UpdateUser(ctx, *existingUser)
	if err != nil {
		slog.ErrorContext(ctx, "failed to update user", "error", err)
		return nil, err
	}

	result := structures.ConvertUserToResponse(updatedUser)
	return result, nil
}

func (u *UsersService) DeleteUser(ctx context.Context, userId int) (bool, error) {
	deleted, err := u.UsersProvider.DeleteUser(ctx, userId)
	if err != nil {
		return false, err
	}

	return deleted, nil
}

func NewUsersService(usersProvider users_provider.UsersProviderInterface) *UsersService {
	return &UsersService{
		UsersProvider: usersProvider,
	}
}
