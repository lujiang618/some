package filter

import (
	"some/api/application/models"
	"some/api/application/objects/request"
	"some/api/application/service"
	"some/api/pkg/api"
	"some/api/pkg/api/code"

	"github.com/gin-gonic/gin"
)

type ExpendDetails struct {
	BaseFilter
	service *service.WealthExpend
}

func NewExpendDetails() *ExpendDetails {
	return &ExpendDetails{
		service: service.NewWealthExpend(),
	}
}

func (f ExpendDetails) GetList(c *gin.Context) (*[]models.ExpendDetailObject, *api.Error) {
	params := &request.ExpendParamList{}

	if err := c.ShouldBindUri(params); err != nil {

		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	// 调用service对应的方法
	data, err := f.service.GetList(params)

	return data, err
}

func (f ExpendDetails) GetDetail(c *gin.Context) (*models.ExpendDetailObject, *api.Error) {
	params := &request.ExpendParamDetail{}

	if err := c.ShouldBindUri(params); err != nil {

		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	// 调用service对应的方法
	data, err := f.service.GetDetail(params)

	return data, err
}

func (f ExpendDetails) Create(c *gin.Context) *api.Error {
	params := &request.ExpendParamCreate{}

	if err := c.ShouldBindUri(params); err != nil {

		return api.NewError(code.ErrorRequest, err.Error())
	}

	// 调用service对应的方法
	err := f.service.Create(params)

	return err
}
