package local

import (
	"testing"
	//"context"
	//"time"
	"fmt"
)

func TestCreateVideo(t *testing.T) {
	// db
	manager, _ := New()
	// Setup context to test
	//ctx := context.Background()
	// ctxWithTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	// defer cancel()

	video := &Video{
        UserId:   int64(1234),
        PlayURL:  "http://47.108.66.104:9000/videos/test1.mp4",
		CoverURL: "http://47.108.66.104:9000/covers/test1.jpg",
		Title:    "the Galaxy",
    }

	//err := manager.CreateVideo(ctx, video)
	err := manager.CreateVideo(video)
    if err != nil {
		t.Errorf("failed to create video since: %v", err)
	}

	// Assert variable was created
	db := manager.handler.Table("t_video")
	var videoRes Video
	err = db.Where("user_id = ?", int64(1234)).Find(&videoRes).Error
	if err != nil {
		t.Errorf("failed to find video: %v", err)
	}
	fmt.Println(videoRes)
}

func TestQueryVideoByUserId(t *testing.T) {
	manager, _ := New()
	// TODO: 表里没数据，无法测试，先改回来吧
	videos, err := manager.QueryVideoByUserId(1234)
	if err!= nil {
		t.Errorf("failed to QueryVideoByUserId: %v", err)
	}
	fmt.Println("video length: ", len(videos))
}

/*
func (m *Manager) QueryVideosByUserId(userId int64) ([]*Video, error) {
	var videos []*Video
	db := m.handler.Table("t_video")
	err := db.Where("user_id = ?", userId).Order("created_date desc").Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}
*/