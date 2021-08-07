package filter

import (
	"some/api/application/objects/request"
	"some/api/application/objects/response"
	"some/api/application/service"
	"some/api/pkg/api"
	"some/api/pkg/api/code"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

type CostFilter struct {
	BaseFilter
	service *service.CostService
}

func NewCostFilter() *CostFilter {
	return &CostFilter{
		service: service.NewCostService(),
	}
}

func (f CostFilter) GetList(c *gin.Context) (*response.CostListObject, *api.Error) {
	params := &request.CostParamList{}

	if err := c.ShouldBindQuery(params); err != nil {
		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	// 调用service对应的方法
	data, err := f.service.GetList(params)

	return data, err
}

func (f CostFilter) GetDetail(c *gin.Context) (*response.CostDetailObject, *api.Error) {
	params := &request.CostParamDetail{}

	if err := c.ShouldBindUri(params); err != nil {

		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	// 调用service对应的方法
	data, err := f.service.GetDetail(params)

	return data, err
}

func (f CostFilter) Create(c *gin.Context) *api.Error {
	params := &request.CostParamCreate{}

	if err := c.ShouldBindJSON(params); err != nil {
		return api.NewError(code.ErrorRequest, err.Error())
	}

	// 调用service对应的方法
	err := f.service.Create(params)

	return err
}

func (f CostFilter) Update(c *gin.Context) *api.Error {
	params := &request.CostParamUpdate{}

	if err := c.ShouldBindJSON(params); err != nil {

		return api.NewError(code.ErrorRequest, err.Error())
	}

	// 调用service对应的方法
	err := f.service.Update(params)

	return err
}

func (f CostFilter) GetCategoryList(c *gin.Context) (*response.CostCategoryListObject, *api.Error) {
	params := &request.CostParamList{}

	if err := c.ShouldBindQuery(params); err != nil {
		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	spew.Dump(params)

	// 调用service对应的方法
	data, err := f.service.GetCategoryList(params)

	return data, err
}
