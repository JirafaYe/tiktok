package local

import (
	//"context"
	"gorm.io/gorm"
	//"time"
)

// DB == m.handler

// type Video struct {
// 	gorm.Model
// 	UserId        int64  `gorm:"column:user_id; type:bigint"`
// 	PlayURL       string `gorm:"column:play_url; type:varchar(200)"`
// 	CoverURL      string `gorm:"column:cover_url; type:varchar(200)"`
// 	Title         string `gorm:"column:title; type:varchar(200)"`
// }

// type User struct {
// 	gorm.Model
// 	UserId		  int64  `gorm:"column:user_id; type:bigint"`
// }

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
	UserId		  int64  `gorm:"column:user_id; type:bigint"`
	Name          string `gorm:"column:name; type:varchar(200)"`
	FollowerCount int64  `gorm:"column:follower_count; type:bigint"`
	FollowCount   int64  `gorm:"column:follow_count; type:bigint"`
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

// func (m *Manager) QueryVideosByUserId(userId int64) ([]*Video, error) {
// 	var videos []*Video
// 	db := m.handler.Table("t_video")
// 	err := db.Where("user_id = ?", userId).Order("created_date desc").Find(&videos).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return videos, nil
// }

// TODO: 验证正确性
func (m *Manager) QueryVideosByUserId(userId int64) (videos []*Video) {
	db := m.handler.Table("t_video")
	db.Where("user_id = ?", userId).Order("created_at desc").Find(&videos)
	return
}

// func (m *Manager) QueryNameById(id int64) (name string) {
// 	db := m.handler.Table("t_user")
// 	var user User
// 	db.Where("user_id = ?", id).Find(&user)
// 	name = user.Name
// 	return
// }

// TODO 修改
func (m *Manager) QueryUserById(id int64) (user User) {
	db := m.handler.Table("t_user")
	db.Where("id = ?", id).Find(&user)
	return
}
