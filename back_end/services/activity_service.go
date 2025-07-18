package services

import (
	"context"
	"log/slog"
	"type_writer_api/providers"
	"type_writer_api/structures"
)

type ActivityServiceInterface interface {
	GetActivities(ctx context.Context) ([]*structures.Activity, error)
	GetActivityByIdOrName(ctx context.Context, activityId int, name string) (*structures.Activity, error)
	CreateActivity(ctx context.Context, activityInfo structures.ActivityReq) (*structures.Activity, error)
	UpdateActivity(ctx context.Context, activityInfo structures.ActivityReq, activityId int) (*structures.Activity, error)
	DeleteActivity(ctx context.Context, activityId int) (bool, error)
}

type ActivityService struct {
	ActivityProvider providers.ActivityProviderInterface
}

func (a *ActivityService) GetActivities(ctx context.Context) ([]*structures.Activity, error) {
	var result []*structures.Activity

	activities, err := a.ActivityProvider.GetActivities(ctx)
	if err != nil {
		return nil, err
	}

	for _, activity := range activities {
		result = append(result, activity)
	}

	return result, nil
}

func (a *ActivityService) GetActivityByIdOrName(ctx context.Context, activityId int, name string) (*structures.Activity, error) {
	activity, err := a.ActivityProvider.GetActivityByIdOrName(ctx, activityId, name)
	if err != nil {
		return nil, err
	}

	result := activity
	return result, nil
}

func (a *ActivityService) CreateActivity(ctx context.Context, activityInfo structures.ActivityReq) (*structures.Activity, error) {
	activityToCreate := structures.ConvertRequestToActivity(&activityInfo)

	createdActivity, err := a.ActivityProvider.CreateActivity(ctx, *activityToCreate)
	if err != nil {
		slog.ErrorContext(ctx, "failed to create activity", "error", err)
		return nil, err
	}

	result := createdActivity
	return result, nil
}

func (a *ActivityService) UpdateActivity(ctx context.Context, activityInfo structures.ActivityReq, activityId int) (*structures.Activity, error) {
	existingActivity, err := a.ActivityProvider.GetActivityByIdOrName(ctx, activityId, "")

	if err != nil {
		slog.ErrorContext(ctx, "failed to update activity", "error", err)
		return nil, err
	}

	if activityInfo.Name != "" {
		existingActivity.Name = activityInfo.Name
	}
	if activityInfo.Description != "" {
		existingActivity.Description = activityInfo.Description
	}

	updatedActivity, err := a.ActivityProvider.UpdateActivity(ctx, *existingActivity)
	if err != nil {
		slog.ErrorContext(ctx, "failed to update activity", "error", err)
		return nil, err
	}

	result := updatedActivity
	return result, nil
}

func (t *ActivityService) DeleteActivity(ctx context.Context, activityId int) (bool, error) {
	deleted, err := t.ActivityProvider.DeleteActivity(ctx, activityId)
	if err != nil {
		return false, err
	}

	return deleted, nil
}

func NewActivityService(activityProvider *providers.ActivityProvider) *ActivityService {
	return &ActivityService{
		ActivityProvider: activityProvider,
	}
}
