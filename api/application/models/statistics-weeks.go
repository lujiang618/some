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

func (m *WealthStatisticsWeak) GetTotal(userId uint64, startData, endDate string) (string, error) {
	total := ""
	err := m.Db().Select("sum(amount) total").Where("user_id = ? and start_date=? and end_date=?", userId, startData, endDate).Find(&total).Error

	return total, err
}
