package controllers

import (
	"net/http"
	"some/api/application/filter"
	"some/api/pkg/api"
	"some/api/pkg/api/code"

	"github.com/gin-gonic/gin"
)

type CostController struct {
	BaseController
	paramFilter *filter.CostFilter
}

func NewCostController() *CostController {
	return &CostController{
		paramFilter: filter.NewCostFilter(),
	}
}

func (ctl *CostController) GetList(c *gin.Context) {
	data, err := ctl.paramFilter.GetList(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, code.Success, data)
}

func (ctl *CostController) GetDetail(c *gin.Context) {
	detail, err := ctl.paramFilter.GetDetail(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, code.Success, detail)
}

func (ctl *CostController) Create(c *gin.Context) {
	if err := ctl.paramFilter.Create(c); err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, code.Success, "")
}

func (ctl *CostController) Update(c *gin.Context) {
	if err := ctl.paramFilter.Update(c); err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, code.Success, "")
}

func (ctl *CostController) GetCostCategories(c *gin.Context) {
	data, err := ctl.paramFilter.GetCategoryList(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, code.Success, data)
}
