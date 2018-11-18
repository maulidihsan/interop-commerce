package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"encoding/json"
	"gopkg.in/go-playground/validator.v9"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/maulidihsan/flashdeal-webservice/pkg/user/usecase"
	"github.com/maulidihsan/flashdeal-webservice/pkg/models"
	"github.com/maulidihsan/flashdeal-webservice/config"
)

type UserController struct{
	_user usecase.UserUsecase
	validate *validator.Validate
}

func NewUserController(u usecase.UserUsecase, v *validator.Validate) *UserController {
	return &UserController{
		_user: u,
		validate: v,
	}
}

func (u UserController) Auth(c *gin.Context) {
	config := config.GetConfig()
	var creds models.Credential
	c.Bind(&creds)
	err := u.validate.Struct(&creds)
	if(err != nil) {
		c.JSON(401, gin.H{"message": "Field validation error", "error": err})
		c.Abort()
		return
	}
	user, err := u._user.Login(&creds)
	if(err != nil) {
		c.JSON(500, gin.H{"message": "Error to retrieve user", "error": err})
		c.Abort()
		return
	}
	payload, err := json.Marshal(user)
    if err != nil {
        c.JSON(500, gin.H{"message": "Error to marshall user data", "error": err})
		c.Abort()
		return
    }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"data": string(payload)})
	tokenString, err := token.SignedString([]byte(config.GetString("jwt.secret")))
	if err != nil {
        c.JSON(500, gin.H{"message": "Error to marshall user data", "error": err})
		c.Abort()
		return
    }
	c.JSON(http.StatusOK, gin.H{"success": true, "token": tokenString})
	return
}

func (u UserController) Register(c *gin.Context) {
	var newUser models.User
	c.Bind(&newUser)
	err := u.validate.Struct(&newUser)
	if(err != nil) {
		var errors []models.ValidationErr
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, models.ValidationErr{
				Tag: err.Tag(),
				Value: err.Value(),
			})
		}
		c.JSON(401, gin.H{"message": "Field validation error", "error": errors})
		c.Abort()
		return
	}
	err = u._user.AddUser(&newUser)
	if(err != nil) {
		c.JSON(500, gin.H{"message": "Error to retrieve user", "error": err})
		c.Abort()
		return
	}
	res := models.Response{
		Message: "User Added",
		Success: true,
		Code: 200,
	}
	c.JSON(http.StatusOK, res)
	return
}

func (u UserController) Get(c *gin.Context) {
	payload := c.MustGet("user").(string)
	var user models.User
	err := json.Unmarshal([]byte(payload), &user)
	if(err != nil) {
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, user)
	return
}