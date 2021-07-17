package service

import (
	"some/api/application/models"
	"some/api/application/objects/request"
	"some/api/pkg/api"
	"some/api/pkg/api/code"
)

type WealthExpend struct {
}

func NewWealthExpend() *WealthExpend {
	return &WealthExpend{}
}

func (s *WealthExpend) GetList(params *request.ExpendParamList) (*[]models.ExpendDetailObject, *api.Error) {

	data, err := models.NewExpendDetail().GetList(int(params.UserId))

	return data, api.NewError(code.ErrorDatabase, err.Error())
}

func (s *WealthExpend) GetDetail(params *request.ExpendParamDetail) (*models.ExpendDetailObject, *api.Error) {

	data, err := models.NewExpendDetail().GetDetail(params.Id)

	return data, api.NewError(code.ErrorDatabase, err.Error())
}

func (s *WealthExpend) Create(params *request.ExpendParamCreate) *api.Error {

	err := models.NewExpendDetail().Create(params)

	return api.NewError(code.ErrorDatabase, err.Error())
}
