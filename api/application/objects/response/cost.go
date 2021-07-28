package response

import "some/api/pkg/utils"

type CostDetailObject struct {
	Id          int            `json:"id" gorm:"primary_key"`
	UserId      uint64         `json:"user_id"`
	Category    int            `json:"category"`
	OccurDate   string         `json:"occur_date"`
	Content     string         `json:"content"`
	Amount      string         `json:"amount"`
	Description string         `json:"description"`
	CreatedAt   utils.JSONTime `json:"created_at"  gorm:"-"`
	UpdatedAt   utils.JSONTime `json:"updated_at"  gorm:"-"`
}

type CostListObject struct {
	PageSize   int                 `json:"page_size"`
	PageNo     int                 `json:"page_no"`
	TotalCount int                 `json:"total_count"`
	TotalPage  int                 `json:"total_page"`
	Data       *[]CostDetailObject `json:"data"`
}

type CostCategoryObject struct {
	Id        int            `json:"id" gorm:"primary_key"`
	Name      string         `json:"name"`
	UserId    int            `json:"category"`
	ParrentId int            `json:"parrent_id"`
	CreatedAt utils.JSONTime `json:"created_at"  gorm:"-"`
	UpdatedAt utils.JSONTime `json:"updated_at"  gorm:"-"`
}

type CostCategoryListObject struct {
	PageSize   int                   `json:"page_size"`
	PageNo     int                   `json:"page_no"`
	TotalCount int                   `json:"total_count"`
	TotalPage  int                   `json:"total_page"`
	Data       *[]CostCategoryObject `json:"data"`
}
