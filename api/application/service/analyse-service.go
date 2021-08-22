package service

import (
	"some/api/application/models"
	"some/api/application/objects/request"
	"some/api/application/objects/response"
	"some/api/pkg/api"
	"some/api/pkg/api/code"
	"some/api/pkg/utils"
	"strconv"
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
	objDetail   *models.WealthCostDetail
	objMonth    *models.WealthStatisticsMonth
	objWeek     *models.WealthStatisticsWeak
	objEarning  *models.WealthEarning
	objCategory *models.WealthCostCategory
}

func NewAnalyseService() *AnalyseService {
	return &AnalyseService{
		objDetail:   models.NewWealthCostDetail(),
		objMonth:    models.NewWealthStatisticsMonth(),
		objWeek:     models.NewWealthStatisticsWeak(),
		objEarning:  models.NewWWealthEarning(),
		objCategory: models.NewWealthCostCategory(),
	}
}
func (s *AnalyseService) GetCharts(params *request.AnalyseParamsCharts) (*response.Charts, *api.Error) {
	totalEarning, err := s.objEarning.GetCurrentTotal(params.UserId)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	totalCost, err := s.GetTotalCost(params.UserId)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	pieScaleWeek, err := s.GetPieScaleWeek(params.UserId)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	pieScaleWeekNoLoad, err := s.GetPieScaleWeekNoLoad(params.UserId)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	pieScaleMonth, err := s.GetPieScaleMonth(params.UserId)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	pieScaleMonthNoLoad, err := s.GetPieScaleMonthNoLoad(params.UserId)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	totalDays, err := s.GetTotalDay(params.UserId)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	costTopWeek, err := s.GetCostTopWeek(params.UserId)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	costTopMonth, err := s.GetCostTopMonth(params.UserId)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	totalMonths, err := s.GetTotalMonths(params.UserId)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	totalWeeks, err := s.GetTotalWeeks(params.UserId)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	totalYearCategories, err := s.GetTotalYearCategory(params.UserId)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	totalMonthCategories, err := s.GetTotalMonthCategory(params.UserId)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	// 获取当前月支出
	totalCurrentMonth, err := s.objDetail.GetCurrentMonth(params.UserId)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	totalCurrentDay, err := s.objDetail.GetCurrentDay(params.UserId)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	avgYear, err := s.objMonth.GetAvgYear(params.UserId, utils.GetCurrentYear())
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	avgMonth, err := s.objDetail.GetAvgCurrentMonth(params.UserId)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	avgWeek, err := s.objDetail.GetAvgCurrentWeek(params.UserId)
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	data := &response.Charts{
		TotalEarning:         totalEarning,
		TotalCost:            totalCost,
		TotalCurrentMonth:    totalCurrentMonth,
		TotalCurrentDay:      totalCurrentDay,
		AvgYear:              avgYear,
		AvgMonth:             avgMonth,
		AvgWeek:              avgWeek,
		PieScaleWeek:         pieScaleWeek,
		PieScaleMonth:        pieScaleMonth,
		PieScaleMonthNoLoad:  pieScaleMonthNoLoad,
		PieScaleWeekNoLoad:   pieScaleWeekNoLoad,
		TotalDays:            totalDays,
		CostTopMonth:         costTopMonth,
		CostTopWeek:          costTopWeek,
		TotalMonths:          totalMonths,
		TotalWeeks:           totalWeeks,
		TotalYearCategories:  totalYearCategories,
		TotalMonthCategories: totalMonthCategories,
	}

	return data, nil
}

func (s *AnalyseService) GetTotalCost(userId uint64) (string, error) {
	return s.objMonth.GetTotalYear(userId)
}

func (s *AnalyseService) GetTotalWeeks(userId uint64) ([]response.TotalWeek, error) {
	return s.objWeek.GetTotalWeeks(userId)
}

func (s *AnalyseService) GetTotalMonths(userId uint64) ([]response.TotalMonth, error) {
	return s.objMonth.GetTotalMonths(userId)
}

func (s *AnalyseService) GetCostTopWeek(userId uint64) ([]response.CostTop, error) {
	dateStart, dateEnd := utils.GetCurrentWeekScope()
	tops, err := s.objDetail.GetTop(userId, dateStart, dateEnd)
	if err != nil {
		return nil, err
	}

	return s.FormatCostTop(tops)
}

func (s *AnalyseService) GetCostTopMonth(userId uint64) ([]response.CostTop, error) {
	dateStart, dateEnd := utils.GetCurrentMonthScope()

	tops, err := s.objDetail.GetTop(userId, dateStart, dateEnd)
	if err != nil {
		return nil, err
	}

	return s.FormatCostTop(tops)
}

