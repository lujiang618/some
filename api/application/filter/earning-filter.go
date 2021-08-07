package filter

import (
	"some/api/application/objects/request"
	"some/api/application/objects/response"
	"some/api/application/service"
	"some/api/pkg/api"
	"some/api/pkg/api/code"

	"github.com/gin-gonic/gin"
)

type EarningFilter struct {
	BaseFilter
	service *service.EarningService
}

func NewEarningFilter() *EarningFilter {
	return &EarningFilter{
		service: service.NewEarningService(),
	}
}

func (f EarningFilter) GetList(c *gin.Context) (*response.CostListObject, *api.Error) {
	params := &request.CostParamList{}

	if err := c.ShouldBindQuery(params); err != nil {
		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	// 调用service对应的方法
	data, err := f.service.GetList(params)

	return data, err
}

func (f EarningFilter) Create(c *gin.Context) *api.Error {
	params := &request.ParamEarningCreate{}

	if err := c.ShouldBindJSON(params); err != nil {
		return api.NewError(code.ErrorRequest, err.Error())
	}

	// 调用service对应的方法
	err := f.service.Create(params)

	return err
}

func (f EarningFilter) Update(c *gin.Context) *api.Error {
	params := &request.ParamEarningUpdate{}

	if err := c.ShouldBindJSON(params); err != nil {

		return api.NewError(code.ErrorRequest, err.Error())
	}

	// 调用service对应的方法
	err := f.service.Update(params)

	return err
}
