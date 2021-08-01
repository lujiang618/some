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
	CurrentMonth string `json:"current_month"`
	LastMonth    string `json:"last_month"`
	CurrentWeek  string `json:"current_week"`
	LastWeek     string `json:"last_week"`
}

type StatTotal struct {
	Total string
}
