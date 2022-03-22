package dto

type ResourceTypeCreateDto struct {
	Name string `json:"name" form:"name" binding:"required,min=2,max=20"`
	Icon string `json:"icon" form:"icon" binding:"required,min=2,max=60"`
	Url  string `json:"url" form:"url" binding:"required,min=2,max=60"`
}

type ResourceTypeUpdateDto struct {
	Name string `json:"name" form:"name" binding:"min=2,max=20"`
	Icon string `json:"icon" form:"icon" binding:"min=2,max=60"`
	Url  string `json:"url" form:"url" binding:"min=2,max=60"`
}

type ResourceCreateDto struct {
	Title     string `json:"title" form:"title" binding:"required,min=2,max=20"`
	SubTitle  string `json:"sub_title" form:"sub_title" binding:"required"`
	Image     string `json:"image" form:"image" binding:"required,min=2,max=60"`
	Recommend bool   `json:"recommend" form:"recommend" `
	Rid       int    `json:"rid" form:"rid" binding:"required"`
}

type ResourceUpdateDto struct {
	Title     string `json:"title" form:"title" binding:"required,min=2,max=20"`
	SubTitle  string `json:"sub_title" form:"sub_title" binding:"required"`
	Image     string `json:"image" form:"image" binding:"required,min=2,max=60"`
	Recommend bool   `json:"recommend" form:"recommend" `
	Rid       int    `json:"rid" form:"rid" binding:"required"`
}
