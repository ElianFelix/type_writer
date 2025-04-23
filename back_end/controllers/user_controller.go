package controllers

import (
	"net/http"
	"type_writer_api/services"
	"type_writer_api/structures"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService *services.UserService
}

func (u *UserController) GetUsers(ctx echo.Context) error {
	users, err := u.UserService.GetUsers(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error fetching users" )
	}
	return ctx.JSON(http.StatusOK, struct{Users []*structures.UserResp}{Users: users})
}

func (u *UserController) GetUser(ctx echo.Context) error {
	req := structures.UserReq{}
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Error fetching user, incomplete or bad request")
	}
	return ctx.NoContent(http.StatusNotImplemented)
}



func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}
