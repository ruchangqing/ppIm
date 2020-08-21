package model

type User struct {
	Id           int
	Username     string
	Password     string
	PasswordSalt string
	Nickname     string
	Avatar       string
	Country      string
	City         string
	Sex          int
	RealName     string
	IdCard       string
	Status       int
	Longitude    string
	Latitude     string
	LastIp       string
	RegisterAt   string `gorm:"default:''"`
	LoginAt      string `gorm:"default:''"`
}
