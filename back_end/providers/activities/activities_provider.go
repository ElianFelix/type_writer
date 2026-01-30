package activities_provider

import (
	"context"
	"type_writer_api/structures"

	"gorm.io/gorm"
)

type ActivitiesProviderInterface interface {
	GetActivities(ctx context.Context) ([]*structures.Activity, error)
	GetActivityByIdOrName(ctx context.Context, activityId *int, title *string) (*structures.Activity, error)
	CreateActivity(ctx context.Context, textInfo structures.Activity) (*structures.Activity, error)
	UpdateActivity(ctx context.Context, updatedtextInfo structures.Activity) (*structures.Activity, error)
	DeleteActivity(ctx context.Context, activityId int) (bool, error)
}

type ActivitiesProvider struct {
	Db *gorm.DB
}

func (t *ActivitiesProvider) GetActivities(ctx context.Context) ([]*structures.Activity, error) {
	var activities []*structures.Activity
	err := t.Db.WithContext(ctx).Table(structures.ACTIVITY_TABLE_NAME).Find(&activities).Error
	if err != nil {
		return nil, err
	}
	return activities, nil
}

func (t *ActivitiesProvider) GetActivityByIdOrName(ctx context.Context, activityId *int, name *string) (*structures.Activity, error) {
	var activity *structures.Activity
	err := t.Db.WithContext(ctx).Table(structures.ACTIVITY_TABLE_NAME).
		First(&activity, "id = ? OR name = ?", activityId, name).Error
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func (t *ActivitiesProvider) CreateActivity(ctx context.Context, activityInfo structures.Activity) (*structures.Activity, error) {
	var activity *structures.Activity
	err := t.Db.WithContext(ctx).Table(structures.ACTIVITY_TABLE_NAME).FirstOrCreate(&activity, &activityInfo).Error
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func (t *ActivitiesProvider) UpdateActivity(ctx context.Context, updatedActivityInfo structures.Activity) (*structures.Activity, error) {
	err := t.Db.WithContext(ctx).Table(structures.ACTIVITY_TABLE_NAME).Updates(&updatedActivityInfo).Error
	if err != nil {
		return nil, err
	}
	return &updatedActivityInfo, nil
}

func (t *ActivitiesProvider) DeleteActivity(ctx context.Context, activityId int) (bool, error) {
	var deleteActivity = structures.Activity{Id: activityId}
	err := t.Db.WithContext(ctx).Table(structures.ACTIVITY_TABLE_NAME).Delete(&deleteActivity).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewActivitiesProvider(db *gorm.DB) *ActivitiesProvider {
	return &ActivitiesProvider{
		Db: db,
	}
}
