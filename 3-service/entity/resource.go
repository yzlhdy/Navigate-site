package entity

type ResourceType struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	Url  string `json:"url"`
	Model
}
