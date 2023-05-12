package controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/esimov/xm/app/models"
	"github.com/esimov/xm/app/response"
	"github.com/esimov/xm/auth"
	"github.com/esimov/xm/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB, c *config.Config, ctx *gin.Context) {
	user := models.User{}

	users, err := user.FindAll(db)
	if err != nil {
		response.Error(ctx.Writer, http.StatusInternalServerError, err)
		return
	}
	response.Status(ctx.Writer, http.StatusOK, users)
}

func CreateUser(db *gorm.DB, c *config.Config, ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}

	err = user.BeforeSave()
	if err != nil {
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	err = user.Save(db)
	if err != nil {
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	response.Status(ctx.Writer, http.StatusOK, user)
}

func UpdateUser(db *gorm.DB, c *config.Config, ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}

	uid, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.Error(ctx.Writer, http.StatusBadRequest, err)
		return
	}
	tokenID, err := auth.ExtractTokenID(c, ctx.Request)
	if err != nil {
		response.Error(ctx.Writer, http.StatusUnauthorized, errors.New("Unauthorized user"))
		return
	}
	if tokenID != uint32(uid) {
		response.Error(ctx.Writer, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	err = user.BeforeSave()
	if err != nil {
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	err = user.Update(db)
	if err != nil {
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	response.Status(ctx.Writer, http.StatusOK, user)
}
