package scores_service

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

func TestGetScores(t *testing.T) {
	var (
		mockResult1 = []*structures.Score{
			{1, 1, 1, 1, 100, 60, 0, time.Now(), time.Now()},
			{2, 1, 1, 1, 100, 60, 0, time.Now(), time.Now()},
		}
		expectedResult1 = []*structures.Score{
			{1, 1, 1, 1, 100, 60, 0, time.Now(), time.Now()},
			{2, 1, 1, 1, 100, 60, 0, time.Now(), time.Now()},
		}
		mockResult2 = []*structures.Score{}
		expectedResult2 = []*structures.Score{}
	)
	data := []struct{
		testName string
		mockResult []*structures.Score
		mockErr error
		expectedResult []*structures.Score
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

	mockScoresProvider := mockProviders.NewMockScoresProviderInterface(ctrl)
	scoresService := NewScoresService(mockScoresProvider)

	for _, testCase := range data {
		t.Run(testCase.testName, func(t *testing.T) {
			mockScoresProvider.EXPECT().GetScores(context.Background()).Return(testCase.mockResult, testCase.mockErr).Times(1)

			result, err := scoresService.GetScores(context.Background())

			if testCase.expectedErr != nil {
				if err != testCase.expectedErr {
					t.Fatalf("expected error: %v but got %v instead", testCase.expectedErr, err)
				}
			}
			if len(result) != len(testCase.expectedResult) {
				t.Fatalf("slice length missmatch: got %v, expected %v", len(result), len(testCase.expectedResult))
			}

			for idx, activity := range result {
				err := helpers.CompareReflectedStructFields(*activity, *testCase.expectedResult[idx])
				if err != nil {
					t.Fatalf("row %v failed: %v\n", idx, err.Error())
				}
			}
		})
	}
}

func TestGetScoreById(t *testing.T) {
	data := []struct{
		testName string
		inputId int
		mockResult *structures.Score
		mockErr error
		expectedResult *structures.Score
		expectedErr error
	}{
		{
			"valid id",
			1,
			&structures.Score{1, 1, 1, 1, 100, 60, 0, time.Now(), time.Now()},
			nil,
			&structures.Score{1, 1, 1, 1, 100, 60, 0, time.Now(), time.Now()},
			nil,
		},
		{
			"not found error",
			0,
			nil,
			gorm.ErrRecordNotFound,
			nil,
			gorm.ErrRecordNotFound,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockScoresProvider := mockProviders.NewMockScoresProviderInterface(ctrl)
	scoresService := NewScoresService(mockScoresProvider)

	for _, testCase := range data {
		t.Run(testCase.testName, func(t *testing.T) {
			mockScoresProvider.EXPECT().GetScoreById(context.Background(), testCase.inputId).Return(testCase.mockResult, testCase.mockErr).Times(1)

			result, err := scoresService.GetScoreById(context.Background(), testCase.inputId)

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

func TestCreateScore(t *testing.T) {
	data := []struct{
		testName string
		inputScore structures.ScoreReq
		mockResult *structures.Score
		mockErr error
		expectedResult *structures.Score
		expectedErr error
	}{
		{
			"valid input activity request",
			structures.ScoreReq{1, 1, 1, 100, 60, 0},
			&structures.Score{1, 1, 1, 1, 100, 60, 0, time.Now(), time.Now()},
			nil,
			&structures.Score{1, 1, 1, 1, 100, 60, 0, time.Now(), time.Now()},
			nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockScoresProvider := mockProviders.NewMockScoresProviderInterface(ctrl)
	scoresService := NewScoresService(mockScoresProvider)

	for _, testCase := range data {
		t.Run(testCase.testName, func(t *testing.T) {
			mockScoresProvider.EXPECT().CreateScore(context.Background(), *structures.ConvertRequestToScore(&testCase.inputScore)).Return(testCase.mockResult, testCase.mockErr).Times(1)

			result, err := scoresService.CreateScore(context.Background(), testCase.inputScore)

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

func TestUpdateScore(t *testing.T) {
	data := []struct{
		testName string
		inputScore structures.ScoreReq
		inputUpdateId int
		mockQueryResult *structures.Score
		mockResult *structures.Score
		mockErr error
		expectedResult *structures.Score
		expectedErr error
	}{
		{
			"valid input activity request",
			structures.ScoreReq{1, 1, 1, 100, 60, 0},
			1,
			&structures.Score{1, 1, 1, 1, 100, 60, 0, time.Now(), time.Now()},
			&structures.Score{1, 1, 1, 1, 100, 60, 0, time.Now(), time.Now()},
			nil,
			&structures.Score{1, 1, 1, 1, 100, 60, 0, time.Now(), time.Now()},
			nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockScoresProvider := mockProviders.NewMockScoresProviderInterface(ctrl)
	scoresService := NewScoresService(mockScoresProvider)

	for _, testCase := range data {
		t.Run(testCase.testName, func(t *testing.T) {
			mockScoresProvider.EXPECT().GetScoreById(context.Background(), testCase.inputUpdateId).Return(testCase.mockQueryResult, testCase.mockErr).Times(1)
			mockScoresProvider.EXPECT().UpdateScore(
				context.Background(),
				gomock.Cond(func(input structures.Score) bool { return helpers.CompareReflectedStructFields(input, *testCase.expectedResult) == nil}),
			).Return(testCase.mockResult, testCase.mockErr).Times(1)

			result, err := scoresService.UpdateScore(context.Background(), testCase.inputScore, testCase.inputUpdateId)

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

func TestDeleteActivity(t *testing.T) {
	data := []struct{
		testName string
		inputDeleteId int
		mockResult bool
		mockErr error
		expectedResult bool
		expectedErr error
	}{
		{
			"valid input score id",
			1,
			true,
			nil,
			true,
			nil,
		},
		{
			"not existant input score id",
			99,
			false,
			nil,
			false,
			nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockScoresProvider := mockProviders.NewMockScoresProviderInterface(ctrl)
	scoresService := NewScoresService(mockScoresProvider)

	for _, testCase := range data {
		t.Run(testCase.testName, func(t *testing.T) {
			mockScoresProvider.EXPECT().DeleteScore(context.Background(), testCase.inputDeleteId).Return(testCase.mockResult, testCase.mockErr).Times(1)

			result, err := scoresService.DeleteScore(context.Background(), testCase.inputDeleteId)

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
