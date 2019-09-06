package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/xdean/goex/xecho"
	"github.com/xdean/miniboardgame/go/server/config"
	_const "github.com/xdean/miniboardgame/go/server/const"
	"github.com/xdean/miniboardgame/go/server/model"
	"github.com/xdean/miniboardgame/go/server/web/handler/openid"
	"github.com/xdean/miniboardgame/go/server/web/middleware"
	"net/http"
	"time"
)

func SignUp(c echo.Context) error {
	type Param struct {
		Username string `json:"username" form:"username" query:"username" validate:"required,regexp=USERNAME"`
		Password string `json:"password" form:"password" query:"password" validate:"required,regexp=PASSWORD"`
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
		if t, err := user.GenerateToken(config.Global.Security.Key); err == nil {
			c.SetCookie(generateTokenCookie(t))
			return c.JSON(http.StatusCreated, xecho.J{
				"message": "Sign up success",
				"token":   t,
			})
		} else {
			return err
		}
	} else {
		return err
	}
}

type LoginParam struct {
	Type     string `json:"type" form:"type" query:"type"`
	Username string `json:"username" form:"username" query:"username" validate:"regexp=USERNAME"`
	Password string `json:"password" form:"password" query:"password"`
	Provider string `json:"provider" form:"provider" query:"provider"`
	Token    string `json:"token" form:"token" query:"token"`
}

func Login(c echo.Context) error {
	param := new(LoginParam)
	if err := c.Bind(param); err != nil {
		return err
	}
	if err := c.Validate(param); err != nil {
		return err
	}
	switch param.Type {
	case "openid":
		return LoginOpenid(c, *param)
	case "password":
		fallthrough
	case "":
		return LoginPassword(c, *param)
	default:
		return echo.NewHTTPError(http.StatusBadRequest, "Unknown login type: "+param.Type)
	}
}

func LoginPassword(c echo.Context, param LoginParam) error {
	if param.Username == "" || param.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Username and password are required")
	}
	user := new(model.User)
	if err := user.FindByUsername(GetDB(c), param.Username); err == nil {
		if user.MatchPassword(param.Password) {
			if t, err := user.GenerateToken(config.Global.Security.Key); err == nil {
				c.SetCookie(generateTokenCookie(t))
				return c.JSON(http.StatusOK, xecho.J{
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

func LoginOpenid(c echo.Context, param LoginParam) error {
	if param.Provider == "" || param.Token == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Provider and token are required")
	}
	oid, err := openid.Get(param.Provider, param.Token)
	if err == nil {
		user := &model.User{
			Username: oid + "@" + param.Provider,
			Password: oid,
			Roles:    []model.Role{{Name: _const.ROLE_USER}},
		}
		db := GetDB(c)
		yes, err := model.UserExistByUsername(db, user.Username)
		if err != nil {
			return err
		}
		if yes {
			if err := user.FindByUsername(db, user.Username); err != nil {
				return err
			}
		} else {
			if err := user.CreateAccount(db); err != nil {
				return err
			}
		}
		if t, err := user.GenerateToken(config.Global.Security.Key); err == nil {
			c.SetCookie(generateTokenCookie(t))
			return c.JSON(http.StatusOK, xecho.J{
				"message": "Login success",
				"token":   t,
			})
		} else {
			return err
		}
	} else {
		return err
	}
}

func Logout(c echo.Context) error {
	// TODO use refresh token
	c.SetCookie(&http.Cookie{
		Path:    "/",
		Name:    middleware.JwtKey,
		Expires: time.Now(),
	})
	return c.JSON(http.StatusOK, xecho.M("Logout success"))
}

func generateTokenCookie(token string) *http.Cookie {
	return &http.Cookie{
		Path:    "/",
		Name:    middleware.JwtKey,
		Value:   token,
		Expires: time.Now().Add(2 * time.Hour),
	}
}
