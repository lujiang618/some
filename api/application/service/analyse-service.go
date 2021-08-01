package service

import (
	"some/api/application/models"
	"some/api/application/objects/request"
	"some/api/application/objects/response"
	"some/api/pkg/api"
	"some/api/pkg/api/code"
	"some/api/pkg/utils"
)

type AnalyseService struct {
}

func NewAnalyseService() *AnalyseService {
	return &AnalyseService{}
}

func (s *AnalyseService) GetAll(params *request.AnalyseParamsAll) (*response.Stat, *api.Error) {
	statisticsMonth := models.NewWealthStatisticsMonth()
	// 获取当前月支出
	currentMonth := utils.GetCurrentYearMonth()
	currentMonthStat, err := statisticsMonth.GetMonthTotal(params.UserId, currentMonth)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	// 获取上一月至此
	lastMonth := utils.GetLastYearMonth()
	lastMonthStat, err := statisticsMonth.GetMonthTotal(params.UserId, lastMonth)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	// 获取当前周支出
	currentWeekStart, currentWeekEnd := utils.GetCurrentWeekScope()

	statisticsWeek := models.NewWealthStatisticsWeak()
	currentWeekStat, err := statisticsWeek.GetTotal(params.UserId, currentWeekStart, currentWeekEnd)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	// 获取上一周支出
	lastWeekStart, lastWeekEnd := utils.GetLastWeekScope()
	lastWeekStat, err := statisticsWeek.GetTotal(params.UserId, lastWeekStart, lastWeekEnd)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	stat := &response.Stat{
		CurrentMonth: currentMonthStat,
		LastMonth:    lastMonthStat,
		CurrentWeek:  currentWeekStat,
		LastWeek:     lastWeekStat,
	}

	return stat, nil
}
