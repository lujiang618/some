package models

type WealthExpendStatisticsMonth struct {
	DetailColumns []string
	BaseModel
}

func NewWealthExpendStatisticsMonth() *WealthExpendStatisticsMonth {
	statisticsMonth := &WealthExpendStatisticsMonth{}

	statisticsMonth.SetTableName("wealth_expend_statistics_months")

	return statisticsMonth
}

func (m *WealthExpendStatisticsMonth) GetList(userId int) (*[]WealthExpendStatisticsMonth, error) {
	statisticsMonth := make([]WealthExpendStatisticsMonth, 0)
	err := m.Db().Where("user_id = ?", userId).Find(&statisticsMonth).Error

	return &statisticsMonth, err
}
