package service

import (
	"some/api/application/models"
	"some/api/application/objects/request"
	"some/api/application/objects/response"
	"some/api/pkg/api"
	"some/api/pkg/api/code"
	"some/api/pkg/utils"
)

/***
统计内容：
+ 当前月
+ 当前周
+ 上月
+ 上周
+ 本年
+ 本年日均
+ 本周日均
**/
type AnalyseService struct {
	objDetail *models.WealthCostDetail
	objMonth  *models.WealthStatisticsMonth
	objWeek   *models.WealthStatisticsWeak
}

func NewAnalyseService() *AnalyseService {
	return &AnalyseService{
		objDetail: models.NewWealthCostDetail(),
		objMonth:  models.NewWealthStatisticsMonth(),
		objWeek:   models.NewWealthStatisticsWeak(),
	}
}

func (s *AnalyseService) GetAll(params *request.AnalyseParamsAll) (*response.Stat, *api.Error) {
	// 获取当前月支出
	currentMonthStat, err := s.GetCurrentMonth(params)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	// 获取上一月支出
	lastMonthStat, err := s.GetLastMonth(params)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	// 获取当前周支出
	currentWeekStat, err := s.GetCurrentWeek(*params)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	// 获取上一周支出
	lastWeekStat, err := s.GetLastWeek(params)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	avgCurrentWeek, err := s.GetAvgCurrentWeek(params)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	avgCurrentWeekNoLoad, err := s.GetAvgCurrentWeekNoLoad(params)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	avgCurrentYear, err := s.GetAvgCurrentYear(params)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	avgCurrentYearNoLoad, err := s.GetAvgCurrentYearNoLoad(params)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	totalYear, err := s.GetTotalCurrentYear(params)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	stat := &response.Stat{
		CurrentMonth:         currentMonthStat,
		LastMonth:            lastMonthStat,
		CurrentWeek:          currentWeekStat,
		LastWeek:             lastWeekStat,
		AvgCurrentWeek:       avgCurrentWeek,
		AvgCurrentWeekNoLoad: avgCurrentWeekNoLoad,
		AvgCurrentYear:       avgCurrentYear,
		AvgCurrentYearNoLoad: avgCurrentYearNoLoad,
		TotalYear:            totalYear,
	}

	return stat, nil
}

func (s *AnalyseService) GetCurrentMonth(params *request.AnalyseParamsAll) (string, error) {
	return s.objDetail.GetCurrentMonth(params.UserId)
}

func (s *AnalyseService) GetLastMonth(params *request.AnalyseParamsAll) (string, error) {
	lastMonth := utils.GetLastYearMonth()
	return s.objMonth.GetMonthTotal(params.UserId, lastMonth)
}

func (s *AnalyseService) GetCurrentWeek(params request.AnalyseParamsAll) (string, error) {
	return s.objDetail.GetCurrentWeek(params.UserId)
}

func (s *AnalyseService) GetLastWeek(params *request.AnalyseParamsAll) (string, error) {
	currentWeekStart, currentWeekEnd := utils.GetLastWeekScope()
	return s.objWeek.GetTotal(params.UserId, currentWeekStart, currentWeekEnd)
}

func (s *AnalyseService) GetAvgCurrentWeek(params *request.AnalyseParamsAll) (string, error) {
	return s.objDetail.GetAvgCurrentWeek(params.UserId)
}

func (s *AnalyseService) GetAvgCurrentYear(params *request.AnalyseParamsAll) (string, error) {
	return s.objMonth.GetAvgYear(params.UserId, utils.GetCurrentYear())
}

func (s *AnalyseService) GetAvgCurrentWeekNoLoad(params *request.AnalyseParamsAll) (string, error) {
	return s.objDetail.GetAvgCurrentWeekNoLoad(params.UserId)
}

func (s *AnalyseService) GetAvgCurrentYearNoLoad(params *request.AnalyseParamsAll) (string, error) {
	return s.objMonth.GetAvgYearNoLoad(params.UserId, utils.GetCurrentYear())
}

func (s *AnalyseService) GetTotalCurrentYear(params *request.AnalyseParamsAll) (string, error) {
	return s.objMonth.GetYearTotal(params.UserId, utils.GetCurrentYear())
}
