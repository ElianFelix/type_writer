package local_middleware

import (
	"net/http"
	"type_writer_api/structures"

	"github.com/casbin/casbin/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func contextUserGetter(ctx echo.Context) (string, error) {
	token, err := echo.ContextGet[*jwt.Token](ctx, "user")
	if err != nil {
		return "", err
	}	
	claims := token.Claims.(*structures.JwtCustomClaims)
	return claims.UserType, nil
}

func CasbinMiddleware(enforcer *casbin.Enforcer) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			user_type, err := contextUserGetter(ctx)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err)
			}
			if pass, err := enforcer.Enforce(user_type, ctx.Request().URL.Path, ctx.Request().Method); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			} else if !pass {
				return echo.NewHTTPError(http.StatusForbidden, "access denied")
			}
			return next(ctx)
		}
	}
}


