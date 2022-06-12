package model

type FavoriteActionResponse struct {
	// 0-成功，其他值-失败
	StatusCode int32 `json:"status_code"`
	// 返回状态描述
	StatusMsg string `json:"status_msg"`
}

type FavoriteListResponse struct {

	// 0-成功，其他值-失败
	StatusCode int32 `json:"status_code"`

	// 返回状态描述
	StatusMsg string `json:"status_msg"`

	// 用户点赞视频列表
	VideoList []VideoList `json:"video_list"`
}

type FavoriteListResponseVideoListAuthor struct {

	// 用户id
	Id int32 `json:"id"`

	// 用户名称
	Name string `json:"name"`

	// 关注总数
	FollowCount int32 `json:"follow_count"`

	// 粉丝总数
	FollowerCount int32 `json:"follower_count"`

	// true-已关注，false-未关注
	IsFollow bool `json:"is_follow"`
}
