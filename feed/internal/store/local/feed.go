package local

import (
	"gorm.io/gorm"
	"time"
)

type Video struct {
	gorm.Model
	PlayURL       string `gorm:"column:play_url; type:varchar(200)"`
	CoverURL      string `gorm:"column:cover_url; type:varchar(200)"`
	Title         string `gorm:"column:title; type:varchar(200)"`
	FavoriteCount int64  `gorm:"column:favorite_count; type:bigint"`
	CommentCount  int64  `gorm:"column:comment_count; type:bigint"`
	IsFavorite    int16  `gorm:"column:is_favorite; type:tinyint"`
	UserId        int64  `gorm:"column:user_id; type:bigint"`
}

type User struct {
	gorm.Model
	Name          string `gorm:"column:name; type:varchar(200)"`
	FollowerCount int64  `gorm:"column:follower_count; type:bigint"`
	FollowCount   int64  `gorm:"column:follow_count; type:bigint"`
}

func (m *Manager) QueryVideosAfter(n int, date time.Time) (videos []*Video) {
	db := m.handler.Table("t_video")
	db.Where("created_at < ?", date).Order("created_at desc").Limit(n).Find(&videos)
	return
}

func (m *Manager) QueryNameById(id int64) (name string) {
	db := m.handler.Table("t_user")
	var user User
	db.Where("id = ?", id).Find(&user)
	name = user.Name
	return
}

func (m *Manager) QueryUserById(id int64) (user User) {
	db := m.handler.Table("t_user")
	db.Where("id = ?", id).Find(&user)
	return
}
