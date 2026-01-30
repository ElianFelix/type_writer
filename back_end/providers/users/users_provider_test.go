package users_provider

import (
	"context"
	"testing"
	"time"
	"type_writer_api/helpers"
	"type_writer_api/structures"
	"type_writer_api/testing/mocks"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetUsersSuccess(t *testing.T) {
	mockGorm, mockDB := mocks.NewMockDB()
	usersProvider := NewUsersProvider(mockGorm)

	expectedRows := []structures.User{
		{
			Id:         1,
			UserType:   "regular",
			Username:   "testUser1",
			PasswdHash: "hashedpassword",
			Name:       "testivo",
			Email:      "test1@user.com",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			Id:         2,
			UserType:   "regular",
			Username:   "testUser2",
			PasswdHash: "hashedpassword",
			Name:       "testivo",
			Email:      "test2@user.com",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	resultRows := sqlmock.NewRows([]string{
		"id",
		"user_type",
		"username",
		"passwd_hash",
		"name",
		"email",
		"created_at",
		"updated_at",
	})

	for _, expectedRow := range expectedRows {
		resultRows.AddRow(
			expectedRow.Id,
			expectedRow.UserType,
			expectedRow.Username,
			expectedRow.PasswdHash,
			expectedRow.Name,
			expectedRow.Email,
			expectedRow.CreatedAt,
			expectedRow.UpdatedAt,
		)
	}

	mockDB.ExpectQuery(`SELECT \* FROM "users"`).WillReturnRows(resultRows)

	result, err := usersProvider.GetUsers(context.Background())

	if err != nil {
		t.Fatalf("error in fetching users %v", err)
	}

	if len(result) != 2 {
		t.Fatalf("unexpected result length: expected %v, got %v", 2, len(result))
	}

	for indx, resultRow := range result {
		err := helpers.CompareReflectedStructFields(*resultRow, expectedRows[indx])
		if err != nil {
			t.Fatalf("row %v failed: %v\n", indx, err.Error())
		}
	}
}

func TestGetUserByIdOrNameSuccess(t *testing.T) {
	mockGorm, mockDB := mocks.NewMockDB()
	usersProvider := NewUsersProvider(mockGorm)

	expectedRow := structures.User{
		Id:         1,
		UserType:   "regular",
		Username:   "testUser1",
		PasswdHash: "hashedpassword",
		Name:       "testivo",
		Email:      "test1@user.com",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	resultRows := sqlmock.NewRows([]string{
		"id",
		"user_type",
		"username",
		"passwd_hash",
		"name",
		"email",
		"created_at",
		"updated_at",
	}).AddRow(
		expectedRow.Id,
		expectedRow.UserType,
		expectedRow.Username,
		expectedRow.PasswdHash,
		expectedRow.Name,
		expectedRow.Email,
		expectedRow.CreatedAt,
		expectedRow.UpdatedAt,
	)

	mockDB.ExpectQuery(`SELECT \* FROM "users" WHERE id = .+ OR username = .+ ORDER BY "users"\."id" LIMIT .+`).WillReturnRows(resultRows)

	inputId := 1
	inputUsername := ""
	result, err := usersProvider.GetUserByIdOrUsername(context.Background(), &inputId, &inputUsername)

	if err != nil {
		t.Fatalf("error in fetching user %v", err)
	}

	if err := helpers.CompareReflectedStructFields(*result, expectedRow); err != nil {
		t.Fatal(err)
	}
}

func TestCreateUserSuccess(t *testing.T) {
	mockGorm, mockDB := mocks.NewMockDB()
	usersProvider := NewUsersProvider(mockGorm)

	expectedRow := structures.User{
		Id:         1,
		UserType:   "regular",
		Username:   "testUser1",
		PasswdHash: "hashedpassword",
		Name:       "testivo",
		Email:      "test1@user.com",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	resultRows := sqlmock.NewRows([]string{
		"id",
		"user_type",
		"username",
		"passwd_hash",
		"name",
		"email",
		"created_at",
		"updated_at",
	}).AddRow(
		expectedRow.Id,
		expectedRow.UserType,
		expectedRow.Username,
		expectedRow.PasswdHash,
		expectedRow.Name,
		expectedRow.Email,
		expectedRow.CreatedAt,
		expectedRow.UpdatedAt,
	)

	mockDB.ExpectQuery(`SELECT \* FROM "users" WHERE "users"\."id" = .+ AND "users"\."user_type" = .+ AND "users"\."username" = .+ AND "users"\."passwd_hash" = .+ AND "users"\."name" = .+ AND "users"\."email" = .+ AND "users"\."created_at" = .+ AND "users"\."updated_at" = .+ ORDER BY "users"\."id" LIMIT .+`).WillReturnRows(resultRows)

	result, err := usersProvider.CreateUser(context.Background(), expectedRow)

	if err != nil {
		t.Fatalf("error in creating user %v", err)
	}

	if err := helpers.CompareReflectedStructFields(*result, expectedRow); err != nil {
		t.Fatal(err)
	}
}

func TestUpdateUserSuccess(t *testing.T) {
	mockGorm, mockDB := mocks.NewMockDB()
	usersProvider := NewUsersProvider(mockGorm)

	expectedRow := structures.User{
		Id:         1,
		UserType:   "regular",
		Username:   "testUser1",
		PasswdHash: "hashedpassword",
		Name:       "testivo",
		Email:      "test1@user.com",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	mockDB.ExpectBegin()
	mockDB.ExpectExec(`UPDATE "users" SET "user_type"=.+,"username"=.+,"passwd_hash"=.+,"name"=.+,"email"=.+,"created_at"=.+,"updated_at"=.+ WHERE "id" = .+`).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	result, err := usersProvider.UpdateUser(context.Background(), expectedRow)

	if err != nil {
		t.Fatalf("error in updating user %v", err)
	}

	if err := helpers.CompareReflectedStructFields(*result, expectedRow); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteUserSuccess(t *testing.T) {
	mockGorm, mockDB := mocks.NewMockDB()
	usersProvider := NewUsersProvider(mockGorm)

	expectedRow := structures.User{
		Id: 1,
	}

	mockDB.ExpectBegin()
	mockDB.ExpectExec(`FROM "users" WHERE "users"\."id" = .+`).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	result, err := usersProvider.DeleteUser(context.Background(), expectedRow.Id)

	if err != nil {
		t.Fatalf("error in deleting user %v", err)
	}

	if result != true {
		t.Fatalf("unexpected result: expected %v, got %v", true, result)
	}
}
