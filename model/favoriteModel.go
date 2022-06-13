package model

type Favorite struct {
	Id      int64
	UserId  int64
	VideoId int64
	Video   Video           `gorm:"ForeignKey:VideoId"`
	User    UserInformation `gorm:"ForeignKey:UserId"`
	// Video   Video `gorm:"ForeignKey:VideoId"`
}
