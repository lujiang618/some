package models

type User struct {
	Id    uint64
	Name  string
	Phone uint64
	BaseModel
}

func NewUser() *User {
	user := &User{}

	user.SetTableName("users")

	return user
}

func (m *User) GetDetail(id uint64) (*User, error) {
	user := NewUser()
	err := m.Db().Where("id = ?", id).First(user).Error

	return user, err
}

// 根据token获取会员账号信息
func (m *User) GetDetailByToken(token string) (*User, error) {
	member := &User{}
	err := m.Db().Where("token = ?", token).First(member).Error

	return member, err
}
