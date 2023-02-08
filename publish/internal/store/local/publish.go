package local

import (
	"context"

	"gorm.io/gorm"
	// "github.com/JirafaYe/feed/internal/store/local"// Video结构体
)

// DB == m.handler

type Video struct {
	gorm.Model
	UserId        int64  `gorm:"column:user_id; type:bigint"`
	PlayURL       string `gorm:"column:play_url; type:varchar(200)"`
	CoverURL      string `gorm:"column:cover_url; type:varchar(200)"`
	Title         string `gorm:"column:title; type:varchar(200)"`
}

type User struct {
	gorm.Model
	UserId		  int64  `gorm:"column:user_id; type:bigint"`
}

// CreateVideo 创建视频
func (m *Manager) CreateVideo(ctx context.Context, video *Video) error {
	err := m.handler.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(video).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// PublishList 返回一系列视频，带有authorID
func (m *Manager) PublishList(ctx context.Context, authorId int64) ([]*Video, error) {
	var pubList []*Video
	err := m.handler.WithContext(ctx).Model(&Video{}).Where(&Video{AuthorID: int(authorId)}).Find(&pubList).Error
	if err != nil {
		return nil, err
	}
	return pubList, nil
}
