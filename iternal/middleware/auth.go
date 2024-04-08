package middleware

import (
	"awesomeProject/iternal/models"
	"awesomeProject/iternal/repository"
	"awesomeProject/iternal/tools"
	"errors"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var identityKey = "id"
var userRepo = repository.NewUserRepo()

// User demo
type User struct {
	ID string
}

func Passport() *jwt.GinJWTMiddleware {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		// Realm имя
		Realm: "user",
		// Key секретный ключ
		Key:         []byte("JWTSECRET"),
		Timeout:     time.Hour * 4,
		MaxRefresh:  time.Hour * 24,
		IdentityKey: identityKey,

		// PayloadFunc просто переписать тк сложная хуйня
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},

		// IdentityHandler просто переписать тк сложная хуйня
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				ID: claims[identityKey].(string),
			}
		},

		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals models.Login
			// ShouldBind заполняет loginVals если поля совпадают
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			userLogin := loginVals.Login
			password := loginVals.Password

			result, err := userRepo.GetByKey("login", userLogin)

			if err != nil {
				return nil, errors.New("не верный логин или пароль")
			}

			if tools.CheckPasswordHash(password, result.Password) {
				return &User{
					ID: result.ID.String(),
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": "вы не авторизированны",
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
		LoginResponse: func(c *gin.Context, i int, s string, t time.Time) {
			value, _ := Passport().ParseTokenString(s)
			id := jwt.ExtractClaimsFromToken(value)["id"]
			result, err := userRepo.GetByKey("id", id.(string))

			if err != nil {
				tools.CreateError(http.StatusUnauthorized, err, c)
				return
			}

			c.Header("token", value.Raw)
			c.JSON(http.StatusOK, result)
		},
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	return authMiddleware
}
