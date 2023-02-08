package pkg

// GetUserFeed implements the FeedSrvImpl interface.
func (s *FeedSrvImpl) GetUserFeed(ctx context.Context, req *feed.DouyinFeedRequest) (resp *feed.DouyinFeedResponse, err error) {
	var uid int64 = 0
	if *req.Token != "" {
		claim, err := Jwt.ParseToken(*req.Token)
		if err != nil {
			resp = pack.BuildVideoResp(errno.ErrTokenInvalid)
			return resp, nil
		} else {
			uid = claim.Id
		}
	}

	vis, nextTime, err := command.NewGetUserFeedService(ctx).GetUserFeed(req, uid)
	if err != nil {
		resp = pack.BuildVideoResp(err)
		return resp, nil
	}

	resp = pack.BuildVideoResp(errno.Success)
	resp.VideoList = vis
	resp.NextTime = &nextTime
	return resp, nil
}

func GetUidFromToken(token string) (uid int64, err error) {
	var uid int64 = 0;
	if token != "" {
		claim, err := Jwt.ParseToken(token)
		if err != nil {
			log.Printf("token invalid")
			return -1, err
		} else {
			uid = claim.Id
		}
	}
	log.Printf("token is empty")
	return -1, nil
}