package model

type User struct {
	Id       int64  `gorm:"primaryKey"`
	Username string `gorm:"size:32;not null"`
	Password string `gorm:"size:32;not null"`
	//	FollowCount   int64  `json:"follow_count,omitempty"`
	//	FollowerCount int64  `json:"follower_count,omitempty"`
	//	IsFollow      bool   `json:"is_follow,omitempty"`
}

type UserInformation struct {
	Id            int64  `gorm:"primaryKey"`
	Name          string `gorm:"size:32;not null"` // 用户名称
	FollowCount   int64  `gorm:"size:64;not null"`
	FollowerCount int64  `gorm:"size:64;not null"`
	// 使得 UserId 成为外键并关联至 User 表
	User   User `gorm:"ForeignKey:UserId"`
	UserId int64
}

/*一共需要以下几个数据表
前四个必须表的详细字段见本文件最后
User表
UserInformation表
Video表
Comment表
favorite表 （可要可不要）
follow表 （可要可不要）
follower表 （可要可不要）
*/
