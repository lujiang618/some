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

type AnalyseFilter struct {
	BaseFilter
	service *service.AnalyseService
}

func NewAnalyseFilter() *AnalyseFilter {
	return &AnalyseFilter{
		service: service.NewAnalyseService(),
	}
}

func (f AnalyseFilter) GetAll(c *gin.Context) (*response.Stat, *api.Error) {
	params := &request.AnalyseParamsAll{}

	if err := c.ShouldBindJSON(params); err != nil {
		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	spew.Dump(params)

	// 调用service对应的方法
	data, err := f.service.GetAll(params)

	return data, err
}

func (f AnalyseFilter) GetCharts(c *gin.Context) (*response.Charts, *api.Error) {
	params := &request.AnalyseParamsCharts{}

	if err := c.ShouldBindJSON(params); err != nil {
		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	spew.Dump(params)

	// 调用service对应的方法
	data, err := f.service.GetCharts(params)

	return data, err
}
