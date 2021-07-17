package request

type CostParamList struct {
	UserId uint64
}

type CostParamDetail struct {
	Id int
}

type CostParamCreate struct {
	UserId      uint64 `json:"user_id"`
	Category    int    `json:"category"`
	OccurDate   string `json:"occur_date"`
	Content     string `json:"content"`
	Amount      string `json:"amount"`
	Description string `json:"description"`
}
