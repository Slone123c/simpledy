package model

type CommentActionResponse struct {
	// 0-成功，其他值-失败
	StatusCode int `json:"status_code"`
	// 返回状态描述
	StatusMsg string `json:"status_msg"`
}

// 评论
type CommentRespElement struct {
	ID         int               `json:"id"`          // 评论id
	User       CommentAuthorInfo `json:"user"`        // 评论用户信息
	Content    string            `json:"content"`     // 评论内容
	CreateDate string            `json:"create_date"` // 评论发布日期，格式 mm-dd
}

// 视频作者信息
type CommentAuthorInfo struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	ID            int64  `json:"id"`             // 用户id
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Name          string `json:"name"`           // 用户名称
}

// 评论列表
type CommentListResponse struct {
	// 0-成功，其他值-失败
	StatusCode int `json:"status_code"`
	// 返回状态描述
	StatusMsg   string               `json:"status_msg"`
	CommentList []CommentRespElement `json:"comment_list"` // 评论列表
}
