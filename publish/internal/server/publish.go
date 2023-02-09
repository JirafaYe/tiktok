package server

import (
	"context"
	"net/http"
	"bytes"
	"fmt"
	"image/jepg"
	"os"
	"strings"
	"gorm.io/gorm"
	"github.com/gofrs/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/Jirafa/publish/internel/store/obs"
	"github.com/Jirafa/publish/internel/store/local"
	"github.com/Jirafa/publish/pkg/jwt""
)

type PublishSrv struct {
	service.UnimplementedPublishServer
}

/*publish的参数*/
// type PublishClient interface {
// 	PubAction(ctx context.Context, in *PublishActionRequest, opts ...grpc.CallOption) (*PublishActionResponse, error)
// 	PubList(ctx context.Context, in *PublishListRequest, opts ...grpc.CallOption) (*PublishListResponse, error)
// }

func (c *PublishSrv) PublishAction (ctx context.Context, request *service.PublishActionRequest) (*service.PublishActionResponse, error){
	//TODO: 解决token问题
	// 解析token，得到user_id
	claim, err := Jwt.ParseToken(request.Token)
	if err!= nil {
		log.Printf("failed to parse token: %v", err)
        return nil, err
    }
	uid := int(claim.Id)

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
		UserId: 			uid,
	}
	return m.localer.CreateVideo(c,ctxm videoModel)
}

// func (c *PublishSrv) PublishList (ctx context.Context, request *service.DouyinPublishListRequest) (*service.DouyinPublishListResponse, error){
// 	var response service.DouyinPublishListResponse

// 	return &response, nil
// }


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