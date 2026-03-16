package scores_provider

import (
	"context"
	"encoding/json"
	"testing"
	"time"
	"type_writer_api/helpers"
	"type_writer_api/structures"
	"type_writer_api/testing/mocks"

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
			Duration:   60,
			Result: 	map[string]any{"wpm": float64(300), "errors": float64(300)}, // type casting ints because json unmarshalls ints into floats64
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			Id:         2,
			UserId:     2,
			ActivityId: 2,
			TextId:     2,
			Duration:   60,
			Result: 	map[string]any{"wpm": float64(300), "errors": float64(300)},
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	resultRows := sqlmock.NewRows([]string{
		"id",
		"user_id",
		"activity_id",
		"text_id",
		"duration",
		"result",
		"created_at",
		"updated_at",
	})

	for _, expectedRow := range expectedRows {
		sResult, err := json.Marshal(expectedRow.Result)
		if err != nil {
			t.Fatalf("error in serializing text tags %v", err)
		}

		resultRows.AddRow(
			expectedRow.Id,
			expectedRow.UserId,
			expectedRow.ActivityId,
			expectedRow.TextId,
			expectedRow.Duration,
			sResult,
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
		t.Fatalf("unexpected result length: expected %v,\n got %v\n", 2, len(result))
	}

	for indx, resultRow := range result {
		// fmt.Printf("result row %v, expected %v", *resultRow, expectedRows[indx])
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
		Duration:   60,
		Result: 	map[string]any{"wpm": float64(300), "errors": float64(300)},
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	sResult, err := json.Marshal(expectedRow.Result)
	if err != nil {
		t.Fatalf("error in serializing text tags %v", err)
	}

	resultRows := sqlmock.NewRows([]string{
		"id",
		"user_id",
		"activity_id",
		"text_id",
		"duration",
		"result",
		"created_at",
		"updated_at",
	}).AddRow(
		expectedRow.Id,
		expectedRow.UserId,
		expectedRow.ActivityId,
		expectedRow.TextId,
		expectedRow.Duration,
		sResult,
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
		Duration:   60,
		Result: 	map[string]any{"wpm": float64(300), "errors": float64(300)},
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	sResult, err := json.Marshal(expectedRow.Result)
	if err != nil {
		t.Fatalf("error in serializing text tags %v", err)
	}

	resultRows := sqlmock.NewRows([]string{
		"id",
		"user_id",
		"activity_id",
		"text_id",
		"duration",
		"result",
		"created_at",
		"updated_at",
	}).AddRow(
		expectedRow.Id,
		expectedRow.UserId,
		expectedRow.ActivityId,
		expectedRow.TextId,
		expectedRow.Duration,
		sResult,
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
		Duration:   60,
		Result: 	map[string]any{"wpm": float64(300), "errors": float64(300)},
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	mockDB.ExpectBegin()
	mockDB.ExpectExec(`UPDATE "scores" SET "user_id"=.+,"activity_id"=.+,"text_id"=.+,"duration"=.+,"created_at"=.+,"updated_at"=.+ WHERE "id" = .+`).WillReturnResult(sqlmock.NewResult(1, 1))
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
