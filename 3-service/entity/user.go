package entity

type User struct {
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"type:varchar(255)"`
	Avatar   string `json:"avatar" gorm:"type:varchar(255)"`
	Token    string `json:"token"`
	Model
}
