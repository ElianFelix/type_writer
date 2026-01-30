package users_provider

import (
	"context"
	"type_writer_api/structures"

	"gorm.io/gorm"
)

type UsersProviderInterface interface {
	GetUsers(ctx context.Context) ([]*structures.User, error)
	GetUserByIdOrUsername(ctx context.Context, userId *int, username *string) (*structures.User, error)
	CreateUser(ctx context.Context, userInfo structures.User) (*structures.User, error)
	UpdateUser(ctx context.Context, updatedUserInfo structures.User) (*structures.User, error)
	DeleteUser(ctx context.Context, userId int) (bool, error)
}

type UsersProvider struct {
	Db *gorm.DB
}

func (u *UsersProvider) GetUsers(ctx context.Context) ([]*structures.User, error) {
	var users []*structures.User
	err := u.Db.WithContext(ctx).Table(structures.USER_TABLE_NAME).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UsersProvider) GetUserByIdOrUsername(ctx context.Context, userId *int, username *string) (*structures.User, error) {
	var user *structures.User
	err := u.Db.WithContext(ctx).Table(structures.USER_TABLE_NAME).
		First(&user, "id = ? OR username = ?", userId, username).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UsersProvider) CreateUser(ctx context.Context, userInfo structures.User) (*structures.User, error) {
	var user *structures.User
	err := u.Db.WithContext(ctx).Table(structures.USER_TABLE_NAME).FirstOrCreate(&user, &userInfo).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UsersProvider) UpdateUser(ctx context.Context, updatedUserInfo structures.User) (*structures.User, error) {
	err := u.Db.WithContext(ctx).Table(structures.USER_TABLE_NAME).Updates(&updatedUserInfo).Error
	if err != nil {
		return nil, err
	}
	return &updatedUserInfo, nil
}

func (u *UsersProvider) DeleteUser(ctx context.Context, userId int) (bool, error) {
	var deleteUser = structures.User{Id: userId}
	err := u.Db.WithContext(ctx).Table(structures.USER_TABLE_NAME).Delete(&deleteUser).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewUsersProvider(db *gorm.DB) *UsersProvider {
	return &UsersProvider{
		Db: db,
	}
}
