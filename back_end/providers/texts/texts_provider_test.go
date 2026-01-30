package texts_provider

import (
	"context"
	"testing"
	"time"
	"type_writer_api/helpers"
	"type_writer_api/structures"
	"type_writer_api/testing/mocks"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetTextsSuccess(t *testing.T) {
	mockGorm, mockDB := mocks.NewMockDB()
	textsProvider := NewTextsProvider(mockGorm)

	expectedRows := []structures.Text{
		{
			Id:         1,
			TextType:   "drill",
			Title:      "test drill",
			Difficulty: "easy",
			TextBody:   "ffff jjjj fj fj",
			TextLength: 15,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			Id:         2,
			TextType:   "drill",
			Title:      "test drill 2",
			Difficulty: "easy",
			TextBody:   "ffff jjjj fj fj fj",
			TextLength: 18,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	resultRows := sqlmock.NewRows([]string{
		"id",
		"text_type",
		"title",
		"difficulty",
		"text_body",
		"text_length",
		"created_at",
		"updated_at",
	})

	for _, expectedRow := range expectedRows {
		resultRows.AddRow(
			expectedRow.Id,
			expectedRow.TextType,
			expectedRow.Title,
			expectedRow.Difficulty,
			expectedRow.TextBody,
			expectedRow.TextLength,
			expectedRow.CreatedAt,
			expectedRow.UpdatedAt,
		)
	}

	mockDB.ExpectQuery(`SELECT \* FROM "texts"`).WillReturnRows(resultRows)

	result, err := textsProvider.GetTexts(context.Background())

	if err != nil {
		t.Fatalf("error in fetching texts %v", err)
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

func TestGetTextByIdOrNameSuccess(t *testing.T) {
	mockGorm, mockDB := mocks.NewMockDB()
	textsProvider := NewTextsProvider(mockGorm)

	expectedRow := structures.Text{
		Id:         1,
		TextType:   "drill",
		Title:      "test drill",
		Difficulty: "easy",
		TextBody:   "ffff jjjj fj fj",
		TextLength: 15,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	resultRows := sqlmock.NewRows([]string{
		"id",
		"text_type",
		"title",
		"difficulty",
		"text_body",
		"text_length",
		"created_at",
		"updated_at",
	}).AddRow(
		expectedRow.Id,
		expectedRow.TextType,
		expectedRow.Title,
		expectedRow.Difficulty,
		expectedRow.TextBody,
		expectedRow.TextLength,
		expectedRow.CreatedAt,
		expectedRow.UpdatedAt,
	)

	mockDB.ExpectQuery(`SELECT \* FROM "texts" WHERE id = .+ ORDER BY "texts"\."id" LIMIT .+`).WillReturnRows(resultRows)

	inputId := 1
	inputUsername := ""
	result, err := textsProvider.GetTextByIdOrTitle(context.Background(), &inputId, &inputUsername)

	if err != nil {
		t.Fatalf("error in fetching text %v", err)
	}

	if err := helpers.CompareReflectedStructFields(*result, expectedRow); err != nil {
		t.Fatal(err)
	}
}

func TestCreateTextSuccess(t *testing.T) {
	mockGorm, mockDB := mocks.NewMockDB()
	textsProvider := NewTextsProvider(mockGorm)

	expectedRow := structures.Text{
		Id:         1,
		TextType:   "drill",
		Title:      "test drill",
		Difficulty: "easy",
		TextBody:   "ffff jjjj fj fj",
		TextLength: 15,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	resultRows := sqlmock.NewRows([]string{
		"id",
		"text_type",
		"title",
		"difficulty",
		"text_body",
		"text_length",
		"created_at",
		"updated_at",
	}).AddRow(
		expectedRow.Id,
		expectedRow.TextType,
		expectedRow.Title,
		expectedRow.Difficulty,
		expectedRow.TextBody,
		expectedRow.TextLength,
		expectedRow.CreatedAt,
		expectedRow.UpdatedAt,
	)

	mockDB.ExpectQuery(`SELECT \* FROM "texts" WHERE "texts"\."id" = .+ AND "texts"\."text_type" = .+ AND "texts"\."title" = .+ AND "texts"\."difficulty" = .+ AND "texts"\."text_body" = .+ AND "texts"\."text_length" = .+ AND "texts"\."created_at" = .+ AND "texts"\."updated_at" = .+ ORDER BY "texts"\."id" LIMIT .+`).WillReturnRows(resultRows)

	result, err := textsProvider.CreateText(context.Background(), expectedRow)

	if err != nil {
		t.Fatalf("error in creating text %v", err)
	}

	if err := helpers.CompareReflectedStructFields(*result, expectedRow); err != nil {
		t.Fatal(err)
	}
}

func TestUpdateTextSuccess(t *testing.T) {
	mockGorm, mockDB := mocks.NewMockDB()
	textsProvider := NewTextsProvider(mockGorm)

	expectedRow := structures.Text{
		Id:         1,
		TextType:   "drill",
		Title:      "test drill",
		Difficulty: "easy",
		TextBody:   "ffff jjjj fj fj",
		TextLength: 15,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	mockDB.ExpectBegin()
	mockDB.ExpectExec(`UPDATE "texts" SET "text_type"=.+,"title"=.+,"difficulty"=.+,"text_body"=.+,"text_length"=.+,"created_at"=.+,"updated_at"=.+ WHERE "id" = .+`).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	result, err := textsProvider.UpdateText(context.Background(), expectedRow)

	if err != nil {
		t.Fatalf("error in updating text %v", err)
	}

	if err := helpers.CompareReflectedStructFields(*result, expectedRow); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteTextSuccess(t *testing.T) {
	mockGorm, mockDB := mocks.NewMockDB()
	textsProvider := NewTextsProvider(mockGorm)

	expectedRow := structures.Text{
		Id: 1,
	}

	mockDB.ExpectBegin()
	mockDB.ExpectExec(`FROM "texts" WHERE "texts"\."id" = .+`).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	result, err := textsProvider.DeleteText(context.Background(), expectedRow.Id)

	if err != nil {
		t.Fatalf("error in deleting text %v", err)
	}

	if result != true {
		t.Fatalf("unexpected result: expected %v, got %v", true, result)
	}
}
