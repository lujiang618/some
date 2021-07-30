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

func (m *WealthCostCategory) GetCategoriesByIds(ids []int) (*[]response.CostCategoryObject, error) {
	categories := make([]response.CostCategoryObject, 0)
	err := m.Db().Where("id in (?)", ids).Find(&categories).Error

	return &categories, err
}

func (m *WealthCostCategory) GetCategoryMap(ids []int) (map[int]string, error) {
	result := make(map[int]string)

	categories, err := m.GetCategoriesByIds(ids)
	if err != nil {
		return result, err
	}

	for _, category := range *categories {
		result[category.Id] = category.Name
	}

	return result, nil
}
