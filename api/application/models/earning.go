package models

import (
	"some/api/application/objects/request"
	"some/api/application/objects/response"
	"some/api/pkg/utils"
)

type Earning struct {
	Id          int            `json:"id" gorm:"primary_key"`
	UserId      uint64         `json:"user_id"`
	Type        int            `json:"type"`
	Content     string         `json:"content"`
	Amount      string         `json:"amount"`
	Description string         `json:"description"`
	CreatedAt   utils.JSONTime `json:"created_at"  gorm:"-"`
	UpdatedAt   utils.JSONTime `json:"updated_at"  gorm:"-"`
}

type WealthEarning struct {
	DetailColumns []string
	BaseModel
}

func NewWWealthEarning() *WealthEarning {
	earning := &WealthEarning{}

	earning.SetTableName("wealth_earnings")

	return earning
}

func (m *WealthEarning) GetList(params *request.CostParamList) (*[]response.EarningObject, int, error) {
	details := make([]response.EarningObject, 0)
	query := m.Db().Where("user_id = ?", params.UserId)
	if len(params.DateRange) == 2 {
		query = query.Where("occur_date between ? and ?", params.DateRange[0], params.DateRange[1])
	}

	if params.CategoryId > 0 {
		query = query.Where("category_id = ?", params.CategoryId)
	}

	if params.Content != "" {
		query = query.Where("content like ?", "%"+params.Content+"%")
	}

	orderType := "DESC"
	if params.SortOrder == "ascend" {
		orderType = "ASC"
	}

	sortStr := ""
	switch params.SortField {
	case "amount":
		sortStr = "amount " + orderType
	case "type":
		sortStr = "type " + orderType
	default:
		sortStr = "created_at desc"
	}

	err := query.Order(sortStr).Find(&details).Error

	return &details, 100, err
}

func (m *WealthEarning) Create(params *request.ParamEarningCreate) error {

	detail := &Earning{
		UserId:      params.UserId,
		Type:        params.Type,
		Content:     params.Content,
		Amount:      params.Amount,
		Description: params.Description,
	}

	return m.Db().Create(detail).Error
}

func (m *WealthEarning) Update(params *request.ParamEarningUpdate) error {
	detail := &Earning{
		UserId:      params.UserId,
		Type:        params.Type,
		Content:     params.Content,
		Amount:      params.Amount,
		Description: params.Description,
	}

	tx := m.Db().Begin()
	err := tx.Where("id = ?", params.Id).Updates(detail).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (m *WealthEarning) GetCurrentTotal(userId uint64) (string, error) {
	var total response.StatTotal
	err := m.Db().Select("sum(amount) total").
		Where("user_id = ?", userId).First(&total).Error

	return total.Total, err
}
