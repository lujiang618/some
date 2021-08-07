package request

import "some/api/pkg/utils"

/**
pageNo: 1
pageSize: 10
user_id: 1
occur_date: 2021-07-30
sortField: amount
sortOrder: ascend

**/
type CostParamList struct {
	UserId     uint64   `json:"user_id" form:"user_id"`
	Content    string   `json:"content" form:"content"`
	Year       int      `json:"year" form:"year"`
	Month      int      `json:"month" form:"month"`
	CategoryId int      `json:"category_id" form:"category_id"`
	DateRange  []string `json:"date_range" form:"date_range[]"`
	PageNo     int      `json:"pageNo" form:"pageNo"`
	PageSize   int      `json:"pageSize" form:"pageSize"`
	SortField  string   `json:"sortField" form:"sortField"`
	SortOrder  string   `json:"sortOrder" form:"sortOrder"`
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
