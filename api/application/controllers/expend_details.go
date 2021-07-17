package controllers

import (
	"net/http"
	"some/api/application/filter"
	"some/api/pkg/api"
	"some/api/pkg/api/code"

	"github.com/gin-gonic/gin"
)

type ExpendDetails struct {
	BaseController
	paramFilter *filter.ExpendDetails
}

func (ctl *ExpendDetails) GetList(c *gin.Context) {
	data, err := ctl.paramFilter.GetList(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, code.Success, data)
}

func (ctl *ExpendDetails) GetDetail(c *gin.Context) {
	detail, err := ctl.paramFilter.GetDetail(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, code.Success, detail)
}

func (ctl *ExpendDetails) Create(c *gin.Context) {
	if err := ctl.paramFilter.Create(c); err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, code.Success, "")
}
