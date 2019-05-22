package handler

import (
	"errors"
	"github.com/XDean/MiniBoardgame/config"
	_const "github.com/XDean/MiniBoardgame/const"
	"github.com/XDean/MiniBoardgame/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

func SignUp(c echo.Context) error {
	type Param struct {
		Username string `json:"username" validate:"required,regexp=USERNAME"`
		Password string `json:"password" validate:"required,regexp=PASSWORD"`
	}
	param := new(Param)
	if err := c.Bind(param); err != nil {
		return err
	}
	if err := c.Validate(param); err != nil {
		return err
	}
	user := &model.User{
		Username: param.Username,
		Password: param.Password,
		Roles:    []model.Role{{Name: _const.ROLE_USER}},
	}
	if err := user.CreateAccount(GetDB(c)); err == nil {
		return c.JSON(http.StatusCreated, M("Sign up success"))
	} else {
		return err
	}
}

func Login(c echo.Context) error {
	type Param struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	param := new(Param)
	if err := c.Bind(param); err != nil {
		return err
	}
	if err := c.Validate(param); err != nil {
		return err
	}
	user := new(model.User)
	if err := user.FindByUsername(GetDB(c), param.Username); err == nil {
		if user.MatchPassword(param.Password) {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
				User: *user,
				StandardClaims: jwt.StandardClaims{
					Subject:   "access token",
					ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
				},
			})
			if t, err := token.SignedString([]byte(config.Global.Security.Key)); err == nil {
				return c.JSON(http.StatusOK, J{
					"message": "Login success",
					"token":   t,
				})
			} else {
				return err
			}
		}
	}
	return echo.NewHTTPError(http.StatusUnauthorized, "Bad Credentials")
}

func GetCurrentUser(c echo.Context) (*model.User, error) {
	if user, ok := c.Get(_const.USER).(*model.User); ok {
		return user, nil
	}
	return nil, errors.New("not authorized")
}

func JwtAuthenticateConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		SigningKey: []byte(config.Global.Security.Key),
		Claims:     Claims{},
		SuccessHandler: func(echo echo.Context) {
			token := echo.Get(middleware.DefaultJWTConfig.ContextKey).(jwt.Token)
			claims := token.Claims.(Claims)
			echo.Set(_const.USER, claims.User)
		},
	}
}

type Claims struct {
	jwt.StandardClaims
	User model.User
}

func AdminAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		if user, ok := context.Get(_const.USER).(model.User); ok {
			for _, role := range user.Roles {
				if _const.ROLE_ADMIN == role.Name {
					return next(context)
				}
			}
		}
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "not authorized or not admin",
		}
	}
}
