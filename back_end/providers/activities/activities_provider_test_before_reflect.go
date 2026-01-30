package activities_provider
//
// import (
// 	"context"
// 	"testing"
// 	"time"
// 	"type_writer_api/structures"
// 	"type_writer_api/testing/mocks"
//
// 	"github.com/DATA-DOG/go-sqlmock"
// )
//
// func TestGetActivitiesSuccess(t *testing.T) {
// 	mockGorm, mockDB := mocks.NewMockDB()
// 	activitiesProvider := NewActivitiesProvider(mockGorm)
//
// 	expectedRows := []structures.Activity{
// 		{
// 			Id:          1,
// 			Name:        "speed drill",
// 			Description: "test speed drill",
// 			CreatedAt:   time.Now(),
// 			UpdatedAt:   time.Now(),
// 		},
// 		{
// 			Id:          2,
// 			Name:        "article",
// 			Description: "test article",
// 			CreatedAt:   time.Now(),
// 			UpdatedAt:   time.Now(),
// 		},
// 	}
//
// 	resultRows := sqlmock.NewRows([]string{
// 		"id",
// 		"name",
// 		"description",
// 		"created_at",
// 		"updated_at",
// 	})
//
// 	for _, expectedRow := range expectedRows {
// 		resultRows.AddRow(
// 			expectedRow.Id,
// 			expectedRow.Name,
// 			expectedRow.Description,
// 			expectedRow.CreatedAt,
// 			expectedRow.UpdatedAt,
// 		)
// 	}
//
// 	mockDB.ExpectQuery(`SELECT \* FROM "activities"`).WillReturnRows(resultRows)
//
// 	result, err := activitiesProvider.GetActivities(context.Background())
//
// 	if err != nil {
// 		t.Fatalf("error in fetching activities %v", err)
// 	}
//
// 	if len(result) != 2 {
// 		t.Fatalf("unexpected result length: expected %v, got %v", 2, len(result))
// 	}
//
// 	for indx, resultRow := range result {
// 		if resultRow.Id != (&expectedRows[indx]).Id {
// 			t.Fatalf("unexpected result field: expected %v, got %v", resultRow.Id, expectedRows[indx].Id)
// 		}
// 		if resultRow.Name != (&expectedRows[indx]).Name {
// 			t.Fatalf("unexpected result field: expected %v, got %v", resultRow.Name, expectedRows[indx].Name)
// 		}
// 		if resultRow.Description != (&expectedRows[indx]).Description {
// 			t.Fatalf("unexpected result field: expected %v, got %v", resultRow.Description, expectedRows[indx].Description)
// 		}
// 	}
// }
//
// func TestGetActivityByIdOrNameSuccess(t *testing.T) {
// 	mockGorm, mockDB := mocks.NewMockDB()
// 	activitiesProvider := NewActivitiesProvider(mockGorm)
//
// 	expectedRow := structures.Activity{
// 		Id:          1,
// 		Name:        "speed drill",
// 		Description: "test speed drill",
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 	}
//
// 	resultRows := sqlmock.NewRows([]string{
// 		"id",
// 		"name",
// 		"description",
// 		"created_at",
// 		"updated_at",
// 	}).AddRow(
// 		expectedRow.Id,
// 		expectedRow.Name,
// 		expectedRow.Description,
// 		expectedRow.CreatedAt,
// 		expectedRow.UpdatedAt,
// 	)
//
// 	mockDB.ExpectQuery(`SELECT \* FROM "activities" WHERE id = .+ OR name = .+ ORDER BY "activities"\."id" LIMIT .+`).WillReturnRows(resultRows)
//
// 	result, err := activitiesProvider.GetActivityByIdOrName(context.Background(), 1, "")
//
// 	if err != nil {
// 		t.Fatalf("error in fetching activity %v", err)
// 	}
//
// 	if result.Id != expectedRow.Id {
// 		t.Fatalf("unexpected result field: expected %v, got %v", result.Id, expectedRow.Id)
// 	}
// 	if result.Name != expectedRow.Name {
// 		t.Fatalf("unexpected result field: expected %v, got %v", result.Name, expectedRow.Name)
// 	}
// 	if result.Description != expectedRow.Description {
// 		t.Fatalf("unexpected result field: expected %v, got %v", result.Description, expectedRow.Description)
// 	}
// }
//
// func TestCreateActivitySuccess(t *testing.T) {
// 	mockGorm, mockDB := mocks.NewMockDB()
// 	activitiesProvider := NewActivitiesProvider(mockGorm)
//
// 	expectedRow := structures.Activity{
// 		Id:          1,
// 		Name:        "speed drill",
// 		Description: "test speed drill",
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 	}
//
// 	resultRows := sqlmock.NewRows([]string{
// 		"id",
// 		"name",
// 		"description",
// 		"created_at",
// 		"updated_at",
// 	}).AddRow(
// 		expectedRow.Id,
// 		expectedRow.Name,
// 		expectedRow.Description,
// 		expectedRow.CreatedAt,
// 		expectedRow.UpdatedAt,
// 	)
//
// 	mockDB.ExpectQuery(`SELECT \* FROM "activities" WHERE "activities"\."id" = .+ AND "activities"\."name" = .+`).WillReturnRows(resultRows)
//
// 	result, err := activitiesProvider.CreateActivity(context.Background(), expectedRow)
//
// 	if err != nil {
// 		t.Fatalf("error in creating activity %v", err)
// 	}
//
// 	if result.Id != expectedRow.Id {
// 		t.Fatalf("unexpected result field: expected %v, got %v", result.Id, expectedRow.Id)
// 	}
// 	if result.Name != expectedRow.Name {
// 		t.Fatalf("unexpected result field: expected %v, got %v", result.Name, expectedRow.Name)
// 	}
// 	if result.Description != expectedRow.Description {
// 		t.Fatalf("unexpected result field: expected %v, got %v", result.Description, expectedRow.Description)
// 	}
// }
//
// func TestUpdateActivitySuccess(t *testing.T) {
// 	mockGorm, mockDB := mocks.NewMockDB()
// 	activitiesProvider := NewActivitiesProvider(mockGorm)
//
// 	expectedRow := structures.Activity{
// 		Id:          1,
// 		Name:        "speed drill",
// 		Description: "test speed drill",
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 	}
//
// 	mockDB.ExpectBegin()
// 	mockDB.ExpectExec(`UPDATE "activities" SET "name"=.+,"description"=.+,"created_at"=.+,"updated_at"=.+ WHERE "id" = .+`).WillReturnResult(sqlmock.NewResult(1, 1))
// 	mockDB.ExpectCommit()
//
// 	result, err := activitiesProvider.UpdateActivity(context.Background(), expectedRow)
//
// 	if err != nil {
// 		t.Fatalf("error in updating activity %v", err)
// 	}
//
// 	if result.Id != expectedRow.Id {
// 		t.Fatalf("unexpected result field: expected %v, got %v", result.Id, expectedRow.Id)
// 	}
// 	if result.Name != expectedRow.Name {
// 		t.Fatalf("unexpected result field: expected %v, got %v", result.Name, expectedRow.Name)
// 	}
// 	if result.Description != expectedRow.Description {
// 		t.Fatalf("unexpected result field: expected %v, got %v", result.Description, expectedRow.Description)
// 	}
// }
//
// func TestDeleteActivitySuccess(t *testing.T) {
// 	mockGorm, mockDB := mocks.NewMockDB()
// 	activitiesProvider := NewActivitiesProvider(mockGorm)
//
// 	expectedRow := structures.Activity{
// 		Id: 1,
// 	}
//
// 	mockDB.ExpectBegin()
// 	mockDB.ExpectExec(`FROM "activities" WHERE "activities"\."id" = .+`).WillReturnResult(sqlmock.NewResult(1, 1))
// 	mockDB.ExpectCommit()
//
// 	result, err := activitiesProvider.DeleteActivity(context.Background(), expectedRow.Id)
//
// 	if err != nil {
// 		t.Fatalf("error in deleting activity %v", err)
// 	}
//
// 	if result != true {
// 		t.Fatalf("unexpected result: expected %v, got %v", true, result)
// 	}
// }
