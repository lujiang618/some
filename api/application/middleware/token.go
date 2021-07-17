package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"some/api/application/models"
	"some/api/pkg/constant"
	"some/api/pkg/utils"

	"some/api/pkg/db"

	"some/api/pkg/api"
	"some/api/pkg/api/code"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 做权限验证，除了auth控制器之外，其他控制器都需要有有效的token
func VerifyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		var userId uint64
		var account string
		var isOk bool

		// 先验证access token
		if userId, account = VerifyAccessToken(c); userId > 0 && len(account) > 0 {
			isOk = true
		}

		// 如果access token 没有验证通过，则验证api token
		if !isOk {
			if userId, account = VerifyToken(c); userId > 0 && len(account) > 0 {
				isOk = true
			}
		}

		if !isOk {
			api.SetResponse(c, http.StatusForbidden, code.ErrorSign, "")
			return
		}

		//将会员信息放到gin框架的context中
		c.Set("user_id", int64(userId)) // gin框架 没有c.Getint()方法，这里转成int64
		c.Set("account", account)

		c.Next()
	}
}

// 验证API TOKEN
func VerifyToken(c *gin.Context) (uint64, string) {
	token := c.GetHeader("Token")

	if len(token) < 1 {
		return 0, ""
	}

	// 先从redis内读取缓存
	User, err := GetUser(token)
	if err != nil {
		return 0, ""
	}

	return User.Id, User.Name
}

// 验证ACCESS TOKEN
func VerifyAccessToken(c *gin.Context) (uint64, string) {
	accessToken := c.GetHeader("Access-Token")
	if len(accessToken) < 1 {
		return 0, ""
	}

	objJwt := utils.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := objJwt.ParseToken(accessToken)
	if err != nil {
		logrus.Error("access token err:", err)
		return 0, ""
	}

	return claims.ID, claims.Account
}

// 根据token获取user信息
func GetUser(token string) (*models.User, error) {
	// 先从redis内读取缓存
	User, _ := GetUserFromRedis(token)
	if User != nil {
		return User, nil
	}

	User, err := models.NewUser().GetDetailByToken(token)

	userData, errJson := json.Marshal(User)
	if errJson != nil {
		logrus.Error(err)
	}

	errJson = db.Redis.Set(fmt.Sprintf("%s%s", constant.UserCacheKey, token), userData, time.Minute*60).Err()
	if errJson != nil {
		logrus.Error(err)
	}

	return User, err
}

// 从缓存中获取会员信息
func GetUserFromRedis(token string) (*models.User, *api.Error) {
	userData, err := db.Redis.Get(fmt.Sprintf("%s%s", constant.UserCacheKey, token)).Bytes()
	if err != nil {
		return nil, api.NewError(code.ErrorSyntax, code.GetMessage(code.ErrorSyntax))
	}
	if len(userData) == 0 {
		return nil, nil
	}

	User := &models.User{}

	if err := json.Unmarshal(userData, User); err != nil {
		return nil, api.NewError(code.ErrorSyntax, code.GetMessage(code.ErrorSyntax))
	}

	return User, nil
}
