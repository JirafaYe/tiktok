package local

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	Id        int32          `json:"id"`
	AuthorId  int32          `json:"author_id"`
	VideoId   int32          `json:"video_id"`
	Msg       string         `json:"msg"`
	Likes     int32          `json:"likes"`
	Unlikes   int32          `json:"unlikes"`
	IsTopped  bool           `json:"is_topped"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (c Comment) TableName() string {
	return "t_comment"
}

func (m *Manager) InsertComment(comment *Comment) error {
	return m.handler.Create(comment).Error
}

func (m *Manager) DeleteComment(comment Comment) error {
	return m.handler.Where("id = ?", comment.Id).Delete(&Comment{}).Error
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
