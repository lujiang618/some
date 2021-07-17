package request

type ExpendParamList struct {
	UserId uint64
}

type ExpendParamDetail struct {
	Id int
}

type ExpendParamCreate struct {
	UserId      uint64 `json:"user_id"`
	Category    int    `json:"category"`
	OccurDate   string `json:"occur_date"`
	Content     string `json:"content"`
	Amount      string `json:"amount"`
	Description string `json:"description"`
}
