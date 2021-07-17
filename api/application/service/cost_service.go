package service

import (
	"some/api/application/models"
	"some/api/application/objects/request"
	"some/api/application/objects/response"
	"some/api/pkg/api"
	"some/api/pkg/api/code"
)

type CostService struct {
}

func NewCostService() *CostService {
	return &CostService{}
}

func (s *CostService) GetList(params *request.CostParamList) (*response.CostListObject, *api.Error) {

	data, count, err := models.NewWealthCostDetail().GetList(int(params.UserId))
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	result := &response.CostListObject{
		PageSize:   10,
		PageNo:     1,
		TotalCount: count,
		TotalPage:  10,
		Data:       data,
	}
	return result, nil
}

func (s *CostService) GetDetail(params *request.CostParamDetail) (*response.CostDetailObject, *api.Error) {

	data, err := models.NewWealthCostDetail().GetDetail(params.Id)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	return data, nil
}

func (s *CostService) Create(params *request.CostParamCreate) *api.Error {

	err := models.NewWealthCostDetail().Create(params)
	if err != nil {
		return api.NewError(code.ErrorDatabase, err.Error())
	}

	return nil
}

func (s *CostService) GetCategoryList(params *request.CostParamList) (*response.CostCategoryListObject, *api.Error) {

	data, count, err := models.NewWealthCostCategory().GetList(int(params.UserId))
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	result := &response.CostCategoryListObject{
		PageSize:   10,
		PageNo:     1,
		TotalCount: count,
		TotalPage:  10,
		Data:       data,
	}
	return result, nil
}
