package model

type VideoPublishResponse struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}
type VideoFeedResponse struct {
	NextTime   int64       `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	StatusCode int64       `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string      `json:"status_msg"`  // 返回状态描述
	VideoList  []VideoList `json:"video_list"`  // 视频列表
}
type PublishListResponse struct {
	StatusCode int64       `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string      `json:"status_msg"`  // 返回状态描述
	VideoList  []VideoList `json:"video_list"`  // 视频列表
}

type VideoList struct {
	Author        Author `json:"author,omitempty"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count,omitempty"`  // 视频的评论总数
	CoverURL      string `json:"cover_url,omitempty"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count,omitempty"` // 视频的点赞总数
	ID            int64  `json:"id,omitempty"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite,omitempty"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url,omitempty"`       // 视频播放地址
	Title         string `json:"title,omitempty"`          // 视频标题
}

type FavoriteListResponse struct {
	StatusCode int64       `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string      `json:"status_msg"`  // 返回状态描述
	VideoList  []VideoList `json:"video_list"`  // 视频列表
}

// 视频作者信息
type Author struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	ID            int64  `json:"id"`             // 用户id
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Name          string `json:"name"`           // 用户名称
}
