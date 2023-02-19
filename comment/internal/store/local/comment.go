package local

import (
	"errors"
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
	AuthorId  int32          `json:"author_id"`
	VideoId   int32          `json:"video_id"`
	Msg       string         `json:"msg"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (c Comment) TableName() string {
	return "t_comment"
}

func (v Video) TableName() string {
	return "t_video"
}

func (m *Manager) InsertComment(comment *Comment) error {
	isExisted, err := m.SelectVideoById(comment.VideoId)
	if err != nil {
		return err
	}
	if !isExisted {
		return errors.New("video_id不存在")
	}
	return m.handler.Create(comment).Error
}

func (m *Manager) DeleteComment(comment Comment) error {
	return m.handler.Where("id = ?", comment.Id).Delete(&Comment{}).Error
}

// videoId合法性校验
func (m *Manager) SelectVideoById(id int32) (bool, error) {
	var video Video
	err := m.handler.Where("id=?", id).Select("id").Find(&video).Error

	if err != nil {
		return false, err
	}
	if 0 != video.ID {
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
func (m *Manager) UpdateCommentsCountByVideoId(id int32) error {
	var cnt int64
	tx := m.handler.Model(&Comment{}).Where("video_id=?", id).Count(&cnt)
	err := tx.Error
	if err != nil {
		return err
	} else if tx.RowsAffected == 0 {
		return errors.New("未查询到评论记录")
	}
	return m.handler.Model(&Video{}).Update("comment_count", cnt).Error
}
