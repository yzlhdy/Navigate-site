package repository

import (
	"navigate/entity"

	"gorm.io/gorm"
)

type ResourceTypeRepository interface {
	ResourceList(page int, limit int) ([]*entity.ResourceType, int64)
	CreateResourceType(resourceType *entity.ResourceType) entity.ResourceType
	UpdateResourceType(id int, resourceType *entity.ResourceType) entity.ResourceType
	FindResourceType(id int) entity.ResourceType
	DeleteResourceType(id int) entity.ResourceType
}

type resourceTypeRepository struct {
	db *gorm.DB
}

func NewResourceTypeRepository(db *gorm.DB) ResourceTypeRepository {
	return &resourceTypeRepository{
		db: db,
	}
}

// 获取资源列表
func (r *resourceTypeRepository) ResourceList(page int, limit int) ([]*entity.ResourceType, int64) {
	var resourceTypes []*entity.ResourceType
	var total int64
	r.db.Offset((page - 1) * limit).Limit(limit).Find(&resourceTypes)
	r.db.Model(&entity.ResourceType{}).Count(&total)
	return resourceTypes, total
}

// 创建资源
func (r *resourceTypeRepository) CreateResourceType(resourceType *entity.ResourceType) entity.ResourceType {
	r.db.Create(&resourceType)
	return *resourceType
}

func (r *resourceTypeRepository) UpdateResourceType(id int, resourceType *entity.ResourceType) entity.ResourceType {
	r.db.Model(&entity.ResourceType{}).Where("id = ?", id).Updates(resourceType)
	return *resourceType
}

func (r *resourceTypeRepository) FindResourceType(id int) entity.ResourceType {
	var resourceType entity.ResourceType
	r.db.First(&resourceType, id)
	return resourceType
}

func (r *resourceTypeRepository) DeleteResourceType(id int) entity.ResourceType {
	var resourceType entity.ResourceType
	r.db.First(&resourceType, id)
	r.db.Delete(&resourceType)
	return resourceType

}
