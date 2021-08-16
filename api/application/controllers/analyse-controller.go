package controllers

import (
	"net/http"
	"some/api/application/filter"
	"some/api/pkg/api"
	"some/api/pkg/api/code"

	"github.com/gin-gonic/gin"
)

type AnalyseController struct {
	BaseController
	paramFilter *filter.AnalyseFilter
}

func NewAnalyseController() *AnalyseController {
	return &AnalyseController{
		paramFilter: filter.NewAnalyseFilter(),
	}
}

func (ctl *AnalyseController) GetAll(c *gin.Context) {
	data, err := ctl.paramFilter.GetAll(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, code.Success, data)
}

func (ctl *AnalyseController) GetCharts(c *gin.Context) {
	data, err := ctl.paramFilter.GetCharts(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, code.Success, data)
}
