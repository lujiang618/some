package filter

import (
	"some/api/application/service"
)

type Auth struct {
	BaseFilter
	service *service.Auth
}

func NewAuth() *Auth {
	return &Auth{
		service: service.NewAuth(),
	}
}

// // 验证登录的请求参数
// func (f *Auth) Login(c *gin.Context) (*models.User, *api.Error) {
// 	params := &request.LoginParams{}

// 	if err := c.ShouldBindJSON(params); err != nil {
// 		return nil, api.NewError(code.ErrorRequest, err.Error())
// 	}

// 	// 调用service对应的方法
// 	member, err := f.service.Login(params)

// 	return member, err
// }

// // 验证出的请求参数
// func (f *Auth) Logout(c *gin.Context) (*request.LogoutParams, error) {
// 	params := &request.LogoutParams{}

// 	if err := c.ShouldBind(params); err != nil {
// 		return nil, err
// 	}

// 	return params, nil
// }

// // 验证注册信息
// func (f *Auth) Register(c *gin.Context) (*models.User, *api.Error) {
// 	params := &request.RegisterParams{}

// 	if err := c.ShouldBind(params); err != nil {
// 		return nil, api.NewError(code.ErrorRequest, err.Error())
// 	}

// 	member, err := f.service.Register(params)

// 	return member, err
// }

// // 验证获取用户详情
// func (f *Auth) Forget(c *gin.Context) (*request.ForgetParams, error) {
// 	params := &request.ForgetParams{}

// 	if err := c.ShouldBindUri(params); err != nil {
// 		logrus.Error(err)

// 		return nil, err
// 	}

// 	return params, nil
// }
