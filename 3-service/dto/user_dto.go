package dto

type UserCreateDto struct {
	Name     string `json:"name" form:"name" binding:"required,min=2,max=20"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=20"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Avatar   string `json:"avatar" form:"avatar"`
}
