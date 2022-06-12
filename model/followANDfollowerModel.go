package model

type Relation struct {
	UserId3  int64
	ToUserId int64
	User     UserInformation `gorm:"ForeignKey:UserId3"`
	ToUser   UserInformation `gorm:"ForeignKey:ToUserId"`
}
