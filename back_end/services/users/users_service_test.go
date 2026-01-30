package users_service

import (
	"context"
	"testing"
	"time"

	"type_writer_api/helpers"
	"type_writer_api/structures"
	mockProviders "type_writer_api/testing/mocks/providers"

	gomock "go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestGetUsers(t *testing.T) {
	password := "testPassword"
	hashedPassword, _ := helpers.HashPassword(password)
	var (
		mockResult1 = []*structures.User{
			{1, "regular", "testuser1", hashedPassword, "test user1", "tu1@regular", time.Now(), time.Now()},
			{2, "regular", "testuser2", hashedPassword, "test user2", "tu2@regular", time.Now(), time.Now()},
		}
		expectedResult1 = []*structures.UserResp{
			{1, "regular", "testuser1", "test user1", "tu1@regular", time.Now(), time.Now()},
			{2, "regular", "testuser2", "test user2", "tu2@regular", time.Now(), time.Now()},
		}
		mockResult2 = []*structures.User{}
		expectedResult2 = []*structures.UserResp{}
	)
	data := []struct{
		testName string
		mockResult []*structures.User
		mockErr error
		expectedResult []*structures.UserResp
		expectedErr error
	}{
		{
			"success",
			mockResult1,
			nil,
			expectedResult1,
			nil,
		},
		{
			"empty result",
			mockResult2,
			nil,
			expectedResult2,
			nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsersProvider := mockProviders.NewMockUsersProviderInterface(ctrl)
	usersService := NewUsersService(mockUsersProvider)

	for _, testCase := range data {
		t.Run(testCase.testName, func(t *testing.T) {
			mockUsersProvider.EXPECT().GetUsers(context.Background()).Return(testCase.mockResult, testCase.mockErr).Times(1)

			result, err := usersService.GetUsers(context.Background())

			if testCase.expectedErr != nil {
				if err != testCase.expectedErr {
					t.Fatalf("expected error: %v but got %v instead", testCase.expectedErr, err)
				}
			}
			if len(result) != len(testCase.expectedResult) {
				t.Fatalf("slice length missmatch: got %v, expected %v", len(result), len(testCase.expectedResult))
			}

			for idx, user := range result {
				err := helpers.CompareReflectedStructFields(*user, *testCase.expectedResult[idx])
				if err != nil {
					t.Fatalf("row %v failed: %v\n", idx, err.Error())
				}
			}
		})
	}
}

func TestGetUserByIdOrName(t *testing.T) {
	password := "testPassword"
	hashedPassword, _ := helpers.HashPassword(password)
	data := []struct{
		testName string
		inputId int
		inputName string
		mockResult *structures.User
		mockErr error
		expectedResult *structures.UserResp
		expectedErr error
	}{
		{
			"valid id",
			1,
			"",
			&structures.User{1, "regular", "testuser1", hashedPassword, "test user1", "tu1@regular", time.Now(), time.Now()},
			nil,
			&structures.UserResp{1, "regular", "testuser1", "test user1", "tu1@regular", time.Now(), time.Now()},
			nil,
		},
		{
			"valid username",
			0,
			"test user 1",
			&structures.User{1, "regular", "testuser1", hashedPassword, "test user1", "tu1@regular", time.Now(), time.Now()},
			nil,
			&structures.UserResp{1, "regular", "testuser1", "test user1", "tu1@regular", time.Now(), time.Now()},
			nil,
		},
		{
			"not found error",
			0,
			"",
			nil,
			gorm.ErrRecordNotFound,
			nil,
			gorm.ErrRecordNotFound,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsersProvider := mockProviders.NewMockUsersProviderInterface(ctrl)
	usersService := NewUsersService(mockUsersProvider)

	for _, testCase := range data {
		t.Run(testCase.testName, func(t *testing.T) {
			mockUsersProvider.EXPECT().GetUserByIdOrUsername(context.Background(), &testCase.inputId, &testCase.inputName).Return(testCase.mockResult, testCase.mockErr).Times(1)

			result, err := usersService.GetUserByIdOrUsername(context.Background(), &testCase.inputId, &testCase.inputName)

			if testCase.expectedErr != nil {
				if err != testCase.expectedErr {
					t.Fatalf("expected error: %v but got %v instead", testCase.expectedErr, err)
				}
			} else {
				if err := helpers.CompareReflectedStructFields(*result, *testCase.expectedResult); err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	password := "testPassword"
	hashedPassword, _ := helpers.HashPassword(password)
	data := []struct{
		testName string
		inputUser structures.UserReq
		mockResult *structures.User
		mockErr error
		expectedResult *structures.UserResp
		expectedErr error
	}{
		{
			"valid input user request",
			structures.UserReq{"regular", "testuser1", password, "test user1", "tu1@regular"},
			&structures.User{1, "regular", "testuser1", hashedPassword, "test user1", "tu1@regular", time.Now(), time.Now()},
			nil,
			&structures.UserResp{1, "regular", "testuser1", "test user1", "tu1@regular", time.Now(), time.Now()},
			nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsersProvider := mockProviders.NewMockUsersProviderInterface(ctrl)
	usersService := NewUsersService(mockUsersProvider)

	for _, testCase := range data {
		t.Run(testCase.testName, func(t *testing.T) {
			mockUsersProvider.EXPECT().CreateUser(context.Background(), gomock.Any()).Return(testCase.mockResult, testCase.mockErr).Times(1)

			result, err := usersService.CreateUser(context.Background(), testCase.inputUser)

			if testCase.expectedErr != nil {
				if err != testCase.expectedErr {
					t.Fatalf("expected error: %v but got %v instead", testCase.expectedErr, err)
				}
			} else {
				if err := helpers.CompareReflectedStructFields(*result, *testCase.expectedResult); err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	password := "testPassword"
	hashedPassword, _ := helpers.HashPassword(password)
	data := []struct{
		testName string
		inputUser structures.UserReq
		inputUpdateId int
		mockQueryResult *structures.User
		mockResult *structures.User
		mockErr error
		expectedResult *structures.UserResp
		expectedErr error
	}{
		{
			"valid input user request",
			structures.UserReq{"regular", "testuser2", "", "test user2", "tu2@regular"},
			1,
			&structures.User{1, "regular", "testuser1", hashedPassword, "test user1", "tu1@regular", time.Now(), time.Now()},
			&structures.User{1, "regular", "testuser2", hashedPassword, "test user2", "tu2@regular", time.Now(), time.Now()},
			nil,
			&structures.UserResp{1, "regular", "testuser2", "test user2", "tu2@regular", time.Now(), time.Now()},
			nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsersProvider := mockProviders.NewMockUsersProviderInterface(ctrl)
	usersService := NewUsersService(mockUsersProvider)

	for _, testCase := range data {
		t.Run(testCase.testName, func(t *testing.T) {
			mockUsersProvider.EXPECT().GetUserByIdOrUsername(context.Background(), &testCase.inputUpdateId, nil).Return(testCase.mockQueryResult, testCase.mockErr).Times(1)
			mockUsersProvider.EXPECT().UpdateUser(
				context.Background(),
				gomock.Cond(func(input structures.User) bool { return helpers.CompareReflectedStructFields(input, *testCase.mockResult) == nil}),
			).Return(testCase.mockResult, testCase.mockErr).Times(1)

			result, err := usersService.UpdateUser(context.Background(), testCase.inputUser, testCase.inputUpdateId)

			if testCase.expectedErr != nil {
				if err != testCase.expectedErr {
					t.Fatalf("expected error: %v but got %v instead", testCase.expectedErr, err)
				}
			} else {
				if err := helpers.CompareReflectedStructFields(*result, *testCase.expectedResult); err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	data := []struct{
		testName string
		inputDeleteId int
		mockResult bool
		mockErr error
		expectedResult bool
		expectedErr error
	}{
		{
			"valid input user id",
			1,
			true,
			nil,
			true,
			nil,
		},
		{
			"not existant input user id",
			99,
			false,
			nil,
			false,
			nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsersProvider := mockProviders.NewMockUsersProviderInterface(ctrl)
	usersService := NewUsersService(mockUsersProvider)

	for _, testCase := range data {
		t.Run(testCase.testName, func(t *testing.T) {
			mockUsersProvider.EXPECT().DeleteUser(context.Background(), testCase.inputDeleteId).Return(testCase.mockResult, testCase.mockErr).Times(1)

			result, err := usersService.DeleteUser(context.Background(), testCase.inputDeleteId)

			if testCase.expectedErr != nil {
				if err != testCase.expectedErr {
					t.Fatalf("expected error: %v but got %v instead", testCase.expectedErr, err)
				}
			} else {
				if result != testCase.expectedResult {
					t.Fatalf("expected %v, got %v", result, testCase.expectedResult)
				}
			}
		})
	}
}
