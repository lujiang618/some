package models

import (
	"some/api/application/objects/response"

	"github.com/davecgh/go-spew/spew"
)

type WealthStatisticsWeak struct {
	DetailColumns []string
	BaseModel
}

func NewWealthStatisticsWeak() *WealthStatisticsWeak {
	statisticsWeak := &WealthStatisticsWeak{}

	statisticsWeak.SetTableName("wealth_cost_statistics_weeks")

	return statisticsWeak
}

func (m *WealthStatisticsWeak) GetTotalWeeks(userId uint64) ([]response.TotalWeek, error) {
	var totalMonths []response.TotalWeek
	err := m.Db().Select("date_format(start_date,'%m%d') as `week`, sum(total) as total").
		Where("user_id = ? and year=2021", userId).
		Group("year,start_date").
		Find(&totalMonths).Error

	return totalMonths, err
}

func (m *WealthStatisticsWeak) GetTotal(userId uint64, startData, endDate string) (string, error) {

	var total response.StatTotal
	err := m.Db().Select("round(sum(total)) total").
		Where("user_id = ? and start_date=? and end_date=?", userId, startData, endDate).
		First(&total).Error

	spew.Dump(total)

	return total.Total, err
}
