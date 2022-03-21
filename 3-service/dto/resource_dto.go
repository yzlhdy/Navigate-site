package dto

type ResourceCreateDto struct {
	Name string `json:"name" form:"name" binding:"required,min=2,max=20"`
	Icon string `json:"icon" form:"icon" binding:"required,min=2,max=60"`
	Url  string `json:"url" form:"url" binding:"required,min=2,max=60"`
}

type ResourceUpdateDto struct {
	ID   int    `json:"id" form:"id" binding:"required"`
	Name string `json:"name" form:"name" binding:"required,min=2,max=20"`
	Icon string `json:"icon" form:"icon" binding:"required,min=2,max=60"`
	Url  string `json:"url" form:"url" binding:"required,min=2,max=60"`
}
