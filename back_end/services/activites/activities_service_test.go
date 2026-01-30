package activities_service

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

func TestGetActivities(t *testing.T) {
	var (
		mockResult1 = []*structures.Activity{
			{1, "test activity 1", "test activity 1 is first", time.Now(), time.Now()},
			{2, "test activity 2", "test activity 2 is second", time.Now(), time.Now()},
		}
		expectedResult1 = []*structures.Activity{
			{1, "test activity 1", "test activity 1 is first", time.Now(), time.Now()},
			{2, "test activity 2", "test activity 2 is second", time.Now(), time.Now()},
		}
		mockResult2 = []*structures.Activity{}
		expectedResult2 = []*structures.Activity{}
	)
	data := []struct{
		testName string
		mockResult []*structures.Activity
		mockErr error
		expectedResult []*structures.Activity
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

	mockActivitiesProvider := mockProviders.NewMockActivitiesProviderInterface(ctrl)
	activitiesService := NewActivitiesService(mockActivitiesProvider)

	for _, testCase := range data {
		t.Run(testCase.testName, func(t *testing.T) {
			mockActivitiesProvider.EXPECT().GetActivities(context.Background()).Return(testCase.mockResult, testCase.mockErr).Times(1)

			result, err := activitiesService.GetActivities(context.Background())

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

func TestGetActivityByIdOrName(t *testing.T) {
	data := []struct{
		testName string
		inputId int
		inputName string
		mockResult *structures.Activity
		mockErr error
		expectedResult *structures.Activity
		expectedErr error
	}{
		{
			"valid id",
			1,
			"",
			&structures.Activity{1, "test activity 1", "test activity 1 is first", time.Now(), time.Now()},
			nil,
			&structures.Activity{1, "test activity 1", "test activity 1 is first", time.Now(), time.Now()},
			nil,
		},
		{
			"valid name",
			0,
			"test activity 1",
			&structures.Activity{1, "test activity 1", "test activity 1 is first", time.Now(), time.Now()},
			nil,
			&structures.Activity{1, "test activity 1", "test activity 1 is first", time.Now(), time.Now()},
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

	mockActivitiesProvider := mockProviders.NewMockActivitiesProviderInterface(ctrl)
	activitiesService := NewActivitiesService(mockActivitiesProvider)

	for _, testCase := range data {
		t.Run(testCase.testName, func(t *testing.T) {
			mockActivitiesProvider.EXPECT().GetActivityByIdOrName(context.Background(), &testCase.inputId, &testCase.inputName).Return(testCase.mockResult, testCase.mockErr).Times(1)

			result, err := activitiesService.GetActivityByIdOrName(context.Background(), &testCase.inputId, &testCase.inputName)

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

func TestCreateActivity(t *testing.T) {
	data := []struct{
		testName string
		inputActivity structures.ActivityReq
		mockResult *structures.Activity
		mockErr error
		expectedResult *structures.Activity
		expectedErr error
	}{
		{
			"valid input activity request",
			structures.ActivityReq{"test activity 1", "test activity 1 is first"},
			&structures.Activity{1, "test activity created", "test activity after being returned by db", time.Now(), time.Now()},
			nil,
			&structures.Activity{1, "test activity created", "test activity after being returned by db", time.Now(), time.Now()},
			nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockActivitiesProvider := mockProviders.NewMockActivitiesProviderInterface(ctrl)
	activitiesService := NewActivitiesService(mockActivitiesProvider)

	for _, testCase := range data {
		t.Run(testCase.testName, func(t *testing.T) {
			mockActivitiesProvider.EXPECT().CreateActivity(context.Background(), *structures.ConvertRequestToActivity(&testCase.inputActivity)).Return(testCase.mockResult, testCase.mockErr).Times(1)

			result, err := activitiesService.CreateActivity(context.Background(), testCase.inputActivity)

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

func TestUpdateActivity(t *testing.T) {
	data := []struct{
		testName string
		inputActivity structures.ActivityReq
		inputUpdateId int
		mockQueryResult *structures.Activity
		mockResult *structures.Activity
		mockErr error
		expectedResult *structures.Activity
		expectedErr error
	}{
		{
			"valid input activity request",
			structures.ActivityReq{"test activity updated", "test activity after being updated"},
			1,
			&structures.Activity{1, "test activity created", "test activity after being returned by db", time.Now(), time.Now()},
			&structures.Activity{1, "test activity updated", "test activity after being updated", time.Now(), time.Now()},
			nil,
			&structures.Activity{1, "test activity updated", "test activity after being updated", time.Now(), time.Now()},
			nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockActivitiesProvider := mockProviders.NewMockActivitiesProviderInterface(ctrl)
	activitiesService := NewActivitiesService(mockActivitiesProvider)

	for _, testCase := range data {
		t.Run(testCase.testName, func(t *testing.T) {
			mockActivitiesProvider.EXPECT().GetActivityByIdOrName(context.Background(), &testCase.inputUpdateId, nil).Return(testCase.mockQueryResult, testCase.mockErr).Times(1)
			mockActivitiesProvider.EXPECT().UpdateActivity(
				context.Background(),
				gomock.Cond(func(input structures.Activity) bool { return helpers.CompareReflectedStructFields(input, *testCase.expectedResult) == nil}),
			).Return(testCase.mockResult, testCase.mockErr).Times(1)

			result, err := activitiesService.UpdateActivity(context.Background(), testCase.inputActivity, testCase.inputUpdateId)

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
			"valid input activity id",
			1,
			true,
			nil,
			true,
			nil,
		},
		{
			"not existant input activity id",
			99,
			false,
			nil,
			false,
			nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockActivitiesProvider := mockProviders.NewMockActivitiesProviderInterface(ctrl)
	activitiesService := NewActivitiesService(mockActivitiesProvider)

	for _, testCase := range data {
		t.Run(testCase.testName, func(t *testing.T) {
			mockActivitiesProvider.EXPECT().DeleteActivity(context.Background(), testCase.inputDeleteId).Return(testCase.mockResult, testCase.mockErr).Times(1)

			result, err := activitiesService.DeleteActivity(context.Background(), testCase.inputDeleteId)

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
