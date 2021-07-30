package request

import "some/api/pkg/utils"

type CostParamList struct {
	UserId uint64 `json:"user_id" form:"user_id"`
}

type CostParamDetail struct {
	Id int `json:"id"`
}

type CostParamCreate struct {
	UserId      uint64         `json:"user_id" form:"user_id"`
	CategoryId  int            `json:"category_id" form:"category_id"`
	OccurDate   utils.JsonDate `json:"occur_date" form:"occur_date"`
	Content     string         `json:"content" form:"content"`
	Amount      string         `json:"amount" form:"amount"`
	Description string         `json:"description" form:"description"`
}

type CostParamUpdate struct {
	Id          int            `json:"id" form:"id"`
	UserId      uint64         `json:"user_id" form:"user_id"`
	CategoryId  int            `json:"category_id" form:"category_id"`
	OccurDate   utils.JsonDate `json:"occur_date" form:"occur_date"`
	Content     string         `json:"content" form:"content"`
	Amount      string         `json:"amount" form:"amount"`
	Description string         `json:"description" form:"description"`
}
