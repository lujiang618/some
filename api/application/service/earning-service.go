package service

import (
	"some/api/application/models"
	"some/api/application/objects/request"
	"some/api/application/objects/response"
	"some/api/pkg/api"
	"some/api/pkg/api/code"
)

type EarningService struct {
	objEarning *models.WealthEarning
}

func NewEarningService() *EarningService {
	return &EarningService{
		objEarning: models.NewWWealthEarning(),
	}
}

func (s *EarningService) GetList(params *request.CostParamList) (*response.CostListObject, *api.Error) {

	data, count, err := s.objEarning.GetList(params)
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

func (s *EarningService) Create(params *request.ParamEarningCreate) *api.Error {

	err := s.objEarning.Create(params)
	if err != nil {
		return api.NewError(code.ErrorDatabase, err.Error())
	}

	return nil
}

func (s *EarningService) Update(params *request.ParamEarningUpdate) *api.Error {

	err := s.objEarning.Update(params)
	if err != nil {
		return api.NewError(code.ErrorDatabase, err.Error())
	}

	return nil
}
