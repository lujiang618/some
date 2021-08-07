package request

type ParamEarningCreate struct {
	UserId      uint64 `json:"user_id" form:"user_id"`
	Type        int    `json:"type" form:"type"`
	Content     string `json:"content" form:"content"`
	Amount      string `json:"amount" form:"amount"`
	Description string `json:"description" form:"description"`
}

type ParamEarningUpdate struct {
	Id          int    `json:"id" form:"id"`
	UserId      uint64 `json:"user_id" form:"user_id"`
	Type        int    `json:"type" form:"type"`
	Content     string `json:"content" form:"content"`
	Amount      string `json:"amount" form:"amount"`
	Description string `json:"description" form:"description"`
}
