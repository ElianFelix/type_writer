package structures

import "time"

const TEXT_TABLE_NAME = "texts"

type Text struct {
	Id         int       	`json:"id"`
	TextType   string    	`json:"text_type"`
	Title      string    	`json:"title"`
	Difficulty string    	`json:"difficulty"`
	Tags       []string  	`json:"tags" gorm:"serializer:json"`
	TextBody   string    	`json:"text_body"`
	TextLength int       	`json:"text_length"`
	CreatedAt  time.Time 	`json:"created_at"`
	UpdatedAt  time.Time 	`json:"updated_at"`
}

type TextReq struct {
	TextType   string 		`json:"text_type,omitempty"`
	Title      string 		`json:"title,omitempty"`
	Difficulty string 		`json:"difficulty,omitempty"`
	Tags       []string 	`json:"tags"`
	TextBody   string 		`json:"text_body,omitempty"`
}

func ConvertRequestToText(req *TextReq) *Text {
	return &Text{
		TextType:   req.TextType,
		Title:      req.Title,
		Difficulty: req.Difficulty,
		Tags:		req.Tags,
		TextBody:   req.TextBody,
		TextLength: len(req.TextBody),
	}
}
