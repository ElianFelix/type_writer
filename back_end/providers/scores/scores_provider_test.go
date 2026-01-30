package scores_provider

import (
	"context"
	"testing"
	"time"
	"type_writer_api/structures"
	"type_writer_api/testing/mocks"
	"type_writer_api/helpers"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetScoresSuccess(t *testing.T) {
	mockGorm, mockDB := mocks.NewMockDB()
	scoresProvider := NewScoresProvider(mockGorm)

	expectedRows := []structures.Score{
		{
			Id:         1,
			UserId:     1,
			ActivityId: 1,
			TextId:     1,
			Points:     10,
			Duration:   60,
			Errors:     0,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			Id:         2,
			UserId:     2,
			ActivityId: 2,
			TextId:     2,
			Points:     10,
			Duration:   60,
			Errors:     0,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	resultRows := sqlmock.NewRows([]string{
		"id",
		"user_id",
		"activity_id",
		"text_id",
		"points",
		"duration",
		"errors",
		"created_at",
		"updated_at",
	})

	for _, expectedRow := range expectedRows {
		resultRows.AddRow(
			expectedRow.Id,
			expectedRow.UserId,
			expectedRow.ActivityId,
			expectedRow.TextId,
			expectedRow.Points,
			expectedRow.Duration,
			expectedRow.Errors,
			expectedRow.CreatedAt,
			expectedRow.UpdatedAt,
		)
	}

	mockDB.ExpectQuery(`SELECT \* FROM "scores"`).WillReturnRows(resultRows)

	result, err := scoresProvider.GetScores(context.Background())

	if err != nil {
		t.Fatalf("error in fetching scores %v", err)
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

func TestGetScoreByIdOrNameSuccess(t *testing.T) {
	mockGorm, mockDB := mocks.NewMockDB()
	scoresProvider := NewScoresProvider(mockGorm)

	expectedRow := structures.Score{
		Id:         1,
		UserId:     1,
		ActivityId: 1,
		TextId:     1,
		Points:     10,
		Duration:   60,
		Errors:     0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	resultRows := sqlmock.NewRows([]string{
		"id",
		"user_id",
		"activity_id",
		"text_id",
		"points",
		"duration",
		"errors",
		"created_at",
		"updated_at",
	}).AddRow(
		expectedRow.Id,
		expectedRow.UserId,
		expectedRow.ActivityId,
		expectedRow.TextId,
		expectedRow.Points,
		expectedRow.Duration,
		expectedRow.Errors,
		expectedRow.CreatedAt,
		expectedRow.UpdatedAt,
	)

	mockDB.ExpectQuery(`SELECT \* FROM "scores" WHERE id = .+ ORDER BY "scores"\."id" LIMIT .+`).WillReturnRows(resultRows)

	result, err := scoresProvider.GetScoreById(context.Background(), 1)

	if err != nil {
		t.Fatalf("error in fetching scores %v", err)
	}

	if err := helpers.CompareReflectedStructFields(*result, expectedRow); err != nil {
		t.Fatal(err)
	}
}

func TestCreateScoreSuccess(t *testing.T) {
	mockGorm, mockDB := mocks.NewMockDB()
	scoresProvider := NewScoresProvider(mockGorm)

	expectedRow := structures.Score{
		Id:         1,
		UserId:     1,
		ActivityId: 1,
		TextId:     1,
		Points:     10,
		Duration:   60,
		Errors:     0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	resultRows := sqlmock.NewRows([]string{
		"id",
		"user_id",
		"activity_id",
		"text_id",
		"points",
		"duration",
		"errors",
		"created_at",
		"updated_at",
	}).AddRow(
		expectedRow.Id,
		expectedRow.UserId,
		expectedRow.ActivityId,
		expectedRow.TextId,
		expectedRow.Points,
		expectedRow.Duration,
		expectedRow.Errors,
		expectedRow.CreatedAt,
		expectedRow.UpdatedAt,
	)

	mockDB.ExpectQuery(`SELECT \* FROM "scores" WHERE "scores"\."id" = .+ AND "scores"\."user_id" = .+ AND "scores"\."activity_id" = .+ AND "scores"\."text_id" = .+`).WillReturnRows(resultRows)

	result, err := scoresProvider.CreateScore(context.Background(), expectedRow)

	if err != nil {
		t.Fatalf("error in fetching score %v", err)
	}

	if err := helpers.CompareReflectedStructFields(*result, expectedRow); err != nil {
		t.Fatal(err)
	}
}

func TestUpdateScoreSuccess(t *testing.T) {
	mockGorm, mockDB := mocks.NewMockDB()
	scoresProvider := NewScoresProvider(mockGorm)

	expectedRow := structures.Score{
		Id:         1,
		UserId:     1,
		ActivityId: 1,
		TextId:     1,
		Points:     10,
		Duration:   60,
		Errors:     0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	mockDB.ExpectBegin()
	mockDB.ExpectExec(`UPDATE "scores" SET "user_id"=.+,"activity_id"=.+,"text_id"=.+,"points"=.+,"duration"=.+,"created_at"=.+,"updated_at"=.+ WHERE "id" = .+`).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	result, err := scoresProvider.UpdateScore(context.Background(), expectedRow)

	if err != nil {
		t.Fatalf("error in updating score %v", err)
	}

	if err := helpers.CompareReflectedStructFields(*result, expectedRow); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteScoreSuccess(t *testing.T) {
	mockGorm, mockDB := mocks.NewMockDB()
	scoresProvider := NewScoresProvider(mockGorm)

	expectedRow := structures.Score{
		Id: 1,
	}

	mockDB.ExpectBegin()
	mockDB.ExpectExec(`FROM "scores" WHERE "scores"\."id" = .+`).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	result, err := scoresProvider.DeleteScore(context.Background(), expectedRow.Id)

	if err != nil {
		t.Fatalf("error in deleting score %v", err)
	}

	if result != true {
		t.Fatalf("unexpected result: expected %v, got %v", true, result)
	}
}
