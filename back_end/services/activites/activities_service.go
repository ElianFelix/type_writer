package activities_service

import (
	"context"
	"log/slog"
	"type_writer_api/providers/activities"
	"type_writer_api/structures"
)

type ActivitiesServiceInterface interface {
	GetActivities(ctx context.Context) ([]*structures.Activity, error)
	GetActivityByIdOrName(ctx context.Context, activityId *int, name *string) (*structures.Activity, error)
	CreateActivity(ctx context.Context, activityInfo structures.ActivityReq) (*structures.Activity, error)
	UpdateActivity(ctx context.Context, activityInfo structures.ActivityReq, activityId int) (*structures.Activity, error)
	DeleteActivity(ctx context.Context, activityId int) (bool, error)
}

type ActivitiesService struct {
	ActivitiesProvider activities_provider.ActivitiesProviderInterface
}

func (a *ActivitiesService) GetActivities(ctx context.Context) ([]*structures.Activity, error) {
	var result []*structures.Activity

	activities, err := a.ActivitiesProvider.GetActivities(ctx)
	if err != nil {
		return nil, err
	}

	for _, activity := range activities {
		result = append(result, activity)
	}

	return result, nil
}

func (a *ActivitiesService) GetActivityByIdOrName(ctx context.Context, activityId *int, name *string) (*structures.Activity, error) {
	activity, err := a.ActivitiesProvider.GetActivityByIdOrName(ctx, activityId, name)
	if err != nil {
		return nil, err
	}

	result := activity
	return result, nil
}

func (a *ActivitiesService) CreateActivity(ctx context.Context, activityInfo structures.ActivityReq) (*structures.Activity, error) {
	activityToCreate := structures.ConvertRequestToActivity(&activityInfo)

	createdActivity, err := a.ActivitiesProvider.CreateActivity(ctx, *activityToCreate)
	if err != nil {
		slog.ErrorContext(ctx, "failed to create activity", "error", err)
		return nil, err
	}

	result := createdActivity
	return result, nil
}

func (a *ActivitiesService) UpdateActivity(ctx context.Context, activityInfo structures.ActivityReq, activityId int) (*structures.Activity, error) {
	existingActivity, err := a.ActivitiesProvider.GetActivityByIdOrName(ctx, &activityId, nil)

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

	updatedActivity, err := a.ActivitiesProvider.UpdateActivity(ctx, *existingActivity)
	if err != nil {
		slog.ErrorContext(ctx, "failed to update activity", "error", err)
		return nil, err
	}

	result := updatedActivity
	return result, nil
}

func (t *ActivitiesService) DeleteActivity(ctx context.Context, activityId int) (bool, error) {
	deleted, err := t.ActivitiesProvider.DeleteActivity(ctx, activityId)
	if err != nil {
		return false, err
	}

	return deleted, nil
}

func NewActivitiesService(activitiesProvider activities_provider.ActivitiesProviderInterface) *ActivitiesService {
	return &ActivitiesService{
		ActivitiesProvider: activitiesProvider,
	}
}
