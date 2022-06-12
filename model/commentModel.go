package model

type Comment struct {
	Id      int64 `gorm:"primaryKey"`
	Content string
	// 使得 UserId 成为外键并关联至 User 表
	User      User `gorm:"ForeignKey:UserId"`
	UserId    int64
	Video     Video `gorm:"ForeignKey:VideoId"`
	VideoId   int64
	CreatedAt int64
}
