package structures

import "time"

const USER_TABLE_NAME = "users"

type User struct {
	Id int `json:"id"`
	UserType string `json:"user_type"`
	Username string `json:"username"`
	PasswdHash string
	Name string `json:"name"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserReq struct {
	Id int `json:"id"`
	UserType string `json:"user_type"`
	Username string `json:"username"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type UserResp struct {
	Id int `json:"id"`
	UserType string `json:"user_type"`
	Username string `json:"username"`
	Name string `json:"name"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
