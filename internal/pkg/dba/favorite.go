package dba

import (
	"github.com/JirafaYe/tiktok/internal/app/service"
	"gorm.io/gorm"
)

type UserFavorite struct {
	gorm.Model
	UserId  int64 `gorm:"primaryKey;autoIncrement:false"`
	VideoId int64 `gorm:"primaryKey;autoIncrement:false"`
}

//type Video struct {
//	gorm.Model
//	PlayURL       string `gorm:"column:play_url; type:varchar(200)"`
//	CoverURL      string `gorm:"column:cover_url; type:varchar(200)"`
//	Title         string `gorm:"column:title; type:varchar(200)"`
//	FavoriteCount int64  `gorm:"column:favorite_count; type:bigint"`
//	CommentCount  int64  `gorm:"column:comment_count; type:bigint"`
//	IsFavorite    bool   `gorm:"column:is_favorite; type:tinyint"`
//	UserId        int64  `gorm:"column:user_id; type:bigint"`
//}
//
//type User struct {
//	gorm.Model
//	Name          string `gorm:"column:name; type:varchar(200)"`
//	FollowerCount int64  `gorm:"column:follower_count; type:bigint"`
//	FollowCount   int64  `gorm:"column:follow_count; type:bigint"`
//}

type UserFavoriteIndex struct {
}

var favoriteTableName = "t_favorite"
var videoTableName = "t_video"
var userTableName = "t_user"

func (m *Manager) SelectUserFavorite(userId int64, videoId int64) (UserFavorite, error) {
	var u UserFavorite
	err := m.handler.Table(favoriteTableName).Where("user_id = ? AND video_id = ?", userId, videoId).Take(&u).Error
	return u, err
}

func (m *Manager) InsertUserFavorite(userId int64, videoId int64) error {
	u := UserFavorite{
		UserId:  userId,
		VideoId: videoId,
	}
	return m.handler.Table(favoriteTableName).Create(&u).Error
}

func (m *Manager) UpdateVideoLike(videoId int64, num int) error {
	v := &Video{}
	v.ID = uint(videoId)
	return m.handler.Table(videoTableName).Model(v).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", num)).Error
}

func (m *Manager) DeleteUserFavorite(userId int64, videoId int64) error {
	// 物理删除
	return m.handler.Table(favoriteTableName).Unscoped().Delete(&UserFavorite{}, "user_id = ? AND video_id = ?", userId, videoId).Error
}

func (m *Manager) SelectLikesByUser(userId int64) ([]int64, error) {
	var userFavorites []int64
	err := m.handler.Table(favoriteTableName).Select("video_id").Where("user_id = ?", userId).Find(&userFavorites).Error
	return userFavorites, err
}

func (m *Manager) SelectVideos(videoIds []int64) ([]Video, error) {
	var videos []Video
	err := m.handler.Table(videoTableName).Where(videoIds).Find(&videos).Error
	return videos, err
}

func (m *Manager) SelectUsers(userIds []int64) ([]service.UserFeed, error) {
	var users []service.UserFeed
	err := m.handler.Table(userTableName).Where(userIds).Find(&users).Error
	return users, err
}
