package providers

import (
	"context"
	"type_writer_api/structures"

	"gorm.io/gorm"
)

type UserProviderInterface interface {
	GetUsers(ctx context.Context) ([]*structures.User, error)
	GetUserByIdOrUsername(ctx context.Context, userId int, username string) (*structures.User, error)
	CreateUser(ctx context.Context, userInfo structures.User) (*structures.User, error)
	UpdateUser(ctx context.Context, updatedUserInfo structures.User) (*structures.User, error)
	DeleteUser(ctx context.Context, userId int) (bool, error)
}

type UserProvider struct {
	Db *gorm.DB
}

func (u *UserProvider) GetUsers(ctx context.Context) ([]*structures.User, error) {
	var users []*structures.User
	err := u.Db.WithContext(ctx).Table(structures.USER_TABLE_NAME).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserProvider) GetUserByIdOrUsername(ctx context.Context, userId int, username string) (*structures.User, error) {
	var user *structures.User
	err := u.Db.WithContext(ctx).Table(structures.USER_TABLE_NAME).
		First(&user, "id = ? OR username = ?", userId, username).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserProvider) CreateUser(ctx context.Context, userInfo structures.User) (*structures.User, error) {
	var user *structures.User
	err := u.Db.WithContext(ctx).Table(structures.USER_TABLE_NAME).FirstOrCreate(&user, &userInfo).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserProvider) UpdateUser(ctx context.Context, updatedUserInfo structures.User) (*structures.User, error) {
	err := u.Db.WithContext(ctx).Table(structures.USER_TABLE_NAME).Updates(&updatedUserInfo).Error
	if err != nil {
		return nil, err
	}
	return &updatedUserInfo, nil
}

func (u *UserProvider) DeleteUser(ctx context.Context, userId int) (bool, error) {
	var deleteUser = structures.User{Id: userId}
	err := u.Db.WithContext(ctx).Table(structures.USER_TABLE_NAME).Delete(&deleteUser).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewUserProvider(db *gorm.DB) *UserProvider {
	return &UserProvider{
		Db: db,
	}
}
