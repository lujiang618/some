package models

import (
	"some/api/application/objects/request"
	"some/api/application/objects/response"
)

type CostDetail struct {
	DetailColumns []string
	BaseModel
}

func NewCostDetail() *CostDetail {
	detail := &CostDetail{}

	detail.SetTableName("wealth_Cost_statistics_months")

	return detail
}

func (m *CostDetail) GetList(userId int) (*[]response.CostDetailObject, error) {
	details := make([]response.CostDetailObject, 0)
	err := m.Db().Where("user_id = ?", userId).Find(&details).Error

	return &details, err
}

func (m *CostDetail) GetDetail(id int) (*response.CostDetailObject, error) {
	detail := &response.CostDetailObject{}
	err := m.Db().Where("id = ?", id).First(detail).Error

	return detail, err
}

func (m *CostDetail) Create(params *request.CostParamCreate) error {
	detail := &response.CostDetailObject{
		UserId:      params.UserId,
		Category:    params.Category,
		OccurDate:   params.OccurDate,
		Content:     params.Content,
		Amount:      params.Amount,
		Description: params.Description,
	}

	return m.Db().Create(detail).Error
}
