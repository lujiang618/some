package response

import "some/api/pkg/utils"

type EarningObject struct {
	Id          int            `json:"id" gorm:"primary_key"`
	UserId      uint64         `json:"user_id"`
	Type        int            `json:"type"`
	TypeName    string         `json:"type_name"`
	Content     string         `json:"content"`
	Amount      string         `json:"amount"`
	Description string         `json:"description"`
	CreatedAt   utils.JSONTime `json:"created_at"  gorm:"-"`
	UpdatedAt   utils.JSONTime `json:"updated_at"  gorm:"-"`
}

type EarningListObject struct {
	PageSize   int              `json:"page_size"`
	PageNo     int              `json:"page_no"`
	TotalCount int              `json:"total_count"`
	TotalPage  int              `json:"total_page"`
	Data       *[]EarningObject `json:"data"`
}