func (s *AnalyseService) FormatCostTop(tops []response.CostObject) ([]response.CostTop, error) {
	categoryIds := make([]int, 0, len(tops))
	for _, top := range tops {
		categoryIds = append(categoryIds, top.CategoryId)
	}

	categoryId2Name, err := s.objCategory.GetCategoryMap(categoryIds)
	if err != nil {
		return nil, err
	}

	result := make([]response.CostTop, 0, len(tops))
	for _, top := range tops {
		total, _ := strconv.ParseFloat(top.Amount, 64)
		result = append(result, response.CostTop{
			Total: total,
			Name:  categoryId2Name[top.CategoryId] + "-" + top.Content,
		})
	}

	return result, nil
}

func (s *AnalyseService) GetTotalDay(userId uint64) ([]response.TotalDay, error) {
	dateStart := utils.GetSubDate(15)
	dateEnd := utils.GetCurrentDate()

	return s.objDetail.GetTotalDay(userId, dateStart, dateEnd)
}

func (s *AnalyseService) GetPieScaleMonth(userId uint64) ([]response.PieScale, error) {
	dateStart, dateEnd := utils.GetCurrentMonthScope()

	categoryPercents, err := s.objDetail.GetPercent(userId, dateStart, dateEnd)
	if err != nil {
		return nil, err
	}

	return s.FormatPieScal(categoryPercents)
}

func (s *AnalyseService) GetPieScaleMonthNoLoad(userId uint64) ([]response.PieScale, error) {
	dateStart, dateEnd := utils.GetCurrentMonthScope()

	categoryPercents, err := s.objDetail.GetPercentNoLoad(userId, dateStart, dateEnd)
	if err != nil {
		return nil, err
	}

	return s.FormatPieScal(categoryPercents)
}

// 类别的金额 / 总金额
func (s *AnalyseService) GetPieScaleWeek(userId uint64) ([]response.PieScale, error) {
	dateStart, dateEnd := utils.GetCurrentWeekScope()

	categoryPercents, err := s.objDetail.GetPercent(userId, dateStart, dateEnd)
	if err != nil {
		return nil, err
	}

	return s.FormatPieScal(categoryPercents)
}

func (s *AnalyseService) GetPieScaleWeekNoLoad(userId uint64) ([]response.PieScale, error) {
	dateStart, dateEnd := utils.GetCurrentWeekScope()

	categoryPercents, err := s.objDetail.GetPercentNoLoad(userId, dateStart, dateEnd)
	if err != nil {
		return nil, err
	}

	return s.FormatPieScal(categoryPercents)
}

func (s *AnalyseService) GetTotalYearCategory(userId uint64) ([]response.TotalYearCategory, error) {
	categories, err := s.objMonth.GetTotalCategory(userId)
	if err != nil {
		return nil, err
	}

	return s.FormatTotalCategory(categories)
}

func (s *AnalyseService) GetTotalMonthCategory(userId uint64) ([]response.TotalYearCategory, error) {
	dateStart, dateEnd := utils.GetCurrentMonthScope()
	categories, err := s.objDetail.GetTotalCategory(userId, dateStart, dateEnd)
	if err != nil {
		return nil, err
	}

	return s.FormatTotalCategory(categories)
}

func (s *AnalyseService) FormatTotalCategory(categoryPercents []response.TotalCategory) ([]response.TotalYearCategory, error) {
	categoryIds := make([]int, 0, len(categoryPercents))
	for _, categoryPercent := range categoryPercents {
		categoryIds = append(categoryIds, categoryPercent.CategoryId)
	}

	categoryId2Name, err := s.objCategory.GetCategoryMap(categoryIds)
	if err != nil {
		return nil, err
	}

	result := make([]response.TotalYearCategory, 0, len(categoryPercents))
	for _, categoryPercent := range categoryPercents {
		result = append(result, response.TotalYearCategory{
			Category: categoryId2Name[categoryPercent.CategoryId],
			Total:    categoryPercent.Total,
		})
	}

	return result, nil
}

func (s *AnalyseService) FormatPieScal(categoryPercents []response.TotalCategory) ([]response.PieScale, error) {
	categoryIds := make([]int, 0, len(categoryPercents))
	for _, categoryPercent := range categoryPercents {
		categoryIds = append(categoryIds, categoryPercent.CategoryId)
	}

	categoryId2Name, err := s.objCategory.GetCategoryMap(categoryIds)
	if err != nil {
		return nil, err
	}

	result := make([]response.PieScale, 0, len(categoryPercents))
	for _, categoryPercent := range categoryPercents {
		result = append(result, response.PieScale{
			Item:  categoryId2Name[categoryPercent.CategoryId],
			Count: categoryPercent.Total,
		})
	}

	return result, nil
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
