package service

type User struct {
}

func NewUser() *User {
	return &User{}
}

// // 获取会员详细信息
// func (s *User) GetUser(userId uint64) (*models.User, error) {
// 	User, err := models.NewUser().GetDetail(userId)

// 	return User, err
// }
