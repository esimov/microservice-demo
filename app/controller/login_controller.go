package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/esimov/xm/app/models"
	"github.com/esimov/xm/auth"
	"github.com/esimov/xm/config"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, c *config.Config, ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}

	err = user.BeforeSave()
	if err != nil {
		Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	err = user.Save(db)
	if err != nil {
		Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	Response(ctx.Writer, http.StatusOK, user)
}

func Login(db *gorm.DB, c *config.Config, ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := SignIn(db, c, user.Email, user.Password)
	if err != nil {
		Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	Response(ctx.Writer, http.StatusOK, token)
}

func SignIn(db *gorm.DB, c *config.Config, email, password string) (string, error) {
	var err error
	user := models.User{}

	err = user.FindByEmail(db, email)
	if err != nil {
		return "", err
	}
	err = user.VerifyPassword(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(c, user.ID)
}
