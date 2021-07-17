package filter

import (
	"github.com/gin-gonic/gin"
)

type Filter func(c *gin.Context)

// type Member struct {
// 	BaseFilter
// 	service *service.Member
// }

// func NewMember() *Member {
// 	return &Member{
// 		service: service.NewMember(),
// 	}
// }

// // 验证获取用户详情
// func (f *Member) View(c *gin.Context) (*models.User, *api.Error) {
// 	params := &request.ViewParams{}

// 	if err := c.ShouldBindUri(params); err != nil {

// 		return nil, api.NewError(code.ErrorRequest, err.Error())
// 	}

// 	// 调用service对应的方法
// 	member, err := f.service.GetUser(params.Id)

// 	return member, err
// }
