package crontab

import (
	"some/api/application/models"
)

type Cost struct {
	objDetail *models.WealthCostDetail
}

func NewCost() *Cost {
	return &Cost{
		objDetail: models.NewWealthCostDetail(),
	}
}

func (c *Cost) Run() {

	// 房贷
	params := readData("./load.json")
	c.objDetail.Create(params)

	// 房租
	params = readData("./rental.json")
	c.objDetail.Create(params)
}
