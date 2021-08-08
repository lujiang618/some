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
	Year        int            `json:"year"`
	Month       int            `json:"month"`
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

func (m *WealthCostDetail) GetList(params *request.CostParamList) (*[]response.CostObject, int, error) {
	details := make([]response.CostObject, 0)
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

	if params.Year > 0 {
		query = query.Where("year = ?", params.Year)
	}

	if params.Month > 0 {
		query = query.Where("month = ?", params.Month)
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

func (m *WealthCostDetail) GetCurrentMonth(userId uint64) (string, error) {
	dateStart, dateEnd := utils.GetCurrentMonthScope()

	return m.GetTotalByDateScope(userId, dateStart, dateEnd)
}

func (m *WealthCostDetail) GetCurrentWeek(userId uint64) (string, error) {
	dateStart, dateEnd := utils.GetCurrentWeekScope()

	return m.GetTotalByDateScope(userId, dateStart, dateEnd)
}

func (m *WealthCostDetail) GetTotalByDateScope(userId uint64, dateStart, dateEnd string) (string, error) {
	var total response.StatTotal
	err := m.Db().Select("round(sum(amount)) total").
		Where("user_id = ? and occur_date between ? and ?", userId, dateStart, dateEnd).
		First(&total).Error

	return total.Total, err
}

func (m *WealthCostDetail) GetAvgCurrentMonth(userId uint64) (string, error) {
	dateStart, dateEnd := utils.GetCurrentMonthScope()
	return m.GetAvgByDateScope(userId, dateStart, dateEnd)
}

func (m *WealthCostDetail) GetAvgCurrentWeek(userId uint64) (string, error) {
	dateStart, dateEnd := utils.GetCurrentWeekScope()
	return m.GetAvgByDateScope(userId, dateStart, dateEnd)
}

func (m *WealthCostDetail) GetAvgByDateScope(userId uint64, dateStart, dateEnd string) (string, error) {
	var avg response.StatAvg
	err := m.Db().Select("round(sum(amount)/?) as avg", utils.GetCurrentWeekDay()).
		Where("user_id = ? and occur_date between ? and ? ", userId, dateStart, dateEnd).
		First(&avg).Error
	return avg.Avg, err
}

func (m *WealthCostDetail) GetAvgCurrentWeekNoLoad(userId uint64) (string, error) {
	dateStart, dateEnd := utils.GetCurrentWeekScope()
	return m.GetAvgByDateScopeNoLoad(userId, dateStart, dateEnd)
}

func (m *WealthCostDetail) GetAvgByDateScopeNoLoad(userId uint64, dateStart, dateEnd string) (string, error) {
	var avg response.StatAvg
	err := m.Db().Select("round(sum(amount)/?) as avg", utils.GetCurrentWeekDay()).
		Where("user_id = ? and occur_date between ? and ? and category_id <> ?", userId, dateStart, dateEnd, 11).
		First(&avg).Error
	return avg.Avg, err
}

func (m *WealthCostDetail) GetDetail(id int) (*response.CostObject, error) {
	detail := &response.CostObject{}
	err := m.Db().Where("id = ?", id).First(detail).Error

	return detail, err
}

func (m *WealthCostDetail) Create(params *request.CostParamCreate) error {
	year, month, err := utils.GetYearMonth(params.OccurDate.String())
	if err != nil {
		return err
	}

	detail := &CostDetail{
		UserId:      params.UserId,
		CategoryId:  params.CategoryId,
		Year:        year,
		Month:       month,
		OccurDate:   params.OccurDate.String(),
		Content:     params.Content,
		Amount:      params.Amount,
		Description: params.Description,
	}

	return m.Db().Create(detail).Error
}

func (m *WealthCostDetail) Update(params *request.CostParamUpdate) error {
	detail := &CostDetail{
		UserId:      params.UserId,
		CategoryId:  params.CategoryId,
		OccurDate:   params.OccurDate.String(),
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
