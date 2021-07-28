package models

import "some/api/application/objects/response"

type WealthCostCategory struct {
	DetailColumns []string
	BaseModel
}

func NewWealthCostCategory() *WealthCostCategory {
	category := &WealthCostCategory{}

	category.SetTableName("wealth_cost_categories")

	return category
}

func (m *WealthCostCategory) GetList(userId int) (*[]response.CostCategoryObject, int, error) {
	categories := make([]response.CostCategoryObject, 0)
	err := m.Db().Where("user_id = ?", userId).Find(&categories).Error

	return &categories, 100, err
}
