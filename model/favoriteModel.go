package model

type Favorite struct {
	UserId2 int64
	VideoId int64
	Video   Video           `gorm:"ForeignKey:VideoId"`
	User    UserInformation `gorm:"ForeignKey:UserId2"`
	// Video   Video `gorm:"ForeignKey:VideoId"`
}
