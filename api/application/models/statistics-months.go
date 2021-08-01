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

func (m *WealthStatisticsMonth) GetMonthTotal(userId uint64, yearMonth string) (string, error) {
	total := ""
	err := m.Db().Select("sum(amount) total").Where("user_id = ? and year_month = ?", userId, yearMonth).First(&total).Error

	return total, err
}
