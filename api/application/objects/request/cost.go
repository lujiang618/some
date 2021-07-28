package request

type CostParamList struct {
	UserId uint64 `json:"user_id" form:"user_id"`
}

type CostParamDetail struct {
	Id int `json:"id"`
}

type CostParamCreate struct {
	UserId      uint64 `json:"user_id" form:"user_id"`
	Category    int    `json:"category" form:"category"`
	OccurDate   string `json:"occur_date" form:"occur_date"`
	Content     string `json:"content" form:"content"`
	Amount      string `json:"amount" form:"amount"`
	Description string `json:"description" form:"description"`
}

type CostParamUpdate struct {
	Id          int    `json:"id"`
	UserId      uint64 `json:"user_id"`
	Category    int    `json:"category"`
	OccurDate   string `json:"occur_date"`
	Content     string `json:"content"`
	Amount      string `json:"amount"`
	Description string `json:"description"`
}
