package models

type WealthStatisticsMonth struct {
	DetailColumns []string
	BaseModel
}

func NewWealthStatisticsMonth() *WealthStatisticsMonth {
	statisticsMonth := &WealthStatisticsMonth{}

	statisticsMonth.SetTableName("wealth_cost_statistics_months")

	return statisticsMonth
}

func (m *WealthStatisticsMonth) GetList(userId int) (*[]WealthStatisticsMonth, error) {
	statisticsMonth := make([]WealthStatisticsMonth, 0)
	err := m.Db().Where("user_id = ?", userId).Find(&statisticsMonth).Error

	return &statisticsMonth, err
}
