package server

import (
	"context"
	//"net/http"
	"bytes"
	"fmt"
	"image/jpeg"
	"image"
	"os"
	"log"
	"math/rand"
	"time"
	"strings"
	//"gorm.io/gorm"
	"github.com/gofrs/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/JirafaYe/publish/internal/service"
	"github.com/JirafaYe/publish/internal/store/local"
)

//TODO: 解决token问题
	// 解析token，得到user_id
	// claim, err := jwt.ParseToken(request.Token)
	// if err!= nil {
	// 	log.Printf("failed to parse token: %v", err)
    //     return nil, err
    // }
	// uid := int(claim.Id)

type PublishSrv struct {
	service.UnimplementedPublishServer
}

func (c *PublishSrv) PubAction (ctx context.Context, request *service.PublishActionRequest) (*service.PublishActionResponse, error){
	// 暂时模拟uid
	rand.Seed(time.Now().UnixNano())
	uid := rand.Intn(10)+1
	//暂时写死
	MinioVideoBucketName := "videos"
	MinioCoverBucketName := "images"
	// 获取视频
	videoData := []byte(request.Data1)
	reader := bytes.NewReader(videoData)// []byte -> reader
	u2, err := uuid.NewV4()// 返回随机的uuid
	if err != nil {
		return nil, err
	}
	fileName := u2.String() + "." + "mp4"

	// 上传视频
	err = m.objectStorer.UploadFile(MinioVideoBucketName, fileName, reader, int64(len(videoData)))
	// 获取视频链接
	url, err := m.objectStorer.GetFileURL(MinioVideoBucketName, fileName, 0)
	playURL := strings.Split(url.String(), "?")[0]
	if err != nil {
		return nil, err
	}

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
	coverURL, err := m.objectStorer.GetFileURL(MinioCoverBucketName, coverPath, 0)
	if err != nil {
		return nil, err
	}
	CoverURL := strings.Split(coverURL.String(), "?")[0]

	// 封装video: user_id, play_url, cover_url, title
	// TODO: 从token中提取user_id
	videoModel := &local.Video{
		PlayURL:  playURL,
		CoverURL: CoverURL,
		Title: 	  request.Title,
		UserId:   int64(uid),
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

// func (c *PublishSrv) PublishList (ctx context.Context, request *service.DouyinPublishListRequest) (*service.DouyinPublishListResponse, error){
// 	var response service.DouyinPublishListResponse

// 	return &response, nil
// }

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