package models

type WealthExpendCategory struct {
	DetailColumns []string
	BaseModel
}

func NewWealthExpendCategory() *WealthExpendCategory {
	category := &WealthExpendCategory{}

	category.SetTableName("wealth_expend_details")

	return category
}

func (m *WealthExpendCategory) GetList(userId int) (*[]WealthExpendCategory, error) {
	categories := make([]WealthExpendCategory, 0)
	err := m.Db().Where("user_id = ?", userId).Find(&categories).Error

	return &categories, err
}
