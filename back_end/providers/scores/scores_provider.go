package scores_provider

import (
	"context"
	"type_writer_api/structures"

	"gorm.io/gorm"
)

type ScoresProviderInterface interface {
	GetScores(ctx context.Context) ([]*structures.Score, error)
	GetScoreById(ctx context.Context, scoreId int) (*structures.Score, error)
	CreateScore(ctx context.Context, scoreInfo structures.Score) (*structures.Score, error)
	UpdateScore(ctx context.Context, updatedtextInfo structures.Score) (*structures.Score, error)
	DeleteScore(ctx context.Context, scoreId int) (bool, error)
}

type ScoresProvider struct {
	Db *gorm.DB
}

func (t *ScoresProvider) GetScores(ctx context.Context) ([]*structures.Score, error) {
	var scores []*structures.Score
	err := t.Db.WithContext(ctx).Table(structures.SCORE_TABLE_NAME).Find(&scores).Error
	if err != nil {
		return nil, err
	}
	return scores, nil
}

func (t *ScoresProvider) GetScoreById(ctx context.Context, scoreId int) (*structures.Score, error) {
	var score *structures.Score
	err := t.Db.WithContext(ctx).Table(structures.SCORE_TABLE_NAME).
		First(&score, "id = ? ", scoreId).Error
	if err != nil {
		return nil, err
	}
	return score, nil
}

func (t *ScoresProvider) CreateScore(ctx context.Context, scoreInfo structures.Score) (*structures.Score, error) {
	var score *structures.Score
	err := t.Db.WithContext(ctx).Table(structures.SCORE_TABLE_NAME).FirstOrCreate(&score, &scoreInfo).Error
	if err != nil {
		return nil, err
	}
	return score, nil
}

func (t *ScoresProvider) UpdateScore(ctx context.Context, updatedScoreInfo structures.Score) (*structures.Score, error) {
	err := t.Db.WithContext(ctx).Table(structures.SCORE_TABLE_NAME).Updates(&updatedScoreInfo).Error
	if err != nil {
		return nil, err
	}
	return &updatedScoreInfo, nil
}

func (t *ScoresProvider) DeleteScore(ctx context.Context, scoreId int) (bool, error) {
	var deleteScore = structures.Score{Id: scoreId}
	err := t.Db.WithContext(ctx).Table(structures.SCORE_TABLE_NAME).Delete(&deleteScore).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewScoresProvider(db *gorm.DB) *ScoresProvider {
	return &ScoresProvider{
		Db: db,
	}
}
