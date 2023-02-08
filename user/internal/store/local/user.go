package local

import (
	"github.com/JirafaYe/user/pkg"
	"strconv"
	"strings"
	"time"
)

type User struct {
	ID        int64      `gorm:"column:id"`
	Username  string     `gorm:"column:username"`
	Password  string     `gorm:"column:password"`
	CreateAt  time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
	Name      string     `gorm:"column:name"`
}

var user = "t_user"

func (m *Manager) GetUsernameExit(username string) bool {

	err := m.handler.Table(user).Where("username=?", username).First(&user, username).Error
	if err != nil {
		return false
	}
	return true
}

func (m *Manager) Register(username, password string) (int64, bool) {
	if err := pkg.Init("2006-01-02", 1); err != nil {
		panic(err)
		return 0, false
	}
	id := pkg.GenID()
	password = pkg.SaltEncodePwd(password)
	num := strconv.FormatInt(id, 10)
	name := "未命名" + num[15:19]
	u := User{
		ID:        id,
		Username:  username,
		Password:  password,
		CreateAt:  time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
		Name:      name,
	}
	if err := m.handler.Table(user).Create(u).Error; err != nil {
		panic(err)
		return 0, false
	}
	return id, true
}

func (m *Manager) Login(username, password string) (bool, int64, error) {
	var u User
	err := m.handler.Table(user).Where("username=?", username).First(&u).Error
	if err != nil {
		panic(err)
		return false, u.ID, err
	}
	pwd := strings.Split(u.Password, "$")
	userPassword := pkg.VerifyUserPassword(password, pwd[1], pwd[2])
	return userPassword, u.ID, nil
}

func (m *Manager) GetUserMsg(username string) User {
	var u User
	err := m.handler.Table(user).Where("username=?", username).First(&u).Error
	if err != nil {
		panic(err)
		return u
	}
	return u
}
