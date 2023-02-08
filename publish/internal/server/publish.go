package server

import (
	"context"
	"net/http"

	"bytes"
	"fmt"
	"image/jepg"
	"os"
	"strings"

	// "github.com/JirafaYe/example/user/internal/service"
	// "github.com/JirafaYe/example/user/internal/store/local"
	// "github.com/JirafaYe/example/user/pkg/ecrypto"
	// "github.com/JirafaYe/example/user/pkg/token"
	"gorm.io/gorm"

	"github.com/gofrs/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"

	// "github.com/Jirafa/publish/pkg/obs"
	"github.com/Jirafa/publish/internel/store/obs"

	"github.com/Jirafa/publish/internel/store/local"

	"math/rand"
)

type PublishSrv struct {
	service.UnimplementedPublishServer
}

/*publish的参数*/
// type PublishClient interface {
// 	Action(ctx context.Context, in *DouyinPublishActionRequest, opts ...grpc.CallOption) (*DouyinPublishActionResponse, error)
// 	List(ctx context.Context, in *DouyinPublishListRequest, opts ...grpc.CallOption) (*DouyinPublishListResponse, error)
// }

func (c *PublishSrv) PublishAction (ctx context.Context, request *service.DouyinPublishActionRequest) (*service.DouyinPublishActionResponse, error){
	//var response service.DouyinPublishActionResponse
	
	//TODO: 解决token问题
	rand.Seed(time.Now().UnixNano())
	uid := int64(rand.Intn(1000))

	//暂时写死
	MinioVideoBucketName := "videos"
	MinioCoverBucketName := "images"

	videoData := []byte(request.Data)

	// []byte -> reader
	reader := bytes.NewReader(videoData)
	u2, err := uuid.NewV4()// returns random generated UUID
	if err != nil {
		return err
	}
	fileName := u2.String() + "." + "mp4"

	// 上传视频
	err = obs.UploadFile(MinioVideoBucketName, fileName, reader, int64(len(videoData)))
	// 获取视频链接
	url, err := obs.GetFileURL(MinioVideoBucketName, fileName, 0)
	playURL := strings.Split(url.String(), "?")[0]
	if err != nil {
		return err
	}

	u3, err := uuid.NewV4()
	if err != nil {
		return err
	}

	// 获取封面
	coverPath := u3.String() + "." + "jpg"
	coverData, err := readFrameAsJpeg(playURL)
	if err != nil {
		return err
	}
	// 上传封面
	coverReader := bytes.NewReader(coverData)
	err := obs.UploadFile(MinioCoverBucketName, coverPath, coverReader, int64(len(coverData)))
	if err != nil {
		return err
	}
	// 获取封面链接
	coverURL, err := obs.GetFileURL(MinioCoverBucketName, coverPath, 0)
	if err != nil {
		return err
	}

	CoverURL := strings.Split(coverURL.String(), "?")[0]

	// 封装video
	//TODO: 从token中提取user_id; 创建create_time
	videoModel := &Video{
		PlayURL: 			playURL,
		CoverURL: 			CoverURL,
		Title: 				request.Title,
		UserId: 			uid,//TODO
	}
	return m.localer.CreateVideo(c,ctxm videoModel)
}

func (c *PublishSrv) PublishList (ctx context.Context, request *service.DouyinPublishListRequest) (*service.DouyinPublishListResponse, error){
	var response service.DouyinPublishListResponse

	return &response, nil
}


//从视频流中截取封面
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



/*action的request和response*/
// type DouyinPublishActionRequest struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
// 	Data  []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
// 	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
// }


// type DouyinPublishActionResponse struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
// 	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
// }


/*feed/store/local 下定义的Video和User结构体*/
// type Video struct {
// 	gorm.Model
// 	PlayURL       string `gorm:"column:play_url; type:varchar(200)"`
// 	CoverURL      string `gorm:"column:cover_url; type:varchar(200)"`
// 	Title         string `gorm:"column:title; type:varchar(200)"`
// 	FavoriteCount int64  `gorm:"column:favorite_count; type:bigint"`
// 	CommentCount  int64  `gorm:"column:comment_count; type:bigint"`
// 	IsFavorite    int16  `gorm:"column:is_favorite; type:tinyint"`
// 	UserId        int64  `gorm:"column:user_id; type:bigint"`
// }

// type User struct {
// 	gorm.Model
// 	Name          string `gorm:"column:name; type:varchar(200)"`
// 	FollowerCount int64  `gorm:"column:follower_count; type:bigint"`
// 	FollowCount   int64  `gorm:"column:follow_count; type:bigint"`
// }

// func (m *Manager) QueryVideosAfter(n int, date time.Time) (videos []*Video) {
// 	db := m.handler.Table("t_video")
// 	db.Where("created_at < ?", date).Order("created_at desc").Limit(n).Find(&videos)
// 	return
// }

// func (m *Manager) QueryNameById(id int64) (name string) {
// 	db := m.handler.Table("t_user")
// 	var user User
// 	db.Where("id = ?", id).Find(&user)
// 	name = user.Name
// 	return
// }

// func (m *Manager) QueryUserById(id int64) (user User) {
// 	db := m.handler.Table("t_user")
// 	db.Where("id = ?", id).Find(&user)
// 	return