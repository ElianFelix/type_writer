package providers

import (
	"context"
	"type_writer_api/structures"

	"gorm.io/gorm"
)

type ScoreProviderInterface interface {
	GetScores(ctx context.Context) ([]*structures.Score, error)
	GetScoreById(ctx context.Context, scoreId int) (*structures.Score, error)
	CreateScore(ctx context.Context, textInfo structures.Score) (*structures.Score, error)
	UpdateScore(ctx context.Context, updatedtextInfo structures.Score) (*structures.Score, error)
	DeleteScore(ctx context.Context, scoreId int) (bool, error)
}

type ScoreProvider struct {
	Db *gorm.DB
}

func (t *ScoreProvider) GetScores(ctx context.Context) ([]*structures.Score, error) {
	var scores []*structures.Score
	err := t.Db.WithContext(ctx).Table(structures.SCORE_TABLE_NAME).Find(&scores).Error
	if err != nil {
		return nil, err
	}
	return scores, nil
}

func (t *ScoreProvider) GetScoreById(ctx context.Context, scoreId int) (*structures.Score, error) {
	var score *structures.Score
	err := t.Db.WithContext(ctx).Table(structures.SCORE_TABLE_NAME).
		First(&score, "id = ? ", scoreId).Error
	if err != nil {
		return nil, err
	}
	return score, nil
}

func (t *ScoreProvider) CreateScore(ctx context.Context, textInfo structures.Score) (*structures.Score, error) {
	var score *structures.Score
	err := t.Db.WithContext(ctx).Table(structures.SCORE_TABLE_NAME).FirstOrCreate(&score, &textInfo).Error
	if err != nil {
		return nil, err
	}
	return score, nil
}

func (t *ScoreProvider) UpdateScore(ctx context.Context, updatedScoreInfo structures.Score) (*structures.Score, error) {
	err := t.Db.WithContext(ctx).Table(structures.SCORE_TABLE_NAME).Updates(&updatedScoreInfo).Error
	if err != nil {
		return nil, err
	}
	return &updatedScoreInfo, nil
}

func (t *ScoreProvider) DeleteScore(ctx context.Context, scoreId int) (bool, error) {
	var deleteScore = structures.Score{Id: scoreId}
	err := t.Db.WithContext(ctx).Table(structures.SCORE_TABLE_NAME).Delete(&deleteScore).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewScoreProvider(db *gorm.DB) *ScoreProvider {
	return &ScoreProvider{
		Db: db,
	}
}
