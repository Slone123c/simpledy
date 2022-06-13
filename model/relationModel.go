package model

type Relation struct {
	UserId   int64
	ToUserId int64
	User     UserInformation `gorm:"ForeignKey:UserId"`
	ToUser   UserInformation `gorm:"ForeignKey:ToUserId"`
}
