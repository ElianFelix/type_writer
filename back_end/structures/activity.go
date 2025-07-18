package structures

import "time"

const ACTIVITY_TABLE_NAME = "activities"

type Activity struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ActivityReq struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func ConvertRequestToActivity(req *ActivityReq) *Activity {
	return &Activity{
		Name:        req.Name,
		Description: req.Description,
	}
}
