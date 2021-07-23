package models

import (
	"some/api/application/objects/request"
	"some/api/application/objects/response"
)

type WealthCostDetail struct {
	DetailColumns []string
	BaseModel
}

func NewWealthCostDetail() *WealthCostDetail {
	detail := &WealthCostDetail{}

	detail.SetTableName("wealth_cost_details")

	return detail
}

func (m *WealthCostDetail) GetList(userId int) (*[]response.CostDetailObject, int, error) {
	details := make([]response.CostDetailObject, 0)
	err := m.Db().Where("user_id = ?", userId).Find(&details).Error

	return &details, 100, err
}

func (m *WealthCostDetail) GetDetail(id int) (*response.CostDetailObject, error) {
	detail := &response.CostDetailObject{}
	err := m.Db().Where("id = ?", id).First(detail).Error

	return detail, err
}

func (m *WealthCostDetail) Create(params *request.CostParamCreate) error {
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

func (m *WealthCostDetail) Update(params *request.CostParamUpdate) error {
	detail := &response.CostDetailObject{
		UserId:      params.UserId,
		Category:    params.Category,
		OccurDate:   params.OccurDate,
		Content:     params.Content,
		Amount:      params.Amount,
		Description: params.Description,
	}

	return m.Db().Updates(detail).Where("id = ?", params.Id).Error
}
