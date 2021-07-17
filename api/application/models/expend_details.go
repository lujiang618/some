package models

import (
	"some/api/application/objects/request"
	"some/api/pkg/utils"
)

type ExpendDetail struct {
	DetailColumns []string
	BaseModel
}

type ExpendDetailObject struct {
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

func NewExpendDetail() *ExpendDetail {
	detail := &ExpendDetail{}

	detail.SetTableName("wealth_expend_statistics_months")

	return detail
}

func (m *ExpendDetail) GetList(userId int) (*[]ExpendDetailObject, error) {
	details := make([]ExpendDetailObject, 0)
	err := m.Db().Where("user_id = ?", userId).Find(&details).Error

	return &details, err
}

func (m *ExpendDetail) GetDetail(id int) (*ExpendDetailObject, error) {
	detail := &ExpendDetailObject{}
	err := m.Db().Where("id = ?", id).First(detail).Error

	return detail, err
}

func (m *ExpendDetail) Create(params *request.ExpendParamCreate) error {
	detail := &ExpendDetailObject{
		UserId:      params.UserId,
		Category:    params.Category,
		OccurDate:   params.OccurDate,
		Content:     params.Content,
		Amount:      params.Amount,
		Description: params.Description,
	}

	return m.Db().Create(detail).Error
}
