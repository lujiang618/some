package service

type Auth struct {
}

func NewAuth() *Auth {
	return &Auth{}
}

// 登录
// func (s *Auth) Login(params *request.LoginParams) (*models.User, *api.Error) {
// 	objuser := models.NewUser()

// 	user, err := objuser.GetDetailByAccount(params)
// 	if err != nil {
// 		return nil, err
// 	}

// 	ojbJwt := utils.NewJWT()
// 	claims := utils.CustomClaims{
// 		ID:      user.Id,
// 		Name:    user.Name,
// 		Mobile:  user.Mobile,
// 		Account: user.Account,
// 		StandardClaims: jwt.StandardClaims{
// 			NotBefore: int64(time.Now().Unix() - 1000),    // 签名生效时间
// 			ExpiresAt: int64(time.Now().Unix() + 3600*24), // 过期时间 一小时
// 			Issuer:    "gelu",                             //签名的发行者
// 		},
// 	}

// 	var errJwt error
// 	user.AccessToken, errJwt = ojbJwt.CreateToken(claims)
// 	if errJwt != nil {
// 		return nil, api.NewError(code.ErrorSyntax, code.GetMessage(code.ErrorSyntax))
// 	}

// 	return user, err
// }

// 注册
// func (s *Auth) Register(params *request.RegisterParams) (*models.User, *api.Error) {
// 	userObj := models.NewUser()

// 	user := &models.User{
// 		Account:  params.Account,
// 		Email:    params.Email,
// 		Name:     params.Name,
// 		Password: utils.GetMd5(params.Password),
// 		Status:   constant.StatusEnable,
// 		Token:    utils.GetRandom(20),
// 	}

// 	user, err := userObj.Add(user)

// 	return user, err
// }
