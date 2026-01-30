package activities_provider

import (
	"context"
	"testing"
	"time"
	"type_writer_api/structures"
	"type_writer_api/testing/mocks"
	"type_writer_api/helpers"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetActivitiesSuccess(t *testing.T) {
	mockGorm, mockDB := mocks.NewMockDB()
	activitiesProvider := NewActivitiesProvider(mockGorm)

	expectedRows := []structures.Activity{
		{
			Id:          1,
			Name:        "speed drill",
			Description: "test speed drill",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Id:          2,
			Name:        "article",
			Description: "test article",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	resultRows := sqlmock.NewRows([]string{
		"id",
		"name",
		"description",
		"created_at",
		"updated_at",
	})

	for _, expectedRow := range expectedRows {
		resultRows.AddRow(
			expectedRow.Id,
			expectedRow.Name,
			expectedRow.Description,
			expectedRow.CreatedAt,
			expectedRow.UpdatedAt,
		)
	}

	mockDB.ExpectQuery(`SELECT \* FROM "activities"`).WillReturnRows(resultRows)

	result, err := activitiesProvider.GetActivities(context.Background())

	if err != nil {
		t.Fatalf("error in fetching activities %v", err)
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

func TestGetActivityByIdOrNameSuccess(t *testing.T) {
	mockGorm, mockDB := mocks.NewMockDB()
	activitiesProvider := NewActivitiesProvider(mockGorm)

	expectedRow := structures.Activity{
		Id:          1,
		Name:        "speed drill",
		Description: "test speed drill",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	resultRows := sqlmock.NewRows([]string{
		"id",
		"name",
		"description",
		"created_at",
		"updated_at",
	}).AddRow(
		expectedRow.Id,
		expectedRow.Name,
		expectedRow.Description,
		expectedRow.CreatedAt,
		expectedRow.UpdatedAt,
	)

	mockDB.ExpectQuery(`SELECT \* FROM "activities" WHERE id = .+ OR name = .+ ORDER BY "activities"\."id" LIMIT .+`).WillReturnRows(resultRows)

	inputId := 1
	inputName := ""
	result, err := activitiesProvider.GetActivityByIdOrName(context.Background(), &inputId, &inputName)

	if err != nil {
		t.Fatalf("error in fetching activity %v", err)
	}

	if err := helpers.CompareReflectedStructFields(*result, expectedRow); err != nil {
		t.Fatal(err)
	}
}

func TestCreateActivitySuccess(t *testing.T) {
	mockGorm, mockDB := mocks.NewMockDB()
	activitiesProvider := NewActivitiesProvider(mockGorm)

	expectedRow := structures.Activity{
		Id:          1,
		Name:        "speed drill",
		Description: "test speed drill",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	resultRows := sqlmock.NewRows([]string{
		"id",
		"name",
		"description",
		"created_at",
		"updated_at",
	}).AddRow(
		expectedRow.Id,
		expectedRow.Name,
		expectedRow.Description,
		expectedRow.CreatedAt,
		expectedRow.UpdatedAt,
	)

	mockDB.ExpectQuery(`SELECT \* FROM "activities" WHERE "activities"\."id" = .+ AND "activities"\."name" = .+`).WillReturnRows(resultRows)

	result, err := activitiesProvider.CreateActivity(context.Background(), expectedRow)

	if err != nil {
		t.Fatalf("error in creating activity %v", err)
	}

	if err := helpers.CompareReflectedStructFields(*result, expectedRow); err != nil {
		t.Fatal(err)
	}
}

func TestUpdateActivitySuccess(t *testing.T) {
	mockGorm, mockDB := mocks.NewMockDB()
	activitiesProvider := NewActivitiesProvider(mockGorm)

	expectedRow := structures.Activity{
		Id:          1,
		Name:        "speed drill",
		Description: "test speed drill",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mockDB.ExpectBegin()
	mockDB.ExpectExec(`UPDATE "activities" SET "name"=.+,"description"=.+,"created_at"=.+,"updated_at"=.+ WHERE "id" = .+`).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	result, err := activitiesProvider.UpdateActivity(context.Background(), expectedRow)

	if err != nil {
		t.Fatalf("error in updating activity %v", err)
	}

	if err := helpers.CompareReflectedStructFields(*result, expectedRow); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteActivitySuccess(t *testing.T) {
	mockGorm, mockDB := mocks.NewMockDB()
	activitiesProvider := NewActivitiesProvider(mockGorm)

	expectedRow := structures.Activity{
		Id: 1,
	}

	mockDB.ExpectBegin()
	mockDB.ExpectExec(`FROM "activities" WHERE "activities"\."id" = .+`).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	result, err := activitiesProvider.DeleteActivity(context.Background(), expectedRow.Id)

	if err != nil {
		t.Fatalf("error in deleting activity %v", err)
	}

	if result != true {
		t.Fatalf("unexpected result: expected %v, got %v", true, result)
	}
}
