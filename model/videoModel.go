package model

type Video struct {
	Id            int64 `gorm:"size:64;not null"`
	User          User  `gorm:"ForeignKey:AuthorId"`
	AuthorId      int64
	PlayUrl       string `gorm:"size:255;not null"`
	CoverUrl      string `gorm:"size:255;not null"`
	FavoriteCount int64  `gorm:"size:32;not null"`
	CommentCount  int64  `gorm:"size:32;not null"`
	IsFavorite    bool   `gorm:"default:false;not null"`
	Title         string `gorm:"size:64;not null"`
}
