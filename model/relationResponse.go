package model

type RelationFollowListResponseAll struct {
	// 0-成功，其他值-失败
	StatusCode int32 `json:"status_code"`
	// 返回状态描述
	StatusMsg string `json:"status_msg"`
	// 关注用户列表
	UserList []RelationFollowListResponseUserList `json:"user_list"`
}

type RelationFollowListResponseUserList struct {
	ID            int64  `json:"id"`             // 用户id
	Name          string `json:"name"`           // 用户名称
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
}

type RelationFollowerListResponseAll struct {
	// 0-成功，其他值-失败
	StatusCode int32 `json:"status_code"`
	// 返回状态描述
	StatusMsg string `json:"status_msg"`
	// 关注用户列表
	UserList []RelationFollowerListResponseUserList `json:"user_list"`
}

type RelationFollowerListResponseUserList struct {
	ID            int64  `json:"id"`             // 用户id
	Name          string `json:"name"`           // 用户名称
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
}
