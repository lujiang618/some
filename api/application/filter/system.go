package filter

import "some/api/application/service"

type System struct {
	BaseFilter
	service *service.System
}

func NewSystem() *System {
	return &System{
		service: service.NewSystem(),
	}
}
