package models

type WealthStatisticsWeak struct {
	DetailColumns []string
	BaseModel
}

func NewWealthStatisticsWeak() *WealthStatisticsWeak {
	statisticsWeak := &WealthStatisticsWeak{}

	statisticsWeak.SetTableName("wealth_cost_statistics_weeks")

	return statisticsWeak
}

func (m *WealthStatisticsWeak) GetList(userId int) (*[]WealthStatisticsWeak, error) {
	statisticsWeek := make([]WealthStatisticsWeak, 0)
	err := m.Db().Where("user_id = ?", userId).Find(&statisticsWeek).Error

	return &statisticsWeek, err
}
