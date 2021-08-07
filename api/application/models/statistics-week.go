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

func (m *WealthStatisticsWeak) GetTotal(userId uint64, startData, endDate string) (string, error) {

	var total response.StatTotal
	err := m.Db().Select("round(sum(total)) total").
		Where("user_id = ? and start_date=? and end_date=?", userId, startData, endDate).
		First(&total).Error

	spew.Dump(total)

	return total.Total, err
}
