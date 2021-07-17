package models

type WealthExpendStatisticsWeak struct {
	DetailColumns []string
	BaseModel
}

func NewWealthExpendStatisticsWeak() *WealthExpendStatisticsWeak {
	statisticsWeak := &WealthExpendStatisticsWeak{}

	statisticsWeak.SetTableName("wealth_expend_statistics_weeks")

	return statisticsWeak
}

func (m *WealthExpendStatisticsWeak) GetList(userId int) (*[]WealthExpendStatisticsWeak, error) {
	statisticsWeek := make([]WealthExpendStatisticsWeak, 0)
	err := m.Db().Where("user_id = ?", userId).Find(&statisticsWeek).Error

	return &statisticsWeek, err
}
