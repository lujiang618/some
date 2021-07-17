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

func (s *CostService) GetList(params *request.CostParamList) (*[]response.CostDetailObject, *api.Error) {

	data, err := models.NewCostDetail().GetList(int(params.UserId))

	return data, api.NewError(code.ErrorDatabase, err.Error())
}

func (s *CostService) GetDetail(params *request.CostParamDetail) (*response.CostDetailObject, *api.Error) {

	data, err := models.NewCostDetail().GetDetail(params.Id)

	return data, api.NewError(code.ErrorDatabase, err.Error())
}

func (s *CostService) Create(params *request.CostParamCreate) *api.Error {

	err := models.NewCostDetail().Create(params)

	return api.NewError(code.ErrorDatabase, err.Error())
}
