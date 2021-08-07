package controllers

import (
	"net/http"
	"some/api/application/filter"
	"some/api/pkg/api"
	"some/api/pkg/api/code"

	"github.com/gin-gonic/gin"
)

type EarningController struct {
	BaseController
	paramFilter *filter.EarningFilter
}

func NewEarningController() *EarningController {
	return &EarningController{
		paramFilter: filter.NewEarningFilter(),
	}
}

func (ctl *EarningController) GetList(c *gin.Context) {
	data, err := ctl.paramFilter.GetList(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, code.Success, data)
}

func (ctl *EarningController) Create(c *gin.Context) {
	if err := ctl.paramFilter.Create(c); err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, code.Success, "")
}

func (ctl *EarningController) Update(c *gin.Context) {
	if err := ctl.paramFilter.Update(c); err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, code.Success, "")
}
