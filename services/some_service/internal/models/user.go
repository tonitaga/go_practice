package models

type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"size:255;not null" json:"username"`
	Email    string `gorm:"size:255"          json:"email,omitempty"`
	Password string `gorm:"size:512;not null" json:"password"`
	Phone    string `gorm:"size:32"           json:"phone"`
}
