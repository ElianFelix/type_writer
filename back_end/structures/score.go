package structures

import "time"

const SCORE_TABLE_NAME = "scores"

type Score struct {
	Id         int       `json:"id"`
	UserId     int       `json:"user_id"`
	ActivityId int       `json:"activity_id"`
	TextId     int       `json:"text_id"`
	Points     int       `json:"points"`
	Duration   int       `json:"duration"`
	Errors     int       `json:"errors"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ScoreReq struct {
	UserId     int `json:"user_id,omitempty"`
	ActivityId int `json:"activity_id,omitempty"`
	TextId     int `json:"text_id,omitempty"`
	Points     int `json:"points,omitempty"`
	Duration   int `json:"duration,omitempty"`
	Errors     int `json:"errors,omitempty"`
}

func ConvertRequestToScore(req *ScoreReq) *Score {
	return &Score{
		UserId:     req.UserId,
		ActivityId: req.ActivityId,
		TextId:     req.TextId,
		Points:     req.Points,
		Duration:   req.Duration,
		Errors:     req.Errors,
	}
}
