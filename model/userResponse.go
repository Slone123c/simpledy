package model

type UserLoginResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	UserId     int    `json:"user_id,omitempty"`
	Token      string `json:"token,omitempty"`
}

type UserInfoResponse struct {
	StatusCode int32   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`  // 返回状态描述
	UserRsp    UserRsp `json:"user"`
}

type UserRsp struct {
	ID            int64  `json:"id"`             // 用户id
	Name          string `json:"name"`           // 用户名称
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
}
