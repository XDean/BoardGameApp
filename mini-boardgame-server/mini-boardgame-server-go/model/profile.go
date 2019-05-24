package model

type Profile struct {
	ID        uint `gorm:"primary_key"`
	UserID    uint `gorm:"unique;not null"`
	Nickname  string
	Male      bool
	AvatarURL string
}
