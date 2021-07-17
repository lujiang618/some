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
