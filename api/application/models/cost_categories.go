package models

type WealthCostCategory struct {
	DetailColumns []string
	BaseModel
}

func NewWealthCostCategory() *WealthCostCategory {
	category := &WealthCostCategory{}

	category.SetTableName("wealth_cost_details")

	return category
}

func (m *WealthCostCategory) GetList(userId int) (*[]WealthCostCategory, error) {
	categories := make([]WealthCostCategory, 0)
	err := m.Db().Where("user_id = ?", userId).Find(&categories).Error

	return &categories, err
}
