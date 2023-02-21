package server

import (
	// Action 
	"context"
	"bytes"
	"fmt"
	"image/jpeg"
	"image"
	"os"
	"log"
	// "math/rand"
	// "time"
	"strings"
	"github.com/gofrs/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/JirafaYe/publish/internal/service"
	"github.com/JirafaYe/publish/internal/store/local"

	// List
	"github.com/JirafaYe/publish/internal/store/obs"
	jwt"github.com/JirafaYe/publish/pkg/jwt"
)

type PublishSrv struct {
	service.UnimplementedPublishServer
}

func (c *PublishSrv) PubAction (ctx context.Context, request *service.PublishActionRequest) (*service.PublishActionResponse, error){
	// 解析token得到user_id
	claims, err := jwt.ParseToken(request.Token)
	if err != nil {
        fmt.Printf("failed to parse token: %v", err)
    }
	uid := claims.Id
	fmt.Printf("uid: %v\n", uid)
	// 暂时模拟uid
	// rand.Seed(time.Now().UnixNano())
	// uid := rand.Intn(10)+1
	//暂时写死
	MinioVideoBucketName := "videos"
	MinioCoverBucketName := "images"
	// 获取视频
	videoData := []byte(request.Data)
	reader := bytes.NewReader(videoData)// []byte -> reader
	u2, err := uuid.NewV4()// 返回随机的uuid
	if err != nil {
		return nil, err
	}
	fileName := u2.String() + "." + "mp4"

	// 上传视频
	err = m.objectStorer.UploadFile(MinioVideoBucketName, fileName, reader, int64(len(videoData)))
	if err != nil {
		return nil, err
	}
	// 获取视频链接
	url, err := m.objectStorer.GetFileURL(MinioVideoBucketName, fileName, 0)
	playURL := strings.Split(url.String(), "?")[0]
	if err != nil {
		return nil, err
	}
	playURL_database := fileName// 存进数据库的，较为简洁

	// 获取封面
	u3, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	coverPath := u3.String() + "." + "jpg"
	coverData, err := readFrameAsJpeg(playURL)
	if err != nil {
		return nil, err
	}

	// 上传封面
	coverReader := bytes.NewReader(coverData)
	err = m.objectStorer.UploadFile(MinioCoverBucketName, coverPath, coverReader, int64(len(coverData)))
	if err != nil {
		return nil, err
	}
	// 获取封面链接
	// coverURL, err := m.objectStorer.GetFileURL(MinioCoverBucketName, coverPath, 0)
	// if err != nil {
	// 	return nil, err
	// }
	// CoverURL := strings.Split(coverURL.String(), "?")[0]
	CoverURL_database := coverPath

	// 封装video: user_id, play_url, cover_url, title
	videoModel := &local.Video{
		PlayURL:  		playURL_database,
		CoverURL: 		CoverURL_database,
		Title: 	  		request.Title,
		UserId:   		int64(uid),
	}
	// 数据库中插入视频
	err = m.localer.CreateVideo(videoModel)
	if err != nil {
		log.Printf("create video failed: %v", err)
		return nil, err
	}
	// TODO: 返回response
	return &service.PublishActionResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	}, err
}

/* t_video 表
+----------------+---------------------+------+-----+---------+----------------+
| Field          | Type                | Null | Key | Default | Extra          |
+----------------+---------------------+------+-----+---------+----------------+
| id             | bigint(20) unsigned | NO   | PRI | NULL    | auto_increment |
| created_at     | datetime(3)         | YES  |     | NULL    |                |
| updated_at     | datetime(3)         | YES  |     | NULL    |                |
| deleted_at     | datetime(3)         | YES  | MUL | NULL    |                |
| play_url       | varchar(200)        | YES  |     | NULL    |                |
| cover_url      | varchar(200)        | YES  |     | NULL    |                |
| title          | varchar(200)        | YES  |     | NULL    |                |
| favorite_count | bigint(20)          | YES  |     | NULL    |                |
| comment_count  | bigint(20)          | YES  |     | NULL    |                |
| is_favorite    | tinyint(1)          | YES  |     | NULL    |                |
| user_id        | bigint(20)          | NO   |     | NULL    |                |
+----------------+---------------------+------+-----+---------+----------------+
*/
// 从视频流中截取封面
// 测试完毕，可以从视频中截取封面（测试时函数做了修改，返回img image.Image）
func readFrameAsJpeg(filePath string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(filePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)

	return buf.Bytes(), err
}


// TODO: publish list 函数
func (c *PublishSrv) PubList (ctx context.Context, request *service.PublishListRequest) (*service.PublishListResponse, error){
	tmpUserId := request.UserId
	
	var response service.PublishListResponse
	response.StatusCode = 0
	//response.StatusMsg = util.NewString("OK")
	response.StatusMsg = "OK"
	// videos, err := m.localer.QueryVideosByUserId(tmpUserId)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	videos := m.localer.QueryVideosByUserId(tmpUserId)
	for _, v := range videos {
		user := m.localer.QueryUserById(int64(v.UserId))
		response.VideoList = append(response.VideoList, &service.PubVideo{
			Id: v.UserId,
			Author: &service.PubUser{
				Id:            int64(user.ID),
				Name:          user.Name,
				FollowerCount: user.FollowerCount,
				FollowCount:   user.FollowCount,
				IsFollow:      true,
			},
			PlayUrl:       obs.GetVideoPrefix() + v.PlayURL,
			CoverUrl:      obs.GetImagePrefix() + v.CoverURL,
			CommentCount:  v.CommentCount,
			FavoriteCount: v.FavoriteCount,
			IsFavorite:    v.IsFavorite == 1,
			Title:         v.Title,
		})
	}
	return &response, nil
}

/*
type PubVideo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Author        *PubUser `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	PlayUrl       string   `protobuf:"bytes,3,opt,name=play_url,json=playUrl,proto3" json:"play_url,omitempty"`
	CoverUrl      string   `protobuf:"bytes,4,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`
	FavoriteCount int64    `protobuf:"varint,5,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"`
	CommentCount  int64    `protobuf:"varint,6,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`
	IsFavorite    bool     `protobuf:"varint,7,opt,name=is_favorite,json=isFavorite,proto3" json:"is_favorite,omitempty"`
	Title         string   `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`
}
*/