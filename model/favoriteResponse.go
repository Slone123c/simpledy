package model

type FavoriteActionResponse struct {
	// 0-成功，其他值-失败
	StatusCode int32 `json:"status_code"`
	// 返回状态描述
	StatusMsg string `json:"status_msg"`
}

type FavoriteListResponseAll struct {

	// 0-成功，其他值-失败
	StatusCode int32 `json:"status_code"`

	// 返回状态描述
	StatusMsg string `json:"status_msg"`

	// 用户点赞视频列表
	VideoList *[]FavoriteListResponseVideoList `json:"video_list"`
}

type FavoriteListResponseVideoList struct {

	// 视频唯一标识
	Id int32 `json:"id,omitempty"`

	Author FavoriteListResponseVideoListAuthor `json:"author,omitempty"`

	// 视频播放地址
	PlayUrl string `json:"play_url,omitempty"`

	// 视频封面地址
	CoverUrl string `json:"cover_url,omitempty"`

	// 视频的点赞总数
	FavoriteCount int32 `json:"favorite_count,omitempty"`

	// 视频的评论总数
	CommentCount int32 `json:"comment_count,omitempty"`

	// true-已点赞，false-未点赞
	IsFavorite bool `json:"is_favorite,omitempty"`

	// 视频标题
	Title string `json:"title,omitempty"`
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
