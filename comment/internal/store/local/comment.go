package local

import (
	"errors"
	"github.com/JirafaYe/comment/internal/service"
	"gorm.io/gorm"
	"time"
)

// video结构体
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

type Comment struct {
	Id        int32          `json:"id"`
	AuthorId  int64          `json:"author_id"`
	VideoId   int32          `json:"video_id"`
	Msg       string         `json:"msg"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// User结构体
type User struct {
	gorm.Model
	Name          string `gorm:"column:username; type:varchar(200)"`
	FollowerCount int64  `gorm:"column:follower_count; type:bigint"`
	FollowCount   int64  `gorm:"column:follow_count; type:bigint"`
}

func (c Comment) TableName() string {
	return "t_comment"
}

func (v Video) TableName() string {
	return "t_video"
}

func (u User) TableName() string {
	return "t_user"
}

func (m *Manager) GetUserMsg(id []int64) ([]service.CommentUser, error) {
	var users []service.CommentUser
	tx := m.handler.Model(&User{}).Select("id,username name").Where(id).Find(&users)
	return users, tx.Error
}

func (m *Manager) InsertComment(comment *Comment) error {
	var err error
	transaction := m.handler.Begin()

	rows := transaction.Create(comment)
	if rows.RowsAffected != 1 {
		err = errors.New("插入失败")
		transaction.Rollback()
		return err
	}
	err = transaction.Model(&Video{}).Where("id = ?", comment.VideoId).
		UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error
	if err != nil {
		transaction.Rollback()
		return err
	}
	transaction.Commit()
	return nil
}

func (m *Manager) DeleteComment(comment Comment) error {
	var err error

	rows := m.handler.Where("id = ?", comment.Id).Delete(&Comment{})
	if rows.RowsAffected != 1 {
		err = errors.New("无效评论")
		return err
	}

	return m.UpdateCommentsCountByVideoId(comment.VideoId, -1)

}

// videoId合法性校验
func (m *Manager) SelectVideoById(id int32) (bool, error) {
	var video Video
	tx := m.handler.Where("id=?", id).Select("id").Find(&video)
	err := tx.Error

	if err != nil {
		return false, err
	}
	if 0 != tx.RowsAffected {
		return true, nil
	} else {
		return false, nil
	}
}

func (m *Manager) SelectCommentNumsByVideoId(id int32) (int64, error) {
	var cnt int64
	err := m.handler.Model(&Comment{}).Where("video_id=?", id).Count(&cnt).Error
	return cnt, err
}

func (m *Manager) SelectCommentListByVideoId(id int32) ([]Comment, error) {
	var list []Comment
	err := m.handler.Where("video_id = ?", id).Order("created_at").Find(&list).Error
	return list, err
}

// 更新评论数
func (m *Manager) UpdateCommentsCountByVideoId(id int32, num int32) error {
	return m.handler.Model(&Video{}).Where("id = ?", id).
		UpdateColumn("comment_count", gorm.Expr("comment_count + ?", num)).Error
}
