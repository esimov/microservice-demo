package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/esimov/xm/app/models"
	"github.com/esimov/xm/app/response"
	"github.com/esimov/xm/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCompanies(db *gorm.DB, c *config.Config, ctx *gin.Context) {
	company := models.Company{}

	companies, err := company.FindAll(db)
	if err != nil {
		response.Error(ctx.Writer, http.StatusInternalServerError, err)
		return
	}
	response.Status(ctx.Writer, http.StatusOK, companies)
}

func GetCompany(db *gorm.DB, c *config.Config, ctx *gin.Context) {
	company := models.Company{}

	cid, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.Error(ctx.Writer, http.StatusBadRequest, err)
		return
	}
	companies, err := company.FindById(db, cid)
	if err != nil {
		response.Error(ctx.Writer, http.StatusInternalServerError, err)
		return
	}
	response.Status(ctx.Writer, http.StatusOK, companies)
}

func CreateCompany(db *gorm.DB, c *config.Config, ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	company := models.Company{}
	err = json.Unmarshal(body, &company)
	if err != nil {
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}

	err = company.Save(db)
	if err != nil {
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	response.Status(ctx.Writer, http.StatusOK, company)
}

func UpdateCompany(db *gorm.DB, c *config.Config, ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	company := models.Company{}
	err = json.Unmarshal(body, &company)
	if err != nil {
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}

	cid, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.Error(ctx.Writer, http.StatusBadRequest, err)
		return
	}
	err = company.Update(db, cid)
	if err != nil {
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	response.Status(ctx.Writer, http.StatusOK, company)
}

func DeleteCompany(db *gorm.DB, c *config.Config, ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	company := models.Company{}
	err = json.Unmarshal(body, &company)
	if err != nil {
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}

	cid, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.Error(ctx.Writer, http.StatusBadRequest, err)
		return
	}
	rows, err := company.Delete(db, cid)
	if err != nil {
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	response.Status(ctx.Writer, http.StatusOK, rows)
}
