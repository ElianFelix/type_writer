package scores_service

import (
	"context"
	"log/slog"
	"type_writer_api/providers/scores"
	"type_writer_api/structures"
)

type ScoresServiceInterface interface {
	GetScores(ctx context.Context) ([]*structures.Score, error)
	GetScoreById(ctx context.Context, scoreId int) (*structures.Score, error)
	CreateScore(ctx context.Context, scoreInfo structures.ScoreReq) (*structures.Score, error)
	UpdateScore(ctx context.Context, scoreInfo structures.ScoreReq, scoreId int) (*structures.Score, error)
	DeleteScore(ctx context.Context, scoreId int) (bool, error)
}

type ScoresService struct {
	ScoresProvider scores_provider.ScoresProviderInterface
}

func (a *ScoresService) GetScores(ctx context.Context) ([]*structures.Score, error) {
	var results []*structures.Score

	scores, err := a.ScoresProvider.GetScores(ctx)
	if err != nil {
		return nil, err
	}

	for _, score := range scores {
		results = append(results, score)
	}

	return results, nil
}

func (a *ScoresService) GetScoreById(ctx context.Context, scoreId int) (*structures.Score, error) {
	score, err := a.ScoresProvider.GetScoreById(ctx, scoreId)
	if err != nil {
		return nil, err
	}

	result := score
	return result, nil
}

func (a *ScoresService) CreateScore(ctx context.Context, scoreInfo structures.ScoreReq) (*structures.Score, error) {
	scoreToCreate := structures.ConvertRequestToScore(&scoreInfo)

	createdScore, err := a.ScoresProvider.CreateScore(ctx, *scoreToCreate)
	if err != nil {
		slog.ErrorContext(ctx, "failed to create score", "error", err)
		return nil, err
	}

	result := createdScore
	return result, nil
}

func (a *ScoresService) UpdateScore(ctx context.Context, scoreInfo structures.ScoreReq, scoreId int) (*structures.Score, error) {
	existingScore, err := a.ScoresProvider.GetScoreById(ctx, scoreId)

	if err != nil {
		slog.ErrorContext(ctx, "failed to update score", "error", err)
		return nil, err
	}

	if scoreInfo.Points != 0 {
		existingScore.Points = scoreInfo.Points
	}
	if scoreInfo.Duration != 0 {
		existingScore.Duration = scoreInfo.Duration
	}
	if scoreInfo.Errors != 0 {
		existingScore.Errors = scoreInfo.Errors
	}

	updatedScore, err := a.ScoresProvider.UpdateScore(ctx, *existingScore)
	if err != nil {
		slog.ErrorContext(ctx, "failed to update score", "error", err)
		return nil, err
	}

	result := updatedScore
	return result, nil
}

func (t *ScoresService) DeleteScore(ctx context.Context, scoreId int) (bool, error) {
	deleted, err := t.ScoresProvider.DeleteScore(ctx, scoreId)
	if err != nil {
		return false, err
	}

	return deleted, nil
}

func NewScoresService(scoresProvider scores_provider.ScoresProviderInterface) *ScoresService {
	return &ScoresService{
		ScoresProvider: scoresProvider,
	}
}
