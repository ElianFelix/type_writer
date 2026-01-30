package texts_service

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

func TestGetTexts(t *testing.T) {
	var (
		mockResult1 = []*structures.Text{
			{1, "drill", "test text 1", "normal", "test text body", len("test text body"), time.Now(), time.Now()},
			{2, "drill", "test text 2", "normal", "test text body", len("test text body"), time.Now(), time.Now()},
		}
		expectedResult1 = []*structures.Text{
			{1, "drill", "test text 1", "normal", "test text body", len("test text body"), time.Now(), time.Now()},
			{2, "drill", "test text 2", "normal", "test text body", len("test text body"), time.Now(), time.Now()},
		}
		mockResult2 = []*structures.Text{}
		expectedResult2 = []*structures.Text{}
	)
	data := []struct{
		testName string
		mockResult []*structures.Text
		mockErr error
		expectedResult []*structures.Text
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

	mockTextsProvider := mockProviders.NewMockTextsProviderInterface(ctrl)
	textsService := NewTextsService(mockTextsProvider)

	for _, testCase := range data {
		t.Run(testCase.testName, func(t *testing.T) {
			mockTextsProvider.EXPECT().GetTexts(context.Background()).Return(testCase.mockResult, testCase.mockErr).Times(1)

			result, err := textsService.GetTexts(context.Background())

			if testCase.expectedErr != nil {
				if err != testCase.expectedErr {
					t.Fatalf("expected error: %v but got %v instead", testCase.expectedErr, err)
				}
			}
			if len(result) != len(testCase.expectedResult) {
				t.Fatalf("slice length missmatch: got %v, expected %v", len(result), len(testCase.expectedResult))
			}

			for idx, text := range result {
				err := helpers.CompareReflectedStructFields(*text, *testCase.expectedResult[idx])
				if err != nil {
					t.Fatalf("row %v failed: %v\n", idx, err.Error())
				}
			}
		})
	}
}

func TestGetTextByIdOrName(t *testing.T) {
	data := []struct{
		testName string
		inputId int
		inputName string
		mockResult *structures.Text
		mockErr error
		expectedResult *structures.Text
		expectedErr error
	}{
		{
			"valid id",
			1,
			"",
			&structures.Text{1, "drill", "test text 1", "normal", "test text body", len("test text body"), time.Now(), time.Now()},
			nil,
			&structures.Text{1, "drill", "test text 1", "normal", "test text body", len("test text body"), time.Now(), time.Now()},
			nil,
		},
		{
			"valid title",
			0,
			"test text 1",
			&structures.Text{1, "drill", "test text 1", "normal", "test text body", len("test text body"), time.Now(), time.Now()},
			nil,
			&structures.Text{1, "drill", "test text 1", "normal", "test text body", len("test text body"), time.Now(), time.Now()},
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

	mockTextsProvider := mockProviders.NewMockTextsProviderInterface(ctrl)
	textsService := NewTextsService(mockTextsProvider)

	for _, testCase := range data {
		t.Run(testCase.testName, func(t *testing.T) {
			mockTextsProvider.EXPECT().GetTextByIdOrTitle(context.Background(), &testCase.inputId, &testCase.inputName).Return(testCase.mockResult, testCase.mockErr).Times(1)

			result, err := textsService.GetTextByIdOrTitle(context.Background(), &testCase.inputId, &testCase.inputName)

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

func TestCreateText(t *testing.T) {
	data := []struct{
		testName string
		inputText structures.TextReq
		mockResult *structures.Text
		mockErr error
		expectedResult *structures.Text
		expectedErr error
	}{
		{
			"valid input text request",
			structures.TextReq{"drill", "test text 1", "normal", "test text body"},
			&structures.Text{1, "drill", "test text 1", "normal", "test text body", len("test text body"), time.Now(), time.Now()},
			nil,
			&structures.Text{1, "drill", "test text 1", "normal", "test text body", len("test text body"), time.Now(), time.Now()},
			nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTextsProvider := mockProviders.NewMockTextsProviderInterface(ctrl)
	textsService := NewTextsService(mockTextsProvider)

	for _, testCase := range data {
		t.Run(testCase.testName, func(t *testing.T) {
			mockTextsProvider.EXPECT().CreateText(context.Background(), *structures.ConvertRequestToText(&testCase.inputText)).Return(testCase.mockResult, testCase.mockErr).Times(1)

			result, err := textsService.CreateText(context.Background(), testCase.inputText)

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

func TestUpdateText(t *testing.T) {
	data := []struct{
		testName string
		inputText structures.TextReq
		inputUpdateId int
		mockQueryResult *structures.Text
		mockResult *structures.Text
		mockErr error
		expectedResult *structures.Text
		expectedErr error
	}{
		{
			"valid input text request",
			structures.TextReq{"drill", "test text 1", "normal", "test text body after update"},
			1,
			&structures.Text{1, "drill", "test text 1", "normal", "test text body", len("test text body"), time.Now(), time.Now()},
			&structures.Text{1, "drill", "test text 1", "normal", "test text body after update", len("test text body after update"), time.Now(), time.Now()},
			nil,
			&structures.Text{1, "drill", "test text 1", "normal", "test text body after update", len("test text body after update"), time.Now(), time.Now()},
			nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTextsProvider := mockProviders.NewMockTextsProviderInterface(ctrl)
	textsService := NewTextsService(mockTextsProvider)

	for _, testCase := range data {
		t.Run(testCase.testName, func(t *testing.T) {
			mockTextsProvider.EXPECT().GetTextByIdOrTitle(context.Background(), &testCase.inputUpdateId, nil).Return(testCase.mockQueryResult, testCase.mockErr).Times(1)
			mockTextsProvider.EXPECT().UpdateText(
				context.Background(),
				gomock.Cond(func(input structures.Text) bool { return helpers.CompareReflectedStructFields(input, *testCase.mockResult) == nil}),
			).Return(testCase.mockResult, testCase.mockErr).Times(1)

			result, err := textsService.UpdateText(context.Background(), testCase.inputText, testCase.inputUpdateId)

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

func TestDeleteText(t *testing.T) {
	data := []struct{
		testName string
		inputDeleteId int
		mockResult bool
		mockErr error
		expectedResult bool
		expectedErr error
	}{
		{
			"valid input text id",
			1,
			true,
			nil,
			true,
			nil,
		},
		{
			"not existant input text id",
			99,
			false,
			nil,
			false,
			nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTextsProvider := mockProviders.NewMockTextsProviderInterface(ctrl)
	textsService := NewTextsService(mockTextsProvider)

	for _, testCase := range data {
		t.Run(testCase.testName, func(t *testing.T) {
			mockTextsProvider.EXPECT().DeleteText(context.Background(), testCase.inputDeleteId).Return(testCase.mockResult, testCase.mockErr).Times(1)

			result, err := textsService.DeleteText(context.Background(), testCase.inputDeleteId)

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
