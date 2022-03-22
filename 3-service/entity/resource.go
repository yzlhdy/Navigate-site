package entity

type ResourceType struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	Url  string `json:"url"`

	Model
}

type Resource struct {
	Title        string        `json:"title"`
	SubTitle     string        `json:"sub_title"`
	Image        string        `json:"image"`
	Recommend    bool          `json:"recommend"`
	Rid          int           `json:"rid"`
	ResourceType *ResourceType `json:"resource_type" gorm:"ForeignKey:Rid"`
	Model
}
