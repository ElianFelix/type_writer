package controllers

import (
	"log/slog"
	"net/http"
	"time"
	"type_writer_api/services/users"
	"type_writer_api/structures"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	keyString []byte
	UsersService users_service.UsersServiceInterface
}

func (a *AuthController) Login(ctx echo.Context) error {
	reqCtx := ctx.Request().Context()
	req := structures.UserLoginReq{}

	err := ctx.Bind(&req)
	if err != nil {
		slog.ErrorContext(reqCtx, "error binding request body", "error", err)
		return ctx.JSON(http.StatusBadRequest, "error binding request body, incomplete or bad request")
	}

	loginUserInfo, err := a.UsersService.ValidateLoginUser(reqCtx, req.Username, req.Password)
	if err != nil {
		slog.ErrorContext(reqCtx, "failed to validate user", "error", err)
		return ctx.JSON(http.StatusUnauthorized, "failed to validate user")
	}

	// Custom claims for authorization middleware
	claims := &structures.JwtCustomClaims{
		UserType: loginUserInfo.UserType,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: loginUserInfo.Username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(a.keyString)
	if err != nil {
		slog.ErrorContext(reqCtx, "failed to validate user", "error", err)
		return ctx.JSON(http.StatusUnauthorized, "failed to validate user")
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"token": t,
		"active_user": loginUserInfo,
	})
}

func NewAuthController(keyString []byte, usersService *users_service.UsersService) *AuthController {
	return &AuthController{
		keyString: keyString,
		UsersService: usersService,
	}
}
