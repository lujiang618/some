package models

import (
	"some/api/application/objects/response"
	"some/api/pkg/utils"
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

func (m *WealthStatisticsMonth) GetYearTotal(userId uint64, year int) (string, error) {
	var total response.StatTotal
	err := m.Db().
		Select("sum(total) total").Where("user_id = ? and `year` = ?", userId, year).
		First(&total).Error

	return total.Total, err
}

func (m *WealthStatisticsMonth) GetAvgYear(userId uint64, year int) (string, error) {
	var avg response.StatAvg
	err := m.Db().Select("round(sum(total)/?) as avg", utils.GetLastMonthDayOfYear()).
		Where("user_id = ? and `year` = ? ", userId, year).
		First(&avg).Error

	return avg.Avg, err
}

func (m *WealthStatisticsMonth) GetAvgYearNoLoad(userId uint64, year int) (string, error) {
	var avg response.StatAvg
	err := m.Db().Select("round(sum(total)/?) as avg", utils.GetLastMonthDayOfYear()).
		Where("user_id = ? and `year` = ?  and category_id <> ?", userId, year, 11).
		First(&avg).Error

	return avg.Avg, err
}
