package service

import (
	"some/api/application/models"
	"some/api/application/objects/request"
	"some/api/application/objects/response"
	"some/api/pkg/api"
	"some/api/pkg/api/code"
)

/***
const types = [
  {
    id: 1,
    name: '活期存款'
  },
  {
    id: 2,
    name: '定期存款'
  },
  {
    id: 3,
    name: '余额宝'
  },
  {
    id: 4,
    name: '支付宝'
  },
  {
    id: 5,
    name: '微信'
  },
  {
    id: 6,
    name: '零钱通'
  }
]
**/

var types = map[int]string{
	1: "活期存款",
	2: "活期存款",
	3: "余额宝",
	4: "支付宝",
	5: "微信",
	6: "零钱通",
	7: "现金",
}

type EarningService struct {
	objEarning *models.WealthEarning
}

func NewEarningService() *EarningService {
	return &EarningService{
		objEarning: models.NewWWealthEarning(),
	}
}

func (s *EarningService) GetList(params *request.CostParamList) (*response.EarningListObject, *api.Error) {

	data, count, err := s.objEarning.GetList(params)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	s.ApplyTepes(data)
	result := &response.EarningListObject{
		PageSize:   10,
		PageNo:     1,
		TotalCount: count,
		TotalPage:  10,
		Data:       data,
	}
	return result, nil
}

func (s *EarningService) ApplyTepes(rows *[]response.EarningObject) {
	if len(*rows) == 0 {
		return
	}

	for k, row := range *rows {
		(*rows)[k].TypeName = types[row.Type]
	}
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
