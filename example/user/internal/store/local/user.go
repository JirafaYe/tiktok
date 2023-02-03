package local

type User struct {
	Id       int32  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var user = "t_user"

func (m *Manager) InsertUser(u User) error {
	return m.handler.Table(user).Create(&u).Error
}

func (m *Manager) SelectUserByUsername(username string) (User, error) {
	var u User
	err := m.handler.Table(user).Where("username = ?", username).Take(&u).Error
	return u, err
}
