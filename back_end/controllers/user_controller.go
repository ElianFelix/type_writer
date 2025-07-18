package controllers

import (
	"log/slog"
	"net/http"
	"strconv"
	"type_writer_api/services"
	"type_writer_api/structures"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController struct {
	UserService services.UserServiceInterface
}

func (u *UserController) GetUsers(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()

	users, err := u.UserService.GetUsers(reqCtx)
	if err != nil {
		slog.ErrorContext(reqCtx, "error fetching users", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error fetching users")
	}

	return ctx.JSON(http.StatusOK, struct{ Users []*structures.UserResp }{Users: users})
}

func (u *UserController) GetUser(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()
	var (
		userId   int
		username string
		err      error
	)

	userId, err = strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		username = ctx.Param("user_id")
	}

	user, err := u.UserService.GetUserByIdOrUsername(reqCtx, userId, username)
	if err != nil && err == gorm.ErrRecordNotFound {
		slog.ErrorContext(reqCtx, "user not found", "error", err)
		return ctx.JSON(http.StatusNotFound, "user not found")
	} else if err != nil {
		slog.ErrorContext(reqCtx, "error fetching user", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error fetching user")
	}

	return ctx.JSON(http.StatusOK, user)
}

func (u *UserController) CreateUser(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()
	req := structures.UserReq{}

	err := ctx.Bind(&req)
	if err != nil {
		slog.ErrorContext(reqCtx, "error binding request body", "error", err)
		return ctx.JSON(http.StatusBadRequest, "error binding request body, incomplete or bad request")
	}

	createdUser, err := u.UserService.CreateUser(reqCtx, req)
	if err != nil {
		slog.ErrorContext(reqCtx, "error creating new user", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error creating new user")
	}

	return ctx.JSON(http.StatusCreated, createdUser)
}

func (u *UserController) UpdateUser(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()
	var (
		req    structures.UserReq
		userId int
		err    error
	)

	userId, err = strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		slog.ErrorContext(reqCtx, "bad user id in request", "error", err)
		return ctx.JSON(http.StatusBadRequest, "bad user id in request")
	}

	err = ctx.Bind(&req)
	if err != nil {
		slog.ErrorContext(reqCtx, "error binding request body", "error", err)
		return ctx.JSON(http.StatusBadRequest, "error binding request body, incomplete or bad request")
	}

	updatedUser, err := u.UserService.UpdateUser(reqCtx, req, userId)
	if err != nil {
		slog.ErrorContext(reqCtx, "error updating user", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error updating user")
	}

	return ctx.JSON(http.StatusOK, updatedUser)
}

func (u *UserController) DeleteUser(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()
	var (
		userId int
		err    error
	)

	userId, err = strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		slog.ErrorContext(reqCtx, "bad user id in request", "error", err)
		return ctx.JSON(http.StatusBadRequest, "bad user id in request")
	}

	user, err := u.UserService.DeleteUser(reqCtx, userId)
	if err != nil && err == gorm.ErrRecordNotFound {
		slog.ErrorContext(reqCtx, "user not found", "error", err)
		return ctx.JSON(http.StatusNotFound, "user not found")
	} else if err != nil {
		slog.ErrorContext(reqCtx, "error deleting user", "error", err)
		return ctx.JSON(http.StatusInternalServerError, "error deleting user")
	}

	return ctx.JSON(http.StatusOK, user)
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}
