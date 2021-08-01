package models

import (
	"some/api/application/objects/response"
)

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
	var total response.StatTotal
	err := m.Db().
		Select("sum(total) total").Where("user_id = ? and `year_month` = ?", userId, yearMonth).
		First(&total).Error

	return total.Total, err
}
