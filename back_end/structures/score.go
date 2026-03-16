package structures

import "time"

const SCORE_TABLE_NAME = "scores"

type Score struct {
	Id         int       `json:"id"`
	UserId     int       		 `json:"user_id"`
	ActivityId int       		 `json:"activity_id"`
	TextId     int       		 `json:"text_id"`
	Duration   int       		 `json:"duration"`
	Result     map[string]any    `json:"result" gorm:"serializer:json"`
	CreatedAt  time.Time 		 `json:"created_at"`
	UpdatedAt  time.Time 		 `json:"updated_at"`
}

type ScoreReq struct {
	UserId     int			 	 `json:"user_id,omitempty"`
	ActivityId int			 	 `json:"activity_id,omitempty"`
	TextId     int			 	 `json:"text_id,omitempty"`
	Duration   int			 	 `json:"duration,omitempty"`
	Result     map[string]any    `json:"result"`
}

func ConvertRequestToScore(req *ScoreReq) *Score {
	return &Score{
		UserId:     req.UserId,
		ActivityId: req.ActivityId,
		TextId:     req.TextId,
		Duration:   req.Duration,
		Result:     req.Result,
	}
}
