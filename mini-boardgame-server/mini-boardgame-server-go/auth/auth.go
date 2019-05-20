package auth

import (
	"errors"
	"github.com/XDean/MiniBoardgame/config"
	_const "github.com/XDean/MiniBoardgame/const"
	"github.com/XDean/MiniBoardgame/db"
	"github.com/XDean/MiniBoardgame/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func AuthenticatePassword(username string, password string) error {
	var user model.User
	if err := db.DB.Where("username = ?", username).Find(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("Username not exist")
		}
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("Password not correct")
	}
	return nil
}

func JwtAuthenticateConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		SigningKey: config.Global.Security.Key,
		Claims:     Claims{},
		SuccessHandler: func(echo echo.Context) {
			token := echo.Get(middleware.DefaultJWTConfig.ContextKey).(jwt.Token)
			claims := token.Claims.(Claims)
			echo.Set(_const.USERID, claims.UserID)
			echo.Set(_const.USERNAME, claims.Username)
			echo.Set(_const.ROLES, claims.Roles)
		},
	}
}

type Claims struct {
	jwt.StandardClaims
	UserID   uint
	Username string
	Roles    []string
}

func AdminAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		if roles, ok := context.Get(_const.ROLES).(_const.ROLES_TYPE); ok {
			for _, role := range roles {
				if _const.ROLE_ADMIN == role {
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
