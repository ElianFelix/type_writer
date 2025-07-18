package structures

import "time"

const USER_TABLE_NAME = "users"

type User struct {
	Id         int       `json:"id"`
	UserType   string    `json:"user_type"`
	Username   string    `json:"username"`
	PasswdHash string    `json:"passwd_hash"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserReq struct {
	UserType string `json:"user_type,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
}

type UserResp struct {
	Id        int       `json:"id"`
	UserType  string    `json:"user_type"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ConvertRequestToUser(req *UserReq) *User {
	return &User{
		UserType: req.UserType,
		Username: req.Username,
		Name:     req.Name,
		Email:    req.Email,
	}
}

func ConvertUserToResponse(user *User) *UserResp {
	return &UserResp{
		Id:        user.Id,
		UserType:  user.UserType,
		Username:  user.Username,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// func ConvertRequestToUpdateFieldMap(req *UserReq) map[string]interface{} {
// 	updateUserFields := map[string]interface{}{}
//
// 	if req.UserType != "" {
// 		updateUserFields["user_type"] = req.UserType
// 	}
// 	if req.Username != "" {
// 		updateUserFields["username"] = req.Username
// 	}
// 	if req.Name != "" {
// 		updateUserFields["name"] = req.Name
// 	}
// 	if req.Email != "" {
// 		updateUserFields["email"] = req.Email
// 	}
//
// 	return updateUserFields
// }
