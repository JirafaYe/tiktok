package main

import (
	"github.com/JirafaYe/publish/internal/store/local"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:xh020914@tcp(47.108.66.104:33306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	generateTestData(db)

}

func generateTable(db *gorm.DB) {
	err := db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&local.User{}, &local.Video{})
	if err != nil {
		return
	}

	err = db.Migrator().RenameTable("users", "t_user")
	if err != nil {
		return
	}
	err = db.Migrator().RenameTable("videos", "t_video")
	if err != nil {
		return
	}
}

func generateTestData(db *gorm.DB) {
	var users []local.User
	var videos []local.Video
	users = append(users, local.User{Name: "zhangsan"}, local.User{Name: "lisi"})
	videos = append(videos, local.Video{
		PlayURL:  "404 Not Found",
		CoverURL: "404 Not Found",
		Title:    "Test Data",
	})
	//db.Table("t_user").Create(&users)
	db.Table("t_video").Create(&videos)

}

