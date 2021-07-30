package models

import (
	"some/api/application/objects/request"
	"some/api/application/objects/response"
	"some/api/pkg/utils"
)

type CostDetail struct {
	Id          int            `json:"id" gorm:"primary_key"`
	UserId      uint64         `json:"user_id"`
	CategoryId  int            `json:"category_id"`
	OccurDate   string         `json:"occur_date"`
	Content     string         `json:"content"`
	Amount      string         `json:"amount"`
	Description string         `json:"description"`
	CreatedAt   utils.JSONTime `json:"created_at"  gorm:"-"`
	UpdatedAt   utils.JSONTime `json:"updated_at"  gorm:"-"`
}

type WealthCostDetail struct {
	DetailColumns []string
	BaseModel
}

func NewWealthCostDetail() *WealthCostDetail {
	detail := &WealthCostDetail{}

	detail.SetTableName("wealth_cost_details")

	return detail
}

func (m *WealthCostDetail) GetList(params *request.CostParamList) (*[]response.CostDetailObject, int, error) {
	details := make([]response.CostDetailObject, 0)
	query := m.Db().Where("user_id = ?", params.UserId)
	if params.OccurDate != "" {
		query = query.Where("occur_date = ?", params.OccurDate)
	}

	orderType := "DESC"
	if params.SortOrder == "ascend" {
		orderType = "ASC"
	}

	sortStr := ""
	switch params.SortField {
	case "amount":
		sortStr = "amount " + orderType
	case "category_id":
		sortStr = "category_id " + orderType
	case "occur_date":
		sortStr = "occur_date " + orderType
	default:
		sortStr = "occur_date desc"
	}

	err := query.Order(sortStr).Find(&details).Error

	return &details, 100, err
}

func (m *WealthCostDetail) GetDetail(id int) (*response.CostDetailObject, error) {
	detail := &response.CostDetailObject{}
	err := m.Db().Where("id = ?", id).First(detail).Error

	return detail, err
}

func (m *WealthCostDetail) Create(params *request.CostParamCreate) error {
	detail := &CostDetail{
		UserId:      params.UserId,
		CategoryId:  params.CategoryId,
		OccurDate:   params.OccurDate.String(),
		Content:     params.Content,
		Amount:      params.Amount,
		Description: params.Description,
	}

	return m.Db().Create(detail).Error
}

func (m *WealthCostDetail) Update(params *request.CostParamUpdate) error {
	detail := &response.CostDetailObject{
		UserId:      params.UserId,
		CategoryId:  params.CategoryId,
		OccurDate:   params.OccurDate,
		Content:     params.Content,
		Amount:      params.Amount,
		Description: params.Description,
	}

	return m.Db().Updates(detail).Where("id = ?", params.Id).Error
}
