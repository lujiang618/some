package response

import "some/api/pkg/utils"

/**
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL DEFAULT '0',
  `year` int(10) unsigned NOT NULL COMMENT '年份',
  `year_month` int(10) unsigned NOT NULL COMMENT '年月',
  `category_id` int(11) NOT NULL DEFAULT '0' COMMENT '类别',
  `total` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '合计',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  **/
type CostMonth struct {
	Id          int            `json:"id" gorm:"primary_key"`
	UserId      uint64         `json:"user_id"`
	YearMonth   string         `json:"year_month"`
	CategoryId  int            `json:"category_id"`
	Year        int            `json:"year"`
	Total       string         `json:"total"`
	Amount      string         `json:"amount"`
	Description string         `json:"description"`
	CreatedAt   utils.JSONTime `json:"created_at"  gorm:"-"`
	UpdatedAt   utils.JSONTime `json:"updated_at"  gorm:"-"`
}

type Stat struct {
	CurrentMonth         string `json:"current_month"`
	LastMonth            string `json:"last_month"`
	CurrentWeek          string `json:"current_week"`
	LastWeek             string `json:"last_week"`
	TotalYear            string `json:"total_year"`
	AvgCurrentWeek       string `json:"avg_current_week"`
	AvgCurrentWeekNoLoad string `json:"avg_current_week_no_load"`
	AvgCurrentYear       string `json:"avg_current_year"`
	AvgCurrentYearNoLoad string `json:"avg_current_year_no_load"`
}

type Charts struct {
	TotalEarning         string              `json:"totalEarning"`
	TotalCost            string              `json:"totalCost"`
	TotalCurrentMonth    string              `json:"totalCurrentMonth"`
	TotalCurrentDay      string              `json:"totalCurrentDay"`
	AvgYear              string              `json:"avgYear"`
	AvgMonth             string              `json:"avgMonth"`
	AvgWeek              string              `json:"avgWeek"`
	PieScaleWeek         []PieScale          `json:"pieScaleWeek"`
	PieScaleWeekNoLoad   []PieScale          `json:"pieScaleWeekNoLoad"`
	PieScaleMonth        []PieScale          `json:"pieScaleMonth"`
	PieScaleMonthNoLoad  []PieScale          `json:"pieScaleMonthNoLoad"`
	TotalDays            []TotalDay          `json:"totalDays"`
	CostTopMonth         []CostTop           `json:"costTopMonth"`
	CostTopWeek          []CostTop           `json:"costTopWeek"`
	TotalMonths          []TotalMonth        `json:"totalMonths"`
	TotalWeeks           []TotalWeek         `json:"totalWeeks"`
	TotalYearCategories  []TotalYearCategory `json:"totalYearCategories"`
	TotalMonthCategories []TotalYearCategory `json:"totalMonthCategories"`
}

type PieScale struct {
	Item  string  `json:"item"`
	Count float64 `json:"count"`
}

type TotalCategory struct {
	CategoryId int     `json:"x"`
	Total      float64 `json:"y"`
}

type TotalYearCategory struct {
	Category string  `json:"x"`
	Total    float64 `json:"y"`
}

type TotalMonth struct {
	Month string  `json:"x"`
	Total float64 `json:"y"`
}

type TotalWeek struct {
	Week  string  `json:"x"`
	Total float64 `json:"y"`
}

type TotalDay struct {
	Date  string  `json:"x"`
	Total float64 `json:"y"`
}

type CostTop struct {
	Name  string  `json:"name"`
	Total float64 `json:"total"`
}

type StatTotal struct {
	Total string
}

type StatAvg struct {
	Avg string
}
