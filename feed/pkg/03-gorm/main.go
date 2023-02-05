package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type Video struct {
	gorm.Model
	PlayURL     string `gorm:"column:play_url; type:varchar(200)"`
	CoverURL    string `gorm:"column:cover_url; type:varchar(200)"`
	Title       string `gorm:"column:title; type:varchar(200)"`
	FavorNum    int64  `gorm:"column:favor_num; type:bigint"`
	CommentsNum int64  `gorm:"column:comments_num; type:bigint"`
}

var video []Video

func main() {
	db, err := gorm.Open(
		mysql.Open(
			fmt.Sprintf("root:xh020914@tcp(47.108.66.104:33306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"),
		),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatal(err)
	}
	db = db.Table("t_video")
	db.Where("created_at > ?", time.Time{}).Limit(1).Find(&video)
	fmt.Println(video)
}
