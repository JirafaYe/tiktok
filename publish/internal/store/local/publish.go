package local

import (
	//"context"
	"gorm.io/gorm"
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

func (v Video) TableName() string {
	return "t_video"
}

func (u User) TableName() string {
	return "t_user"
}

// 创建视频
func (m *Manager) CreateVideo(video *Video) error {
	return m.handler.Create(video).Error
}

// TODO: 采用事务的方式创建视频
// func (m *Manager) CreateVideo(ctx context.Context, video *Video) error {
// 	err := m.handler.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
// 		err := tx.Create(video).Error
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// 	return err
// }

/* t_video 表
+----------------+-------------+------+-----+-------------------+-----------------------------+
| Field          | Type        | Null | Key | Default           | Extra                       |
+----------------+-------------+------+-----+-------------------+-----------------------------+
| video_id       | bigint(20)  | NO   | PRI | NULL              | auto_increment              |
| user_id        | bigint(20)  | YES  |     | NULL              |                             |
| play_url       | varchar(60) | YES  |     |                   |                             |
| cover_url      | varchar(60) | YES  |     |                   |                             |
| favorite_count | int(11)     | YES  |     | 0                 |                             |
| comment_count  | int(11)     | YES  |     | 0                 |                             |
| title          | text        | YES  |     | NULL              |                             |
| create_date    | datetime    | NO   |     | CURRENT_TIMESTAMP |                             |
| update_date    | datetime    | NO   |     | CURRENT_TIMESTAMP | on update CURRENT_TIMESTAMP |
+----------------+-------------+------+-----+-------------------+-----------------------------+
*/

// // PublishList 返回一系列视频，带有authorID
// func (m *Manager) PublishList(ctx context.Context, authorId int64) ([]*Video, error) {
// 	var pubList []*Video
// 	err := m.handler.WithContext(ctx).Model(&Video{}).Where(&Video{AuthorID: int(authorId)}).Find(&pubList).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return pubList, nil
// }
