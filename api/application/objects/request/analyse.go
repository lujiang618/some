package request

type AnalyseParamsAll struct {
	UserId uint64 `json:"user_id" form:"user_id"`
}

type AnalyseParamsCharts struct {
	UserId uint64 `json:"user_id" form:"user_id"`
}
