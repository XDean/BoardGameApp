package model

type Sex uint8

const (
	Unknown Sex = 0
	Male    Sex = 1
	Female  Sex = 2
)

type Profile struct {
	UserID    uint `gorm:"primary_key"`
	Nickname  string
	Sex       Sex
	AvatarURL string
}

func EmptyProfile(id uint) Profile {
	return Profile{
		UserID:    id,
		Nickname:  "Undefined",
		Sex:       Unknown,
		AvatarURL: "",
	}
}
